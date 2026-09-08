# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- In months.go declare jan through dec in one const block using iota, with consecutive values 1 through 12.
- The supplied API exports these as Jan through Dec and requires actual constant expressions.
- GetMonth returns its argument unchanged; callers pass one of the declared month constants.
- Public checks verify all twelve values and the use of iota in the month constant block.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| All twelve actual constants are consecutive 1..12 | TestPublicCase01 and fixed constant-expression consumers | zero-start, gap-start | month ordinal, declared constant and result |
| GetMonth returns the supplied month | TestPublicCase01 all twelve and final-month grader | return-zero, shift-return | month constant and returned result |
| One month constant block uses iota | TestMonthConstBlock reads editable source in disposable workspace | manual-enumeration | month constants must share one const block using iota |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- zero-start
- gap-start
- return-zero
- shift-return
- manual-enumeration

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, selected operation in `exercises/concept/savings-account/.docs/instructions.md`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the MIT notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

Fixed API constant-expression consumers and public AST checks make the actual editable iota declaration contract observable; all twelve values are checked.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 5 required defects, hints, route prerequisites and provenance. `.tmp/checks/month-enumeration-001.yaml` reports technical pass with 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.
