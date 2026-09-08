# Content review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

## Requirements

- Return (0,false) for a nil Coder interface or an interface holding a nil *Result pointer; do not call its method.
- For a non-nil *Result return its Code and true, including zero and negative codes.
- For any other implementing type call CodeValue once and return its result with true. Such values have safe method receivers.
- Preserve the pointed Result. An interface holding a typed nil pointer is not itself nil.

## Requirement, assertion and defect mapping

| Requirement | Public/reference assertion | Required defect | Diagnostic |
|---|---|---|---|
| Nil interface and typed nil *Result return false without method calls | TestPublicCase01 | interface-nil-only | nil dynamic type and result; unsafe method call panics |
| Non-nil *Result returns signed/zero Code and is preserved | TestPublicCase02 | all-result-invalid, zero-absent, mutate-result | returned code/bool and stored code |
| Other implementers are called exactly once | TestPublicCase03 uses stateful otherCoder | other-rejected, double-call | returned code/bool and call count |

For implementation/debugging tasks these are public assertions. For test-writing tasks these are strong learner reference assertions; the mutation profile checks reference compatibility and required defect detection, without requiring a particular test layout.

## Realistic required defect models

- interface-nil-only
- all-result-invalid
- zero-absent
- other-rejected
- double-call
- mutate-result

## Source boundary

Original task under ADR0008. Project identity matches `config/project-identity.yaml`; creation evidence and content digest are recorded, external sources are empty, and commit binding is deferred to the actual release manifest. No external code, tests or text are claimed or reused.

## Independent review

Nil-interface and typed-nil handling are distinguished from zero-code absence, arbitrary implementation rejection and receiver mutation.

Reviewed current English requirements, API/starter, public or learner reference assertions, grader/mutation layout, all 6 required defects, hints, route prerequisites and provenance. The author-provided `.tmp/checks/typed-nil-result-001.yaml` reports technical pass and 6/6 required mutants detected. No authored code was executed by the reviewer; no real learner state was accessed. Full batch gates, lifecycle promotion and the separate public-feedback tooling matrix remain separate. Automated fixtures do not establish learner mastery or human calibration.
