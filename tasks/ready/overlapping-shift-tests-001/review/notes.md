# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return the original final element, shift elements index..len(values)-2 one place right, then write zero at index.
- Modify the supplied storage so aliases observe the shift; preserve elements before index.
- The slice is nonempty and 0 <= index < len(values). Length is unchanged; signed and repeated values are valid.
- The source and destination overlap. Tests must distinguish a correct shift from repeated overwriting while copying forward.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Overlapping move right and zero insertion | TestContractCase01 and Case02 use distinct values and multiple indices | forward-overwrite, missing-zero, ignore-index | shift contents and removed element |
| Return original last and update supplied storage | TestContractCase01 verifies slice and alias | return-new-last, detached-output | actual shifted slice and original removed value |
| Singleton/final index must still clear the selected slot | TestContractCase02 includes singleton and last index | final-index-skipped | index, actual and expected slice/result |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- forward-overwrite
- missing-zero
- return-new-last
- ignore-index
- detached-output
- final-index-skipped

## Source boundary

Original task under ADR0008. Project identity matches `config/project-identity.yaml`; creation evidence and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest. No external code, tests or text are claimed or reused.

## Independent review

A final-index defect keeps the supplied starter insufficient, while the strong suite covers both overlapping movement and a singleton/final slot.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 6 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/overlapping-shift-tests-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
