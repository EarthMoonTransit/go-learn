package subject

import "context"

type Metadata struct{ Trace string }
type metadataKey struct{}

func WithMetadata(parent context.Context, value Metadata) context.Context {
	return context.WithValue(context.Background(), metadataKey{}, value)
}
func MetadataOf(ctx context.Context) (Metadata, bool) {
	value, ok := ctx.Value(metadataKey{}).(Metadata)
	return value, ok
}
func Apply(ctx context.Context, value, multiplier int, next func(context.Context, int) error) error {
	return next(ctx, value*multiplier)
}
