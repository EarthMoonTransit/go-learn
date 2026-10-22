# Measure equivalent string joins

Prerequisites: format-vet-001 and testing.basics-diagnostics. This manual unit has three separate first encounters, in order: benchmarks.loops-timing, benchmarks.allocations, benchmarks.experimental-design.

Concat and Builder must both join strings with a slash, preserving bytes and input. First run supplied correctness checks against strings.Join. Write your own BenchmarkJoin in benchmark_test.go: use the benchmark iteration count, separate fixture setup from measured joining, and retain the result without changing its semantics. Identify these boundaries in your code.

Then report B/op and allocs/op on representative inputs; explain temporary concatenation allocations, builder storage and the lifetime of the returned string. Allocation count alone is not a speed or correctness result.

Finally choose empty, short/long ASCII and multibyte fixtures, compare both implementations with at least three runs, and report observed ranges/variation with equivalence evidence. A single tiny case or single timing cannot establish a general speedup. No numeric timing threshold is a grade.

Submit actual source, inputs, repeated native output and analysis using PROCEDURE.md. Supplied implementations/tests earn no learner credit. The starter supplies only a benchmark entry point; you choose its body and cases. Keep reference material separate from your submission.
