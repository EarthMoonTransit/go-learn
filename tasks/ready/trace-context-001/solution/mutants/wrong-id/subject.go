package subject

import "context"

type traceKey struct{}

func WithTrace(parent context.Context, id string) context.Context {
	return context.WithValue(parent, traceKey{}, "fixed")
}
func Trace(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(traceKey{}).(string)
	return id, ok
}
