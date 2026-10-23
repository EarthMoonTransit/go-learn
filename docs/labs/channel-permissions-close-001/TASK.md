# Restrict helpers and observe closure

## Context
Report stores readings (Value, OK, Ready) and SendPanicked/ClosePanicked.

## API
```go
func Observe(values []int) Report
```

## Requirements
- Create a channel buffered for all values. Use one helper to send them in order and close it, without goroutines.
- Author helper signatures: the producer may only send/close; the receiver may only receive. Choose their channel parameter types.
- After closure, record len(values)+2 nonblocking comma-ok receives. Ready distinguishes an available receive from an open empty channel; queued zero remains data.
- Attempt an owner-side send, then a second close, each inside the supplied recovery boundary. Record both outcomes. The send attempt must be nonblocking.
- Predict the readings and panic flags before running; compare predictions with retained actual results.
- Reproduce native compiler rejection of receive in the producer and send/close in the receiver. Repair each without widening permissions.

## Constraints
Input contains 0–4 integers. Edit only Observe and your helpers; submit predictions separately.
