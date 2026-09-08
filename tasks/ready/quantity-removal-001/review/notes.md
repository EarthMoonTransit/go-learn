# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return false without changing either map if item or unit is missing or subtraction would be negative.
- Otherwise subtract the unit quantity, delete the item when its new quantity is zero, and return true.
- Item quantities are nonnegative; unit quantities are positive. Maps do not share storage; nil maps contain no entries.
- Preserve unrelated items and all unit entries. Keys are exact strings, including empty strings.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Subtract or delete exactly at zero | TestPublicCase01 repeats removal | keep-zero, reject-exact | subtract contents / exact removal rejected / zero key remains |
| Missing item/unit or insufficient quantity preserves both maps | TestPublicCase02 snapshots maps | ignore-insufficient, unknown-unit-success | rejected keys and changed map contents |
| Preserve unrelated items/units; exact and empty keys | TestPublicCase01/03 and nil-map grader | clear-unrelated, mutate-units, empty-key-rejected | maps changed / empty keys / case-normalized key accepted |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- keep-zero
- reject-exact
- ignore-insufficient
- unknown-unit-success
- clear-unrelated
- mutate-units
- empty-key-rejected

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, selected operation in `exercises/concept/gross-store/.docs/instructions.md`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the MIT notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

Removal-on-zero, unsuccessful no-change behavior and preservation of the unit table are distinct from earlier quantity insertion and lookup operations.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 7 required defects, hints, route prerequisites and provenance. `.tmp/checks/quantity-removal-001.yaml` reports technical pass with 7/7 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.
