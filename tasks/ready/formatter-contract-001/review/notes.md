# Review notes

## Decision

Returned to draft. Caller-defined formatting proves open interface dispatch,
but the starter already gives `Pair` the required method set. The learner does
not create implicit satisfaction, and two mutants cover string assembly rather
than the declared interface skill. Nil and typed-nil formatter behavior is also
undefined.

## Required repair

Reclassify this as focused interface dispatch or redesign the editable code so
method-set satisfaction is learner-owned. Replace unrelated mutants, remove
`interfaces.small-contracts` unless API design is observable, state the nil
precondition, shorten the estimate, and add exact source attribution plus
clean-room evidence.


## C2 repair, 2026-10-17

The previous fully solved Pair.Format body is now an unfinished compiling stub.
The learner implements both the method behavior and open interface dispatch.
This is a guided introduction: fixed-tests keeps public declarations compilable;
no claim is made that the scaffold measures unaided interface declaration design.
The implicit value method set and minimal consumer contract are introduced and
exercised through the direct method, Formatter binding, Pair dispatch and two
independent caller-defined implementations. Introduction completion alone cannot
establish mastery; separate independent practice/transfer remains required.

| Requirement | Public test | Wrong variant | Diagnostic |
| --- | --- | --- | --- |
| exact Pair formatting | TestPairAndImplicitContract | omit-right; swap-pair | Pair.Format arguments, actual/want |
| Pair through implicit Formatter | TestPairAndImplicitContract | omit-right; swap-pair | Render(Pair) actual/want |
| caller-defined types, no closed registry | TestCallerDefinedFormatters | concrete-only; stringify-value | caller formatter actual/want |
| exactly one interface call | TestCallerDefinedFormatters | double-format | Format calls actual/want |
| unchanged result, including whitespace/empty | TestCallerDefinedFormatters | trim-result; stringify-value | unchanged result actual/want |

Nil interface input and panicking Format are outside the declared domain. Source
review S2 verified only the exact existing concept metadata; no upstream body,
test or solution is used. Historical review remains above; fresh independent
content/assessment review and executable acceptance are pending.


## Independent content approval

`/root/clamp_e01_content_review` approved this exact content digest on
2026-10-17 after read-only TASK/API/public/grader/mutant/hint/prerequisite and
provenance review. Formatter approval is limited to the explicit P05 guided
introduction; unaided method-set design and independent mastery are not claimed.
P08 remains conditional on reviewed M1/placement before activation. Public/full
execution, normal promotion and ordered repository acceptance follow separately.


## Disposable public/full evidence

The existing learner verifier accepted the reference through both public and
full scopes. All 6 required compiling wrong variants failed public Check with
FailureCategory=test and useful actual/want diagnostics; no compile error,
timeout or hidden grader detection substitutes for this observation. Execution
used temporary learner payloads and fresh disposable checker copies, never the
authored source checkout or a real learner home. The one-off batch record is
`.tmp/pedagogical-c2/public-feedback.jsonl`. Persisted public-feedback harness
rows belong to the subsequent separate tooling batch, per docs/DEBUGGING.md;
that follow-up remains required before claiming the full rollout complete.
