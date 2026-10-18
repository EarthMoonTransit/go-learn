package frame

import (
	"go/ast"
	"go/build"
	"io"
	"strings"
)

func (a *frame) checkSelection(name, source string, file *ast.File) error {
	for _, group := range file.Comments {
		for _, comment := range group.List {
			text := strings.TrimSpace(comment.Text)
			plusBuild := strings.HasPrefix(text, "//") && strings.HasPrefix(strings.TrimSpace(strings.TrimPrefix(text, "//")), "+build")
			if strings.HasPrefix(text, "//go:") || plusBuild || strings.HasPrefix(text, "//line ") || strings.HasPrefix(text, "/*line ") {
				return a.fail(comment, "frame.directives", "build/compiler/line directives are unsupported")
			}
		}
	}
	// A synthetic target rejects every recognized OS/architecture filename
	// suffix and Go-ignored filename. MatchFile reads only this in-memory file.
	context := build.Context{GOOS: "go_learn_common", GOARCH: "go_learn_common", OpenFile: func(string) (io.ReadCloser, error) { return io.NopCloser(strings.NewReader(source)), nil }}
	selected, err := context.MatchFile("", name)
	if err != nil || !selected {
		return a.fail(file, "frame.files", "source filename must be selected on every platform")
	}
	return nil
}
