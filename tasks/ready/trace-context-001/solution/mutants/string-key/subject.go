package subject

import "context"

type traceKey struct{}

func WithTrace(parent context.Context, id string) context.Context {
	return context.WithValue(parent, "trace", id)
}
func Trace(ctx context.Context) (string, bool) { id, ok := ctx.Value("trace").(string); return id, ok }
