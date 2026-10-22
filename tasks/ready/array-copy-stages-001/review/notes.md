# I6 copy-method author evidence

Contract/allocation approved by sole /root/reviewer in i6-split-review.json. Actual content, provenance and method acceptance independently approved by /root/reviewer2026-10-22; final batch gates remain separate.

| Requirement | Public assertion | Wrong variant/control | Useful diagnostic or boundary |
| --- | --- | --- | --- |
| Original is unchanged by assignment mutation | TestPublicDistinctAssignmentAndCallStages / source assertion | alias-original | source actual/want after modifying only the assigned array |
| Assigned index 0 changes before the call | assigned assertion in every public input | snapshot-before-update | assigned before call actual/want with only index 0 replaced |
| Helper changes its local index 1 and returns it | insideCall assertion | callee-wrong-index; discard-callee-result | insideCall actual/want with callee index 1 replaced |
| Caller survives the value call unchanged | afterCall assertion | pointer-call | changing callee index 1 must not change its caller |
| Three-element length contract; preserve final element and zero/negative values | TestPublicZeroReplacementAndLastElement; TestPublicZeroCalleeValue | two-element-model | source/assigned/insideCall/afterCall actual/want includes third element |
| Learner owns literal, copy, helper/call and two predictions | actual source/method inspection, separate from behavior | correct alternative passes; fabricated-stage control also passes behavior but is method-wrong | no supplied return-type or passing-output credit for the missing operations |

The reference prediction file is learner-owned implementation evidence; its starter is package-only. Fixed public/grader tests are separate checker-owned paths and never depend on declarations in learner tests. Public-only copied-fixture runs replace both grader and all submitted prediction test files with package-only files. These runs establish fixed feedback independently of submitted assertions. Canonical checks alone do not measure submitted-test mutation quality; source review checks the actual predictions.

All required variants compile and terminate with ordinary assertions. No AST/source analysis, reflection checker, subprocess from authored Go, time, panic or compile-error detection is used. Integer inputs are finite and there is no concurrent execution. Grader corpora stop at their first mismatch to preserve bounded useful output.

Controls are retained as inert source under .tmp/v1-repairs-06/copies-*: a correct private naming/layout alternative must pass full and public-only checks; fabricated observations also pass behavior and are explicitly method-rejected because they omit actual copy operations. No extra mutant or automatic method claim is inferred from the passing control. Actual copy/helper/field reasoning and prediction assertions require the independent review.

No old task/application/route bytes are modified. Source identity and release binding follow ADR 0008; no external code, tests, solutions or teaching text were copied or closely paraphrased.

Independent actual method acceptance: Actual learner-owned array literal length, assignment, [3]int helper/value call, stage snapshots and prediction assertions accepted. Public-only probes exclude learner tests;6ordinary diagnostics, correct alternatives and passing fabricated-stage method-rejected control retained. Repairs first only; old independent application remains unchanged.
