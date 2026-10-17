# Store and find values by a composite map key

## Context

Implement insertion and lookup for the supplied comparable struct key. Each table owns independent storage.

## API

```go
type Key struct { Parts [2]int; Class int }
func NewTable() *table
func Put(t *table, key *Key, value string)
func Lookup(t *table, key Key) (string, bool)
```

## Requirements

- NewTable returns an independent empty table.
- Equal keys address one entry; another Put replaces its value.
- Changing either array component or Class identifies a different key.
- Copy the key's value at insertion; later caller mutation must not move that entry.
- Return the exact stored string and true, including an empty string.
- Return ("", false) for absent keys without modifying the table.

## Constraints

- table is supplied private storage; obtain it through NewTable. Table and key pointers are non-nil. All int components and strings are valid. Edit subject/subject.go. No I/O or concurrency.
