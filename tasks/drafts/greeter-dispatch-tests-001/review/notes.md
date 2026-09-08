# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return "I can speak " + g.LanguageName() + ": " + g.Greet(name).
- The non-nil Greeter may hold any implementing concrete type and state; dispatch both methods through its current dynamic value.
- Forward name exactly, including empty strings, Unicode and spaces. Preserve returned language and greeting text exactly.
- Composition must work for arbitrary implementing types and configurations, rather than recognizing a fixed list of languages.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Compose both interface method results exactly | TestContractCase01/02 with configurable greeting and echoGreeting | fixed-language, fixed-greeting, missing-separator | quoted actual/expected complete result |
| Preserve empty, Unicode and whitespace arguments/results | TestContractCase02 | trim-name, empty-name, trim-language | quoted dispatch result showing whitespace |
| Arbitrary implementing types/configurations | Two concrete implementations plus multiple configurations of one type in TestContractCase02 | fixed-language, fixed-greeting | dispatch actual/expected; no test-layout requirement |

For implementation/debugging tasks these are public checks. For test-writing tasks these are the strong learner reference suite; the mutation profile checks reference compatibility and detection of required defects, without imposing table layout or helper organization.

## Realistic required defect models

- fixed-language
- fixed-greeting
- trim-name
- empty-name
- missing-separator
- trim-language

## Source boundary

Substantive adaptation of `exercism-go` at commit `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, operation in `exercises/concept/airport-robot/.docs/instructions.md`. Exact consulted artifact paths and SHA256 values are recorded in `provenance.yaml` and were checked against the pinned local source. The upstream MIT notice and attribution were verified. Observable source behavior is retained with technical APIs, independent boundaries and defect models; no reference-only or generated canonical material was used. Private reference implementations and mutants are excluded from learner payloads.

## Independent review

The contract requires arbitrary interface behavior without an ungraded test-layout rule; custom implementations exercise composition and exact whitespace.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 6 required defect implementations, hints, provenance and route prerequisites. The author-provided `.tmp/checks/greeter-dispatch-tests-001.yaml` reports technical pass and 6/6 detected required mutants. Review was read-only for task content: no authored code was executed by the reviewer and no real learner state was accessed. Full batch gates, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or evidence of learner mastery.
