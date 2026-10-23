Supplied `Insertion([]int)` and `Merge([]int)` sort integers in place.

- Author `BenchmarkSort` and `TestSortEquivalence`; prove both produce the same ascending values, preserving duplicates and length, before comparing performance.
- Choose two nonzero sizes within 16–512 differing by at least fourfold. For each, include sorted, reverse and unordered duplicate-heavy inputs. Explain the choices.
- Give every timed iteration equivalent fresh working state. Exclude restoration, allocation of working input, and validation from measured work. Use the native benchmark iteration API and keep results observable.
- Report native ns/op, B/op and allocs/op for both operations over three runs. Attribute algorithm scratch and working-buffer costs to their actual boundaries.
- Interpret per-case variation and allocation/lifetime tradeoffs; neither fewer allocations nor one timing establishes a general winner.
- Submit source, equivalence evidence, repeated output and boundary reasoning.

Inputs are deterministic and bounded; use the standard library. Follow the reviewed disposable procedure. No numeric performance threshold is a grade.
