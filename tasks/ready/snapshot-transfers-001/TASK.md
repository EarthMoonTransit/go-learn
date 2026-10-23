# Snapshot independent channel transfers

## Context
Capture descriptors before asynchronous transfers.

## API
```go
type Transfer struct { Source <-chan int; Destination chan<- int; Offset int }
func StartTransfers(transfers []Transfer, start <-chan struct{}) <-chan struct{}
```

## Requirements
- Snapshot descriptors before returning fresh nonnil completion. Caller may replace slice elements after return.
- One worker per descriptor waits for start closure, then receives exactly once from Source.
- Send a received value plus Offset once to Destination. Closed-empty Source sends nothing.
- Workers progress independently: a blocked receive or send must not delay other transfers.
- Close completion after all required sends finish; send no completion values. Empty input completes without start.
- Preserve descriptors; never close supplied channels. Review explains capture and blocking phases.

## Constraints
- At most four descriptors; nonnil distinct endpoints, open destinations, initially open start. start receives no values. Caller closes start and eventually enables transfers. Sums fit int. Standard library only.
