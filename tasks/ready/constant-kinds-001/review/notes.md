# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Declare server as an untyped string constant "localhost" and port as a typed int constant 8080 in constants.go.
- Values returns server, port and the result of the constant division 3 / 2.0 as float64.
- Preserve the fractional result: integer-only division has a different value.
- Public checks inspect constant kinds with go/types and verify all three returned values; named aliases of int are allowed.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Actual constant kinds: untyped string and typed int, including aliases | TestDeclaredConstantKinds uses go/types over the full non-test package | typed-server, untyped-port | constant type actual/expected |
| Actual declared values equal localhost and 8080 | TestDeclaredConstantKinds compares go/constant values independently of Values | wrong-server-declaration, wrong-port-declaration | constant value actual/expected |
| Values returns required fields and fractional ratio | TestPublicCase01 and repeated grader | integer-division, wrong-port, wrong-server | returned name, port and ratio actual/expected |
| Accept standard-library use and fixed API references | Canonical reference uses fmt.Sprint(serverExpression) and portExpression; stdlib fixture | Correct alternative, not a defect | reference public/full pass |

For implementation/debugging tasks these are public checks. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and detection without prescribing test layout.

## Realistic required defect models

- integer-division
- typed-server
- untyped-port
- wrong-port
- wrong-server
- wrong-server-declaration
- wrong-port-declaration

## Source boundary

Substantive adaptation of `ardanlabs-gotraining` at commit `4ae5a46bfc665909a7cb29e11c18775330c63401`, selected operation in `topics/go/language/constants/exercises/template1/template1.go`. Exact consulted artifacts and SHA256 values were verified against the pinned local source; the Apache-2.0 notice and attribution were checked. Technical APIs retain selected source behavior with independent boundaries and defects. No reference-only or generated canonical material was reused. Private reference implementations and mutants remain outside learner payloads.

## Independent review

Whole-package source type-checking accepts stdlib imports and fixed API references. Constant kinds and values are inspected independently from returned values; canonical reference protects that accepted variation.

Reviewed current prompt, API/starter, public/reference assertions, grader or mutation layout, all 7 required defects, hints, route prerequisites and provenance. `.tmp/checks/constant-kinds-001.yaml` reports technical pass with 7/7 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch acceptance, lifecycle promotion and release binding remain separate. Automated fixtures are not human calibration or learner mastery.

The saved `review/stdlib-submission.txt` and canonical reference both use a standard-library import plus fixed API declarations. The author-provided `.tmp/checks/constant-kinds-stdlib.yaml` passes through the disposable checker. The Ardan template header mentions multiplication but its detailed body requests division; the attribution explicitly records following that detailed operation.
