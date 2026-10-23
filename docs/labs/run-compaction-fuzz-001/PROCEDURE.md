# Finite native compaction procedure

Follow accepted hex target/seed and independent-invariant first encounters and
fuzz-regression reproduction. This manual application owns the full native
FuzzCompact target/seeds, independent lossy semantics and actual finite reduction.
Supplied production/public/grader sources earn no learner authorship credit.

Copy only manifest.initial_files into a fresh private offline disposable project,
removing `.txt`. The inert target template stays outside the compiled tree until
learner-authored full declarations exist. Allowed edits are compact.go,
fuzz_test.go, regression_test.go and the named native corpus file. Keep public,
grader and module files unchanged. Use local Go1.27, fresh home/config/caches/tmp,
telemetry-off files, no network/download, <=60seconds per command. Go tests use
-timeout=20s; active native invocation uses -parallel=1 and -fuzztime=1x.

Write a deterministic byte target and representative finite seeds. Idempotence
and lack of adjacent output duplicates are necessary but do not reject global
uniqueness. Compare actual output with input run boundaries independently, and
check input preservation. Global uniqueness is false for valid aba; execute this
false-property control against correct production and retain its failure.

The reference seed order is fixed in the reviewed source. Execute seed#4 aaabbba,
seed#5 aabba and seed#6 aba before repair: each fails with expected aba versus
actual ab. Execute seed#7 ab, seed#8 aa and seed#9 ba: each passes. These are actual
finite reductions and shorter controls, not an automated minimizer or random
search. Fewer than3bytes cannot separate global uniqueness from adjacent-run
compaction: with0/1byte they coincide, and with2bytes equal values collapse in
both while unequal values survive in both. This argument is distinct from the
executed specimens and does not claim all shorter byte values were enumerated.

Write native corpus testdata/fuzz/FuzzCompact/separated-repeat and an ordinary
TestCompactRegression before changing production. Reproduce each independently
without active fuzzing; retain both failures. Repair CompactRuns, then rerun the
same corpus and ordinary regression, all public/grader/native seed tests and the
bounded native invocation. Retain its actual output; a1x invocation may run only
baseline seeds and proves neither random discovery nor long-campaign coverage.

The fixed additional specimens reject keep-duplicates, missing final run, reversed
result, empty output and input mutation with ordinary useful fuzz assertions.
Weak target passes the broken global-unique implementation and is method-rejected.
Log-only ordinary regression also passes broken source, while the independent
retained corpus still rejects it. Correct run-end oracle alternative passes.
Compiler errors, runtime failures and timeouts never count as semantic detection.

Exact manifest lists10fresh copies and47commands including version. Every source
replacement, seed-only/corpus/Test/full/fuzz argv, prediction, snapshot, status,
stdout/stderr and cleanup is retained. Source/script hashes and argv need sole
preexecution approval; actual evidence then needs separate method review. No
new runner/profile/AST/trace, automatic coverage/progress, or mastery claim.
