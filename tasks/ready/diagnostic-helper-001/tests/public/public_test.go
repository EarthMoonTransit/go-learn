package subject_test

import (
	"fmt"
	"reflect"
	"slices"
	"task.local/diagnostic-helper-001/subject"
	"testing"
)

type recordReporter struct {
	events, messages []string
	format           string
	args             []any
}

func (r *recordReporter) Helper() { r.events = append(r.events, "helper") }
func (r *recordReporter) Errorf(format string, args ...any) {
	r.events = append(r.events, "error")
	r.format = format
	r.args = append([]any(nil), args...)
	r.messages = append(r.messages, fmt.Sprintf(format, args...))
}

func TestPublicCase01(t *testing.T) {
	r := &recordReporter{}
	subject.AssertInt(r, 3, 3)
	if !slices.Equal(r.events, []string{"helper"}) || len(r.messages) != 0 {
		t.Errorf("equal events%v messages%v", r.events, r.messages)
	}
}

func TestPublicCase02(t *testing.T) {
	r := &recordReporter{}
	subject.AssertInt(r, -2, 0)
	if !slices.Equal(r.events, []string{"helper", "error"}) || !slices.Equal(r.messages, []string{"got -2, want 0"}) || r.format != "got %d, want %d" || !reflect.DeepEqual(r.args, []any{-2, 0}) {
		t.Errorf("helper diagnostic events%v messages%v format%q args%v", r.events, r.messages, r.format, r.args)
	}
	subject.AssertInt(r, 7, 7)
	if len(r.events) != 3 || len(r.messages) != 1 {
		t.Error("later equality reported")
	}
	subject.AssertInt(t, 1, 1)
}
