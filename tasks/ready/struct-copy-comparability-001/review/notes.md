# I6 copy-method author evidence

Contract/allocation approved by sole /root/reviewer in i6-split-review.json. Actual content, provenance and method acceptance independently approved by /root/reviewer2026-10-22; final batch gates remain separate.

| Requirement | Public assertion | Wrong variant/control | Useful diagnostic or boundary |
| --- | --- | --- | --- |
| Scalar and inline-array source independence | TestPublicScalarInlineAndPointeeStages / source Count/Inline | alias-record | source Count/Inline actual/want after modifying only the copy |
| Exact copy scalar and inline updates | copied Count and copied Inline assertions | skip-count-update; skip-inline-update | copied Count/Inline actual/want |
| Shared pointee observes copy-side write | shared pointee observations assertion | deep-copy-target; skip-pointee-update | both observations actual/want |
| Copied Target retains pointer identity | sameTarget assertion | detach-target | sameTarget=false; want true |
| Equality includes Count and both Inline elements | TestPublicEqualityChecksEveryField | equality-ignore-count; equality-ignore-inline-first; equality-ignore-inline-last | different field: EqualRecord actual/want |
| Pointer equality uses identity; nil is valid | equal pointed values / one nil / other nil / matching nil cases | equality-pointee-values; nil-records-unequal | equal pointed values do not make different Target pointers equal; matching nil Targets actual/want |
| Comparability considers all field types | TestPublicAllFieldComparability | all-structs-comparable; pointers-not-comparable | Record int/array/pointer comparable; SliceRecord.Extra slice noncomparable even nil/empty |
| Actual shallow assignment, comparison decisions and authored prediction assertions | separate source/method review | alternative fieldwise equality valid; fabricated-stage control behavior-passing but method-wrong | no supplied signature/type or predetermined booleans alone establish ownership/analysis |

The reference prediction file is learner-owned implementation evidence; its starter is package-only. Fixed public/grader tests are separate checker-owned paths and never depend on declarations in learner tests. Public-only copied-fixture runs replace both grader and all submitted prediction test files with package-only files. These runs establish fixed feedback independently of submitted assertions. Canonical checks alone do not measure submitted-test mutation quality; source review checks the actual predictions.

All required variants compile and terminate with ordinary assertions. No AST/source analysis, reflection checker, subprocess from authored Go, time, panic or compile-error detection is used. Integer inputs are finite and there is no concurrent execution. Grader corpora stop at their first mismatch to preserve bounded useful output.

Controls are retained as inert source under .tmp/v1-repairs-06/copies-*: a correct private naming/layout alternative must pass full and public-only checks; fabricated observations also pass behavior and are explicitly method-rejected because they omit actual copy operations. No extra mutant or automatic method claim is inferred from the passing control. Actual copy/helper/field reasoning and prediction assertions require the independent review.

No old task/application/route bytes are modified. Source identity and release binding follow ADR 0008; no external code, tests, solutions or teaching text were copied or closely paraphrased.

Independent actual method acceptance: Actual learner-owned struct assignment, inline copy versus shared pointee, complete-field equality/comparability and prediction assertions accepted.13public-only ordinary diagnostics plus valid alternatives and passing fabricated-stage method-rejected control retained. No invalid comparison compiler/panic evidence. Repairs first only; existing independent application unchanged.
