# Share an append journal through copied handles

## Context
Keep copied handles attached to one concurrent journal.

## API
```go
type Journal journalHandle
func NewJournal() Journal
func (j Journal) Append(value int) int
func (j Journal) Snapshot() []int
```

## Requirements
- New journals are empty and independent. Copying a handle shares its original journal.
- Append stores value once and returns its unique zero-based index. Concurrent appends may choose any order consistent with snapshots.
- Snapshot returns the complete sequence at one instant, detached from journal storage and other snapshots. Nil/empty are equivalent.
- Own storage and a sync.Mutex/sync.RWMutex together behind stable private identity. Never copy a used mutex; synchronize every read and write.
- Implement private representation, construction and methods. Review identifies every access path and explains why handle copies retain one state/lock pair.

## Constraints
- Use constructor-created handles only; at most sixteen Append calls per journal. Standard library only. Concurrent method calls and handle copies are allowed.
