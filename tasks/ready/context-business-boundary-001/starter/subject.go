package subject

import "context"

type Metadata struct{ Trace string }

func WithMetadata(parent context.Context, value Metadata) context.Context {
	return context.WithValue(parent, "metadata", value)
}
func MetadataOf(ctx context.Context) (Metadata, bool) {
	value, ok := ctx.Value("metadata").(Metadata)
	return value, ok
}
func Apply(ctx context.Context, value, multiplier int, next func(context.Context, int) error) error {
	if v, ok := ctx.Value("value").(int); ok {
		value = v
	}
	if m, ok := ctx.Value("multiplier").(int); ok {
		multiplier = m
	}
	if options, ok := ctx.Value("options").(map[string]int); ok {
		multiplier = options["multiplier"]
	}
	return next(context.WithValue(ctx, "value", value), value*multiplier)
}
