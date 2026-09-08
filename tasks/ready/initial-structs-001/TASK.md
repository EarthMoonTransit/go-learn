# Construct initialized struct values

## Context

Complete the specified operations.

## API

```go
type Motor struct{Speed,Drain,Battery,Distance int}
type Track struct{Distance int}
func NewMotor(speed, drain int) Motor
func NewTrack(distance int) Track
```

## Requirements

- NewMotor copies speed and drain into the corresponding fields.
- A new motor has Battery 100 and Distance zero.
- NewTrack stores the supplied distance.
- All inputs are nonnegative integers; zero values are valid inputs.

## Constraints

- Use only the Go standard library.
