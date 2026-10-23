# Separate request metadata from business arguments
## Context
Repair the supplied Context-based overrides.
## API
```go
type Metadata struct { Trace string }
func WithMetadata(parent context.Context, value Metadata) context.Context
func MetadataOf(ctx context.Context) (Metadata, bool)
func Apply(ctx context.Context, value, multiplier int,
    next func(context.Context, int) error) error
```
## Requirements
- Store Metadata behind a private comparable key. Read safely: absent or wrong dynamic type returns zero,false; an empty Trace is present.
- Nested metadata shadows only metadata, preserving parent values, deadline and cancellation.
- Apply computes value*multiplier exclusively from explicit arguments; Context entries named value, multiplier or options cannot override them.
- Call next synchronously exactly once with the incoming context and product, including when canceled; return its exact error.
- Inject no business parameters into Context and retain no context after return. Explain this boundary in your repair.
## Constraints
- Contexts/callbacks are nonnil; integers are in [-8,8]. Method review inspects storage and forwarding.
