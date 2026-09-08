# Content audit: C02–C04

Scope: 30 new tasks across three content batches; checkpoint 52 learner-visible
ready tasks in the numbered route, excluding three legacy ready bank packages
(55 ready packages in total). Independently reviewed after recomputing the counts.
Automated verification is not calibration from real learner solutions.

## Reuse and diversity

The three batches add 23 substantive Exercism adaptations and seven original
Unicode tasks: 15 implementation, nine test-writing and six debugging tasks.
Each batch meets the 5/3/2 delivery target. Existing baseline tasks bring the
course to 26 implementation, 11 test-writing and 15 debugging tasks.

Operations were reviewed separately from naming: rune-marker/replacement/limit
have different observable behavior; boolean policies form one four-function
component; slice retrieval/replacement/removal/prepend assess distinct effects;
column occupancy, rank occupancy and total stored slots distinguish selected
values, map iteration and nested lengths. Quantity/presence reinforces the
existing fallback lesson with an explicit two-result API; it is practice, not
a new skill introduction. Pointer construction and optional reads introduce
lifetime and nil behavior. No cosmetic name change is counted as a new source
operation.

## Current primary-domain distribution

Small domains and the baseline are not yet a complete coverage release. Rows
over 50% in one activity require balancing before the mandatory coverage gate.
They are not reported as satisfying that gate. Upcoming selections must favor
the missing activities rather than silently weakening the policy.

| Domain | Implementation | Test-writing | Debugging | Highest share |
|---|---:|---:|---:|---:|
| arrays | 1 | 0 | 0 | 100% |
| control-flow | 3 | 1 | 5 | 56% |
| expressions | 4 | 1 | 0 | 80% |
| functions | 0 | 1 | 1 | 50% |
| maps | 3 | 2 | 0 | 60% |
| pointers | 1 | 1 | 0 | 50% |
| slices | 3 | 1 | 3 | 43% |
| strings | 3 | 1 | 2 | 50% |
| types | 4 | 0 | 2 | 67% |
| unicode | 4 | 3 | 2 | 44% |

## Test strength and public feedback

Independent task review found and corrected distinct UTF-8 decoder omissions,
inclusive OR versus XOR, exact whitespace/key normalization, and cached pointer
reads. Narrow saved test submissions for Unicode and pointer reads remain
insufficient while the full reference submissions kill every required defect.
Invalid bytes and exact keys have quoted diagnostics. Constraints now state
that the two maps in unit addition do not alias. The constant task uses the
existing workspace profile to require a constant expression without restricting
the learner to one spelling of integer-valued 40.

A separate tooling batch has prepared public-only regression checks for the reviewed
boolean, token and map-key defects. This closes the gap where a hidden test
would still kill the defect after an accidental public assertion removal.
The regression batch is pending its own local/PR acceptance; this audit does
not claim it has merged. No checker contract, source policy or task limit is relaxed.

## Next adjustments

Continue to 150 without a checkpoint stop. Build structs, pointers and
functions in the next batch, then increase test-writing/debugging opportunities
in domains with excess implementation share. Keep domain deficits visible in
each subsequent audit; new 5/3/2 batches alone do not ensure domain balance.
Maintain source-specific provenance and distinct practice/transfer evidence.
No learner feedback has arrived during this preparation, so no automated
recalibration or claimed mastery is recorded.

Independent review confirmed the counts, domain rows and semantic distinctions.
It approves this three-batch audit only, not future content or full coverage.
