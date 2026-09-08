# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Welcome returns "Welcome, <name>!".
- FormatAssignment returns three lines: Welcome(name), "Slot <slot>: <direction>, <distance> m.", and "Peer: <peer>.".
- Format slot as decimal with minimum width three and leading zeroes; format distance with one decimal place using fmt rounding.
- Preserve all supplied strings exactly, including Unicode, whitespace and empty strings. Separate lines with one newline and add no final newline.
- Slot is 0..9999; distance is finite from 0 through 10000.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Welcome and three exact lines without final newline | TestPublicCase01/02 and TestDirectWelcomeAndDirection | final-newline, missing-peer | quoted full record actual/expected |
| Minimum three-digit slot width and one decimal fmt rounding | TestPublicCase02/03/04 covers short and wide slots and rounding up | no-padding, precision-zero, truncate-distance | quoted boundary/wide-slot actual/expected |
| Exact Unicode, empty and whitespace strings | TestPublicCase03/04 plus direct Welcome and padded direction | trim-name, trim-direction | quoted Welcome and padded direction actual/expected |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- no-padding
- precision-zero
- truncate-distance
- trim-name
- final-newline
- missing-peer
- trim-direction

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, selected operation in `exercises/concept/party-robot/.docs/instructions.md`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the MIT notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

The source AssignTable operation is substantively retained through multiline width/precision formatting; direct Welcome and padded-direction checks cover exact string preservation independently.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 7 required defects, hints, route prerequisites and provenance. `.tmp/checks/formatted-assignment-001.yaml` reports technical pass with 7/7 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.
