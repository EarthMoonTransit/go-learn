package subject_test

import (
	"testing"

	"task.local/utf8-prefix-001/subject"
)

func TestPrefixRunes(t *testing.T) {
	invalid := string([]byte{'a', 0xff, 'b'})
	tests := []struct {
		name  string
		text  string
		limit int
		want  string
	}{
		{name: "nonpositive", text: "abc", limit: 0, want: ""},
		{name: "multibyte boundary", text: "a界b", limit: 2, want: "a界"},
		{name: "larger limit", text: "go", limit: 9, want: "go"},
		{name: "invalid byte preserved", text: invalid, limit: 2, want: string([]byte{'a', 0xff})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subject.PrefixRunes(tt.text, tt.limit); got != tt.want {
				t.Fatalf("got bytes %v; want %v", []byte(got), []byte(tt.want))
			}
		})
	}
}
