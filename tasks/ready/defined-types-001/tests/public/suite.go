package publictests

import (
	"reflect"
	"testing"
)

type API struct {
	SamplesZero any
	BytesZero   any
	EncodedSize func(int, int) int
	Combine     func(int, int) int
}

func Run(t *testing.T, api API) {
	t.Helper()
	if name := reflect.TypeOf(api.SamplesZero).Name(); name != "Samples" {
		t.Errorf("Samples must be a defined type: got type name %q", name)
	}
	if name := reflect.TypeOf(api.BytesZero).Name(); name != "Bytes" {
		t.Errorf("Bytes must be a defined type: got type name %q", name)
	}
	if got, want := api.EncodedSize(7, 1), 7; got != want {
		t.Errorf("EncodedSize(7, 1): got %d, want %d", got, want)
	}
	if got, want := api.Combine(4, 3), 7; got != want {
		t.Errorf("Combine(4, 3): got %d, want %d", got, want)
	}
}
