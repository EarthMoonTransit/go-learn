package subject

import (
	"context"
	"strings"
)

type traceKey struct{}

func WithTrace(parent context.Context, id string) context.Context {
	return context.WithValue(parent, traceKey{}, strings.TrimSpace(id))
}
func Trace(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(traceKey{}).(string)
	return id, ok
}
