package subject_test

import (
	"context"
	"errors"
	"reflect"
	subject "task.local/context-business-boundary-001/subject"
	"testing"
	"time"
)

type requestMarker struct{}
type wrongDynamic struct{ context.Context }

func (wrongDynamic) Value(any) any { return 42 }

type noncomparable struct {
	context.Context
	marker []int
}

func metadata(t *testing.T, ctx context.Context, want string, ok bool) {
	t.Helper()
	v, p := subject.MetadataOf(ctx)
	if v.Trace != want || p != ok {
		t.Errorf("metadata=(%q,%t), want(%q,%t)", v.Trace, p, want, ok)
	}
}
func TestMetadataBoundary(t *testing.T) {
	base := context.Background()
	metadata(t, base, "", false)
	metadata(t, wrongDynamic{base}, "", false)
	forged := context.WithValue(base, "metadata", subject.Metadata{Trace: "foreign"})
	metadata(t, forged, "", false)
	empty := subject.WithMetadata(base, subject.Metadata{})
	metadata(t, empty, "", true)
	first := subject.WithMetadata(base, subject.Metadata{Trace: "A"})
	second := subject.WithMetadata(first, subject.Metadata{Trace: "B"})
	metadata(t, first, "A", true)
	metadata(t, second, "B", true)
	metadata(t, base, "", false)
	deadline := time.Date(2099, 2, 3, 0, 0, 0, 0, time.UTC)
	parent, cancel := context.WithDeadline(context.WithValue(base, requestMarker{}, "keep"), deadline)
	defer cancel()
	child := subject.WithMetadata(parent, subject.Metadata{Trace: " \nΩ "})
	metadata(t, child, " \nΩ ", true)
	if child.Value(requestMarker{}) != "keep" {
		t.Error("metadata lost unrelated parent value")
	}
	if d, ok := child.Deadline(); !ok || !d.Equal(deadline) {
		t.Error("metadata lost parent deadline")
	}
	cancel()
	if child.Err() != context.Canceled {
		t.Error("metadata lost parent cancellation")
	}
}
func TestExplicitBusinessArguments(t *testing.T) {
	for _, canceled := range []bool{false, true} {
		for _, pair := range [][2]int{{2, 3}, {0, 7}, {-8, 8}, {-3, -2}} {
			for _, overrides := range []bool{false, true} {
				parent, cancel := context.WithCancel(context.Background())
				if canceled {
					cancel()
				}
				var ctx context.Context = context.WithValue(parent, requestMarker{}, "same request")
				if overrides {
					ctx = context.WithValue(ctx, "value", -7)
					ctx = context.WithValue(ctx, "multiplier", -6)
					ctx = context.WithValue(ctx, "options", map[string]int{"multiplier": -5})
				}
				ctx = subject.WithMetadata(ctx, subject.Metadata{Trace: "request"})
				ctx = noncomparable{ctx, []int{1}}
				marker := errors.New("callback failure")
				calls := 0
				err := subject.Apply(ctx, pair[0], pair[1], func(got context.Context, value int) error {
					calls++
					if value != pair[0]*pair[1] {
						t.Errorf("business result=%d; want explicit product=%d", value, pair[0]*pair[1])
					}
					if got == nil {
						t.Error("callback context is nil")
						return marker
					}
					metadata(t, got, "request", true)
					if got.Value(requestMarker{}) != "same request" || got.Done() != ctx.Done() || got.Err() != ctx.Err() {
						t.Error("callback lost incoming context/control")
					}
					for _, key := range []string{"value", "multiplier", "options"} {
						if !reflect.DeepEqual(got.Value(key), ctx.Value(key)) {
							t.Errorf("callback context changed business entry %s", key)
						}
					}
					return marker
				})
				if calls != 1 {
					t.Errorf("callback calls=%d; want exactly1 even if canceled", calls)
				}
				if err != marker {
					t.Errorf("return=%v; want exact callback error", err)
				}
				cancel()
			}
		}
	}
	calls := 0
	if err := subject.Apply(context.Background(), 1, 0, func(context.Context, int) error { calls++; return nil }); err != nil || calls != 1 {
		t.Errorf("nil error path: err=%v calls=%d", err, calls)
	}
}
