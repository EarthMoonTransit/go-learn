package gradertests

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
	samplesType := reflect.TypeOf(api.SamplesZero)
	bytesType := reflect.TypeOf(api.BytesZero)
	if samplesType.Name() != "Samples" || samplesType.PkgPath() == "" {
		t.Errorf("Samples is not a package-defined type: %v", samplesType)
	}
	if bytesType.Name() != "Bytes" || bytesType.PkgPath() == "" {
		t.Errorf("Bytes is not a package-defined type: %v", bytesType)
	}
	if samplesType == bytesType {
		t.Error("Samples and Bytes must be distinct types")
	}
	for _, test := range []struct {
		count, width, want int
	}{
		{count: 0, width: 8, want: 0},
		{count: 6, width: 2, want: 12},
		{count: 9, width: 4, want: 36},
	} {
		if got := api.EncodedSize(test.count, test.width); got != test.want {
			t.Errorf("EncodedSize(%d, %d): got %d, want %d", test.count, test.width, got, test.want)
		}
	}
	if got, want := api.Combine(0, 12), 12; got != want {
		t.Errorf("Combine(0, 12): got %d, want %d", got, want)
	}
}
