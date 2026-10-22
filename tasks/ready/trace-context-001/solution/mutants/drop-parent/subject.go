package subject

import "context"

type traceKey struct{}

func WithTrace(parent context.Context, id string) context.Context {
	return context.WithValue(context.Background(), traceKey{}, id)
}
func Trace(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(traceKey{}).(string)
	return id, ok
}
