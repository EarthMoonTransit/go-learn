package subject

import "reflect"

func mapEntries[A, B any](values []A, f func(A) B) []Entry[A, B] {
	out := make([]Entry[A, B], len(values))
	for i, value := range values {
		v := reflect.ValueOf(value).Interface().(A)
		out[i] = Entry[A, B]{Input: v, Output: f(v)}
	}
	return out
}
