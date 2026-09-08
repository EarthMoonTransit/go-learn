# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Temperature units format as "°C" or "°F"; speed units as "km/h" or "mph". Only declared units are supplied.
- Temperature and Speed format their integer followed by one space and the formatted unit.
- Reading formats "<Location>: <Temperature>, Wind <Direction> at <Speed>, <Humidity>% Humidity".
- String and fmt.Sprint produce identical results, preserving text, zero and negative integers.
- Implement all five value methods without recursive receiver formatting.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Unit labels and exact spacing | TestPublicCase01 direct String and fmt.Sprint | celsius-always, speed-unit-always, no-unit-space | quoted String/Sprint actual and expected |
| Signed and zero magnitudes | TestPublicCase01/03 | absolute-temperature | quoted numeric text and unit |
| Compose exact nested Reading text, preserving fields | TestPublicCase02/03 direct and formatted output | no-humidity-percent, direction-erased, trim-location | quoted reading/boundary actual and expected |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- celsius-always
- speed-unit-always
- no-unit-space
- absolute-temperature
- no-humidity-percent
- direction-erased
- trim-location

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, operation in `exercises/concept/meteorology/.docs/instructions.md`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

Five concrete Stringer behaviors and nested composition are independently checked through direct calls and fmt formatting.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 7 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/nested-stringers-001.yaml` reports technical pass and 7/7 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
