package subject_test

import (
	"slices"
	"testing"

	"task.local/rune-offset-tests-001/subject"
)

func TestRuneOffsets(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []int
	}{
		{name: "empty", text: "", want: []int{0}},
		{name: "ASCII", text: "go", want: []int{0, 1, 2}},
		{name: "multibyte", text: "éx", want: []int{0, 2, 3}},
		{name: "invalid byte", text: string([]byte{'a', 0xff, 'b'}), want: []int{0, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subject.RuneOffsets(tt.text); !slices.Equal(got, tt.want) {
				t.Fatalf("RuneOffsets bytes %v = %v; want %v", []byte(tt.text), got, tt.want)
			}
		})
	}
}
