# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- English.Speak returns "Hello World"; (*Chinese).Speak returns "你好世界".
- SayHello calls Speak once on any supplied non-nil Speaker and returns its exact result.
- BuiltinGreetings returns English then Chinese output by using values or pointers that satisfy Speaker.
- Keep the declared receiver forms: English values and pointers qualify; Chinese requires a pointer. Calls are independent.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Declared value/pointer method forms and bilingual outputs | TestPublicCase01 assigns English value/pointer and Chinese pointer to Speaker | same-language | quoted speaker actual/expected |
| Dispatch any Speaker exactly once using current exact result | TestPublicCase02 custom mutable speaker | known-only, double-call, trim-result | dispatch text and call count |
| Builtin output order | TestPublicCase01 BuiltinGreetings | reverse-builtins | builtin array |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection without prescribing test layout.

## Realistic required defect models

- same-language
- known-only
- double-call
- trim-result
- reverse-builtins

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/interfaces/exercises/template1/template1.go`. Exact declared artifacts and SHA256 values were checked against the pinned local checkout; Apache-2.0 attribution and notice were verified. Source narrative/process output is replaced by a technical API and observable checks, with independent boundaries and defects. No reference-only or generated canonical material was reused. Private references and mutants remain excluded from learner payloads.

## Independent review

The detailed source Chinese pointer receiver is retained and its header discrepancy documented; caller-defined speakers test open interface dispatch.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 5 required defects, hints, route prerequisites and provenance. `.tmp/checks/receiver-speakers-001.yaml` reports technical pass and 5/5 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion, release binding and separate public-feedback regressions remain independent requirements. Automated fixtures are not human calibration or learner mastery.
