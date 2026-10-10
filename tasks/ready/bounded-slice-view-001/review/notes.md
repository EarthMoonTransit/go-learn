# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return the elements in in[start:end] with length and capacity both end-start.
- Existing view elements share input storage, so changing one changes the corresponding input element.
- Appending an element to the returned view must allocate independent storage and preserve input elements outside the view.
- Bounds satisfy 0 <= start <= end <= len(in). Empty ranges and nil input are valid; do not mutate input during the call.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Exact window length/capacity and sharing of existing elements | TestPublicCase01 changes a returned element after verifying values and bounds | two-index, detached, ignore-start | view length/capacity and detached input |
| Preserve entire input during the call, including outside the view | TestPublicCase01 compares a complete snapshot before mutating the view | mutate-input, prefix-cleared | call changed input actual/expected |
| Appending cannot overwrite outside the view; empty ranges have zero capacity | TestPublicCase01 append and Case02 start/middle/end empty windows; nil grader | two-index, empty-tail-cap | append overwrote outside / empty length/capacity |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- two-index
- detached
- ignore-start
- empty-tail-cap
- mutate-input
- prefix-cleared

## Source boundary

Original task under ADR0008. Project identity matches `config/project-identity.yaml`; creation evidence and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest. No external code, tests or text are claimed or reused.

## Independent review

The complete input snapshot closes the outside-window mutation gap; shared element updates and isolated append capacity remain distinct requirements.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 6 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/bounded-slice-view-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.


P4C: nil public contract

The existing contract permits nil input with valid zero bounds, requiring length
and capacity zero. It does not require a nil representation: a non-nil empty
zero-capacity slice remains valid. TASK, starter, reference, hints and exported
workspace files are unchanged. The new nil-spare-cap defect returns spare
capacity only for nil input and must fail public Check with actual/expected
length and capacity. P4T provides the persistent Service.Check regression and
exact directional content revision compatibility, without rewriting learner files.

| Requirement | Public assertion | Defect | Diagnostic |
| --- | --- | --- | --- |
| elements/length/capacity | TestPublicCase01 | two-index, ignore-start | view values/len/cap |
| shared elements | TestPublicCase01 | detached | view detached |
| append isolation | TestPublicCase01 | two-index | append overwrote outside |
| unchanged input | TestPublicCase01 | mutate-input, prefix-cleared | call changed input |
| empty bounds | TestPublicCase02 | empty-tail-cap | empty bounds len/cap/input |
| nil with zero bounds | TestPublicNilView | nil-spare-cap | len=0 cap=1; want len=0 cap=0 |

Reference and existing learner alternatives remain subject to public/full gates.
This preparation correction adds zero ready IDs and does not change route numbers.
