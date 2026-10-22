# Executed reference review: join-benchmark-001

Date: 2026-10-22. Author inspection and independent actual content/method review by /root/reviewer: approved. No open findings; batch gates and exact-head CI remain separate.
[Actual normalized execution](execution.json) retains every command, status, output, version, source snapshot and cleanup result. All 4 independent roots were removed; no command timed out. Content digest is bound by provenance.json. Private paths are normalized as $DISPOSABLE/$LOCAL_GO_BIN, not invented execution. Raw receipts are retained in .tmp/v1-introductions-02/native-execution.json.

## Reference measurements and interpretation

Actual context: Go1.27.0, darwin/arm64, Apple M4 Max; three native runs per case with 20ms benchtime. These short repeated measurements are a teaching specimen, not a statistically established production speedup. There is no machine-dependent grade.

| Fixture / implementation | ns/op observed min–max | B/op in each run | allocs/op in each run |
| --- | ---: | ---: | ---: |
| empty/Concat | 1.299–1.39 | 0, 0, 0 | 0, 0, 0 |
| empty/Builder | 1.653–1.697 | 0, 0, 0 | 0, 0, 0 |
| shortASCII/Concat | 59.17–61.07 | 24, 24, 24 | 4, 4, 4 |
| shortASCII/Builder | 15.73–15.94 | 8, 8, 8 | 1, 1, 1 |
| longASCII/Concat | 296.8–310.4 | 3488, 3488, 3488 | 4, 4, 4 |
| longASCII/Builder | 323.1–330.9 | 4032, 4032, 4032 | 4, 4, 4 |
| multibyte/Concat | 79.33–80.72 | 435, 435, 435 | 4, 4, 4 |
| multibyte/Builder | 31.12–31.55 | 152, 152, 152 | 2, 2, 2 |

1. **Timing action:** actual reference constructs the four input fixtures before timed loops, checks each result against strings.Join, calls ReportAllocs/ResetTimer, uses dynamic b.N and stores each result in package Result. TestEquivalentJoins also checks full byte semantics and unchanged input. Provided Concat/Builder code is excluded from learner credit.
2. **Allocation action:** B/op is allocated bytes per operation, allocs/op counts allocation events; neither is live heap size or elapsed time. Concat builds intermediate strings on repeated appends; Builder may grow its backing buffer. On this sample shortASCII changes24B/4events→8B/1event, while longASCII Builder allocates4032B versus Concat3488B, both4events. Source growth and runtime allocation rounding explain why a builder is not automatically cheaper. The returned string must remain valid after the local builder ends; Result retains the last output, while earlier results can become garbage. Input storage is created outside the measured operation. Lower event count alone is not a correctness or universal speed result.
3. **Design action:** empty is a control; shortASCII, longASCII and multibyte exercise different byte lengths/patterns. Every pair has equivalent output and three observations. Builder is faster in this short/multibyte specimen but slower for this longASCII fixture; report these local ranges rather than a universal winner. More cases, longer measurements and controlled comparisons would be needed for stronger conclusions.

## Executed wrong specimens

- setup-inside passes correctness and native benchmark execution but adds `append([]string(nil), c.values...)` to every timed iteration. All nonempty cases show48 extraB/op and one extra allocation in this run; source inspection establishes the unwanted setup work, independent of any numerical threshold. Method rejected.
- missing-separator fails a Builder actual/want correctness assertion. It was not benchmarked while unequal; the restored implementation passes.
- single-tiny executes one two-element input once. Those numbers describe only that measurement; it fails the required representative/repeated experimental design. Its single tiny result even illustrates that fewer allocations can coexist with more B/op, so those metrics are not interchangeable.

Three distinct first actions are accepted only after independent code/results/analysis review. None is the independent application encounter for the same skill. Automatic coverage/progress is unchanged.
