package catalog

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"go-practice-bank/internal/safeio"
)

const (
	maxCandidateLineBytes = 16 << 20
	maxCatalogBytes       = 64 << 20
)

func Sort(candidates []Candidate) {
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].SourceID != candidates[j].SourceID {
			return candidates[i].SourceID < candidates[j].SourceID
		}
		if candidates[i].SourcePath != candidates[j].SourcePath {
			return candidates[i].SourcePath < candidates[j].SourcePath
		}
		if candidates[i].ExerciseKind != candidates[j].ExerciseKind {
			return candidates[i].ExerciseKind < candidates[j].ExerciseKind
		}
		return candidates[i].Slug < candidates[j].Slug
	})
}

func MarshalJSONL(candidates []Candidate) ([]byte, error) {
	copyOfCandidates := append([]Candidate(nil), candidates...)
	Sort(copyOfCandidates)
	var output bytes.Buffer
	encoder := json.NewEncoder(&output)
	encoder.SetEscapeHTML(false)
	for _, candidate := range copyOfCandidates {
		if err := encoder.Encode(candidate); err != nil {
			return nil, fmt.Errorf("encode candidate %q: %w", candidate.Slug, err)
		}
		if output.Len() > maxCatalogBytes {
			return nil, fmt.Errorf("candidate catalog exceeds %d bytes", maxCatalogBytes)
		}
	}
	return output.Bytes(), nil
}

func WriteJSONL(root, relative string, candidates []Candidate) (bool, error) {
	b, err := MarshalJSONL(candidates)
	if err != nil {
		return false, err
	}
	return safeio.AtomicWrite(root, relative, b, maxCatalogBytes)
}

func ReadJSONL(root, relative string) ([]Candidate, error) {
	b, err := safeio.ReadFile(root, relative, maxCatalogBytes)
	if err != nil {
		return nil, fmt.Errorf("read candidate catalog: %w", err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	scanner.Buffer(make([]byte, 64*1024), maxCandidateLineBytes)
	var candidates []Candidate
	line := 0
	for scanner.Scan() {
		line++
		var candidate Candidate
		decoder := json.NewDecoder(bytes.NewReader(scanner.Bytes()))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&candidate); err != nil {
			return nil, fmt.Errorf("decode candidate line %d: %w", line, err)
		}
		var extra any
		if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("candidate line %d contains trailing JSON", line)
		}
		candidates = append(candidates, candidate)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read candidate catalog: %w", err)
	}
	return candidates, nil
}
