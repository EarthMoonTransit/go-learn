package subject

import "context"

type traceKey struct{}

func WithTrace(parent context.Context, id string) context.Context {
	if _, ok := Trace(parent); ok {
		return parent
	}
	return context.WithValue(parent, traceKey{}, id)
}
func Trace(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(traceKey{}).(string)
	return id, ok
}
