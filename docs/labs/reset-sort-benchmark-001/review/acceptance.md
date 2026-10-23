# Actual benchmark evidence — author analysis submitted for independent review

The approved fixed recipe completed all six fresh copies and 17 commands. Every
actual expected observation matched, all before/after project sources matched,
and all six disposable roots were removed. No command timed out and no panic or
crash was counted. Exact argv, environment, version, outputs, exits, snapshots and
cleanup are in execution.json; source-manifest.json and preexecution-review.json
bind the bytes reviewed before execution. No source changed during the recipe.

Reference and valid clone-outside-timer alternative each produced 36 measurements:
six input cases, two operations and three repetitions, native 50x each. The table
shows all actual local reference ranges; values are observations, not grades.

| Reference case/operation | ns/op min–max (3 runs) | B/op | allocs/op |
| --- | ---: | ---: | ---: |
| n32/sorted/Insertion | 272.6–510 | [0] | [0] |
| n32/sorted/Merge | 1503–5554 | [256] | [1] |
| n32/reverse/Insertion | 462.5–736.6 | [0] | [0] |
| n32/reverse/Merge | 779.2–1176 | [256] | [1] |
| n32/duplicates/Insertion | 177.5–420.9 | [0] | [0] |
| n32/duplicates/Merge | 671.6–1537 | [256] | [1] |
| n256/sorted/Insertion | 184.2–302.5 | [0] | [0] |
| n256/sorted/Merge | 2910–9187 | [2048] | [1] |
| n256/reverse/Insertion | 13033–16979 | [0] | [0] |
| n256/reverse/Merge | 3462–4042 | [2048] | [1] |
| n256/duplicates/Insertion | 5829–6522 | [0] | [0] |
| n256/duplicates/Merge | 3698–5477 | [2048] | [1] |

Both methods preserve exact values/length, including nil, empty, signed and duplicate
cases and all their own generated benchmark cases. The test oracle uses slices.Sort
on a separate clone; benchmark setup/result assertions are outside the reference
measurement window. Reference benchmark lines 43–64 own the stop/reset/copy/check,
start/operation/stop, result validation and observation. Input templates remain
unchanged. Supplied sorting algorithm bodies earn no learner action credit.

Insertion reported 0 B/op and 0 allocs/op for these measured operations; Merge reported
one scratch allocation per operation, 256 B for 32 values and 2048 B for 256 values. The
reference working slice allocation is excluded before ResetTimer; the alternative
clones its working slice every iteration with timing stopped and still reports only
the measured operation's allocation. Its64-value Merge case reports 512 B/op and its
256-value case 2048 B/op. This does not make either entire experiment allocation-free:
untimed clones allocate, have different lifetimes and may affect GC/cache context.
Bytes describe storage volume and allocation counts describe events; neither alone
predicts timing. Both inplace operations still mutate the provided working slice.

Within this finite run, input order affects timing and relative operation costs;
see the per-case ranges instead of one aggregate winner. The small and larger sizes
and sorted/reverse/duplicate patterns vary actual work. Three short fixed 50x repeats
expose local variation, not statistical significance or a production-wide conclusion.
For example, at 256 values the reverse-input ranges were 13,033–16,979 ns/op for
Insertion and 3,462–4,042 ns/op for Merge; sorted-input Insertion was 184.2–302.5 ns/op
while Merge was 2,910–9,187 ns/op. These observations motivate the order matrix rather
than a single winner. Small-case variation and even the lower observed sorted 256
than sorted 32 Insertion times show why these brief measurements cannot establish
monotonic scaling. Other repository checks shared this host during the finite run;
this was not a controlled machine-isolation experiment. No additional repetition
or cleaner-environment result is invented.
Machine/toolchain context is retained in native stdout (Go 1.27.0 darwin/arm64).
The alternative's 64/256 cases and different duplicate pattern are valid; its numbers
need not equal the reference's 32/256 observations.

## Actual wrong and correct controls

| Specimen | Actual result | Review boundary |
| --- | --- | --- |
| clone-outside-timer | correctness pass; 36 measured rows | Valid fresh-clone state with setup/validation excluded; source/timing differs legitimately |
| no-restore | version/correctness pass; benchmark exit 1 | Real iteration 1 pre-state assertions name non-sorted case and actual/original state; copying only before loop is insufficient |
| timed-clone | all commands pass; 36 rows | Source starts timing before clone/check and stops after validation/observation; actual B/op and allocs/op include working-copy cost, so scope is wrong |
| single-order | all commands pass; 2 rows | Only sorted 32 and one repetition; absence of order/size/repeated comparison is a method/design failure |
| unequal-sort | version pass; correctness exit 1; no benchmark | Exact Merge result assertions show changed values; semantically unequal algorithms are not timed |

Timed-clone actually reports an extra working-slice allocation: Insertion uses
256/2048B and 1 alloc/op; Merge 512/4096B and 2 allocs/op for 32/256. These observations
support the line-level scope finding; numeric thresholds are not used to reject it.
No-restore's assertion can be deleted by a learner, so source review must still
verify per-iteration restoration. Neither command success nor supplied signatures
prove loop ownership, meaningful equivalence checks or representative design.

This author analysis records observed evidence only. Sole independent actual review
must decide acceptance of the three application slots; no automatic coverage,
mastery, additional sorting/component credit or release binding is asserted.

## Subsequent independent acceptance

The sole reviewer accepted actual source/method/evidence on2026-10-23.
See independent-review.json for exact file/command/snapshot and measured-output
checks. This accepts three manual applications as content evidence, with final
batch gates and delivery still required. No automatic coverage or mastery claim.
