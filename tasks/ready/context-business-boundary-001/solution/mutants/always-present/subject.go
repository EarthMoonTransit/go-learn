package subject

import "context"

type Metadata struct{ Trace string }
type metadataKey struct{}

func WithMetadata(parent context.Context, value Metadata) context.Context {
	return context.WithValue(parent, metadataKey{}, value)
}
func MetadataOf(ctx context.Context) (Metadata, bool) {
	value, ok := ctx.Value(metadataKey{}).(Metadata)
	_ = ok
	return value, true
}
func Apply(ctx context.Context, value, multiplier int, next func(context.Context, int) error) error {
	return next(ctx, value*multiplier)
}
