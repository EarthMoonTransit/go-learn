# I11 atomic-count test-writing independent actual review

Root/sole contract approved original159-word TASK. Official task-init, ADR0008 creation/digest/sources=[]; The initial author precheck was followed by sole independent actual acceptance; current metadata and independent-review.json record that decision. Existing interval, named-cases, diagnostic-helper and decision firsts are unchanged preparation only.

| Requirement | Learner reference named child / assertion | Required defect |
| --- | --- | --- |
| Empty updates preserve seed | empty-keeps-existing / full map | empty-clears-map |
| Existing and missing keys, unrelated state | existing-and-unrelated,missing-positive / map and exact error | ignore-missing,drop-unrelated,wrong-sign |
| Ordered repeated updates | repeated-key,positive-before-subtraction / full state | overwrite-repeat,reverse-order |
| Zero retained and missing zero inserted | existing-reaches-zero,missing-becomes-present-zero / full key set | reject-zero,delete-zero |
| Reject negative intermediate | first-negative,negative-then-recovery,missing-negative / exact sentinel and unchanged state | fresh-error,accept-negative,check-final-only |
| Reject late failure atomically | late-rejection-rolls-back / full original map | commit-as-you-go |
| Preserve delta input | helper on all success/error cases | mutate-deltas |

Thirteen finite wrong implementations; no panic, timeout, compiler/race failure credit. The starter's positive case passes public compatibility and misses defects. The strong learner suite passes compatibility and detects every required ordinary defect. Public/full roles differ for mutation-tests: correct-implementation compatibility does not establish sufficiency. Paired copied fixtures put each wrong operation in verification/reference with the strong suite to retain its ordinary named public-path diagnostic, without hidden tests. Full canonical checks assess the declared defect matrix.

## Actual method and controls

Learner chooses twelve named partitions; each serial child clones its own map/delta input. Shared checkCountOutcome accepts that child t, calls Helper before checking exact error, complete map/key presence, unchanged input and sentinel. Nonfatal Errorf leaves all independent rows runnable. This is not supplied RunNamed or AssertInt wiring. Method review checks actual table/cases/setup/t.Run/helper responsibility; no particular row/helper count, signature or library call is required.

Correct controls use pending-key operation, split marked helpers and another row order; all preserve ordinary behavior. Four method-wrong submitted suites keep all substantive checks but omit table, subtests, Helper marker or correct child scope. They are intended behavior-pass/method-reject controls, not required operation mutants. No AST extension or automatic source-method claim. No separate native selective-run or cleanup skill credit. Sole actual source/method/evidence review accepted on2026-10-23; see independent-review.json. Root promotion and final batch gates remain separate.
