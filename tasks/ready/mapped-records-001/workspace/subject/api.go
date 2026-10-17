package subject

// Entry is supplied support: its type arguments select the two field types.
// The learner implements a generic function, not this record type.
type Entry[A, B any] struct {
	Input  A
	Output B
}
