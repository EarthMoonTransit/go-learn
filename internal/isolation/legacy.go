package isolation

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var packagePattern = regexp.MustCompile(`^\./(?:\.\.\.|[a-zA-Z0-9_-]+(?:/[a-zA-Z0-9_-]+)*(?:/\.\.\.)?)$`)

// LegacyCommand is the parsed, non-shell representation of one v1 manifest
// validation command.
type LegacyCommand struct {
	ExpectFailure bool
	Args          []string
}

func ParseLegacyCommand(value string) (LegacyCommand, error) {
	if strings.ContainsAny(value, "`|&;<>\\\r\n\x00") {
		return LegacyCommand{}, errors.New("shell syntax and expansion are forbidden")
	}
	words, err := splitLegacyWords(value)
	if err != nil {
		return LegacyCommand{}, err
	}
	result := LegacyCommand{}
	if len(words) > 0 && words[0] == "!" {
		result.ExpectFailure = true
		words = words[1:]
	}
	if len(words) < 3 || words[0] != "go" || (words[1] != "test" && words[1] != "vet") {
		return LegacyCommand{}, errors.New("only go test and go vet commands are allowed")
	}
	result.Args = append(result.Args, words[1])
	packageSeen := false
	for index := 2; index < len(words); index++ {
		word := words[index]
		if strings.Contains(word, "$") && word != "^$" && word != "-run=^$" {
			return LegacyCommand{}, errors.New("environment and command expansion are forbidden")
		}
		if packagePattern.MatchString(word) {
			packageSeen = true
			result.Args = append(result.Args, word)
			continue
		}
		if words[1] == "vet" {
			return LegacyCommand{}, fmt.Errorf("vet flag or argument %q is not allowed", word)
		}
		switch {
		case word == "-race", word == "-count=1", word == "-run=^$":
			result.Args = append(result.Args, word)
		case word == "-run" && index+1 < len(words) && words[index+1] == "^$":
			result.Args = append(result.Args, word, words[index+1])
			index++
		default:
			return LegacyCommand{}, fmt.Errorf("test flag or argument %q is not allowed", word)
		}
	}
	if !packageSeen {
		return LegacyCommand{}, errors.New("at least one explicit ./ package is required")
	}
	return result, nil
}

func splitLegacyWords(value string) ([]string, error) {
	var words []string
	var current strings.Builder
	quote := rune(0)
	flush := func() {
		if current.Len() != 0 {
			words = append(words, current.String())
			current.Reset()
		}
	}
	for _, character := range strings.TrimSpace(value) {
		switch {
		case quote != 0 && character == quote:
			quote = 0
		case quote != 0:
			current.WriteRune(character)
		case character == '\'' || character == '"':
			quote = character
		case character == ' ' || character == '\t':
			flush()
		default:
			current.WriteRune(character)
		}
	}
	if quote != 0 {
		return nil, errors.New("unterminated quote")
	}
	flush()
	return words, nil
}
