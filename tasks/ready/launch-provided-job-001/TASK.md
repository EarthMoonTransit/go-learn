# Launch a supplied job and return its handle

## Context

Complete StartJob using the supplied runner and synchronization handle.

## API

```go
type Job struct { Run func(); Handle Handle }
func StartJob(job Job) Handle
// supplied: NewJob() Job; Handle.Release(), Done(), Calls()
```

## Requirements

- Launch job.Run in a new goroutine exactly once.
- Return job.Handle without waiting for the runner to finish.
- Before Release, the supplied runner remains blocked and completion remains pending.
- After Release, the runner finishes and its supplied completion channel closes.
- Keep separate jobs independent; return each job's own handle.

## Constraints

- Inputs come from NewJob; Run and Handle are valid. Edit only subject/implementation.go. Handle owns all synchronization; do not implement or change its internals. Use only the standard library.
