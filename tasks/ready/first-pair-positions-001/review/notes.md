# Content review

Independent application of `maps.key-comparability`. Source/design approved by /root/m2_2_review. Distinct contract: Supplied table Put/Lookup becomes selecting a stable comparable representation to classify repeated ordered string tuples by first occurrence; lossy encodings and pointer identity are observable errors.

| Requirement | Public assertion / reference learner test | Realistic defect |
| --- | --- | --- |
| ordered full-value key identity without lossy encoding | TestOrderedPairFirstPositions | concatenated-key, nul-delimited-key, first-component-only, second-component-only, unordered-key, address-key, rune-normalized-key |
| first occurrence including index zero and repeats | TestOrderedPairFirstPositions | latest-occurrence, zero-position-is-missing |
| input preservation and independent calls | TestOrderedPairFirstPositions / TestPairCallsAndEmptyInputs | erase-input, reuse-prior-call |

Diagnostics name the affected input/position and actual versus expected values or types. Compile errors/timeouts are not semantic detection. Equivalent implementations are accepted; no algorithm/source-shape, performance, full-skill or human-mastery claim.

All declared defects are checked again by existing task-check in disposable public-only fixtures with the grader removed; hidden detection alone is insufficient.

Independent content review APPROVE on 2026-10-22. Targeted technical checks and all applicable public-only fixtures passed. Exact content digests are recorded in the independent batch receipt; route/final diff and repository acceptance remain separate.

Review correction: all-byte equality includes invalid UTF-8. The retained public assertion contrasts distinct \xff, \xfe and UTF-8 U+FFFD strings in both components, plus repeats. The rune-normalized-key defect must fail this public assertion with input/actual/expected indices; the disposable public-only fixture provides automatic regression evidence. No Unicode-normalization or invalid-byte exclusion is allowed by this contract.

The invalid-byte distinction is intentionally asserted in public only. The grader does not repeat that case; removing this public assertion makes rune-normalized-key survive the existing complete task-check. A disposable assertion-removal experiment records that exact survivor, so CI coverage cannot silently fall back to the grader for this fixed gap.
