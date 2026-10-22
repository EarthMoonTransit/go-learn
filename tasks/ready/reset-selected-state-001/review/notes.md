# Content review

Independent application of `types.zero-values`. Source/design approved by /root/m2_2_review. Distinct contract: zero-results-001 constructs fresh zeros; receiver-reset-001 resets one receiver. This operation selectively zeroes a slice of aggregates, preserves unselected owners and shared data, and counts selection rather than changed/nonzero states.

| Requirement | Public assertion / reference learner test | Realistic defect |
| --- | --- | --- |
| exact zeros for selected aggregate fields | TestSelectedZeroValuesAndSharing | range-copy-only, skip-final-state, keep-active-name, allocated-empty-containers, keep-target |
| unselected state and shared storage preservation | TestSelectedZeroValuesAndSharing / TestResetCountsAndRepeatedCalls | reset-unselected, clear-old-slice, clear-old-map, clear-old-pointee |
| selection, count and repeated reset | TestSelectedZeroValuesAndSharing / TestResetCountsAndRepeatedCalls | count-nonzero-only, erase-selection |

Diagnostics name the affected input/position and actual versus expected values or types. Compile errors/timeouts are not semantic detection. Equivalent implementations are accepted; no algorithm/source-shape, performance, full-skill or human-mastery claim.

All declared defects are checked again by existing task-check in disposable public-only fixtures with the grader removed; hidden detection alone is insufficient.

Independent content review APPROVE on 2026-10-22. Targeted technical checks and all applicable public-only fixtures passed. Exact content digests are recorded in the independent batch receipt; route/final diff and repository acceptance remain separate.
