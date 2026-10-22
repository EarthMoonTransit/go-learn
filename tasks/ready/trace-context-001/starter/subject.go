package subject

import "context"

func WithTrace(parent context.Context, id string) context.Context { return parent }
func Trace(ctx context.Context) (string, bool)                    { return "", false }
