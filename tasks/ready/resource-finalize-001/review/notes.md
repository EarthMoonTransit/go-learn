# Review notes

## Decision

Returned to draft. The task mixes cleanup ownership, `defer`, named-result
mutation, the `io.Writer` short-write contract, wrapping, and an `errors.Join`
policy. The grader does not verify the promised operation context, an empty
write, or all simultaneous write/close failure paths. The final hints reveal the
reference's named-return and join strategy.

## Required repair

Split basic cleanup/order from advanced multi-error preservation, weaken hints
to ownership and error-policy invariants, and add focused context and combined
failure evidence. Original-content promotion remains blocked by the AGENTS.md
versus ADR 0008 provenance conflict.


## C2 repair, 2026-10-17

P10 explicitly retains this existing combined-error ownership contract. P06/P07,
error wrapping/Is and reader/writer doubles prepare its mechanisms before the
planned transfer placement. The historical request to split does not create a
second task: the repair preserves the user-requested existing policy, replaces
solution-revealing hints and adds all missing evidence. Current AGENTS rule 2
explicitly admits ADR 0008 originals; the old provenance-policy conflict is no
longer applicable. Actual content digest is refreshed; creation/source identity
and real future release binding remain intact.

| Requirement | Public assertion | Wrong variant | Diagnostic |
| --- | --- | --- | --- |
| one Write(data), including empty input | success/empty plus captured bytes | retry-write; wrong-data; empty-skips-calls | calls/order or data actual/want |
| one Close after any attempt | success/write-only/both/empty/panic | close-before-write; skip-close-on-write-error | calls/order actual/want |
| write operation context and errors.Is | write-only/both/empty-and-errors | lose-write-context; overwrite-write-error | write context or identity actual/want |
| short write with nil write error | short/short-and-close | ignore-short-write | short-write identity actual/want |
| close-only failure | close-only | ignore-close-error | close identity actual/want |
| simultaneous write/short and close failures | both/short-and-close/empty-and-errors; distinct wrapped causes with equal text | overwrite-write-error; ignore-close-error; deduplicate-error-text | both wrapper/cause identities actual/want |
| no retry/recovery | event matrix; panic identity | retry-write; recover-write-panic | calls/order or original panic actual/want |

The context requirement names the operation but imposes no exact prefix, wrapper
type or formatting strategy. Context matching is case-insensitive; classifications
use errors.Is. Public cases cover every promised branch and grader expands the
same domain across sizes/counts and all write/close combinations. No hidden API,
allocation, nilness, buffer ownership or error-order guarantee is introduced.
The defensive Write domain explicitly admits short counts with nil error. The
equal-text public case uses distinct wrapped errors and distinct underlying
causes, requiring all four errors.Is identities; comparing Error strings cannot
substitute for preserving both failures. Fresh independent review and execution
remain pending.


## Independent content approval

`/root/clamp_e01_content_review` approved this exact content digest on
2026-10-17 after read-only TASK/API/public/grader/mutant/hint/prerequisite and
provenance review. Formatter approval is limited to the explicit P05 guided
introduction; unaided method-set design and independent mastery are not claimed.
P08 remains conditional on reviewed M1/placement before activation. Public/full
execution, normal promotion and ordered repository acceptance follow separately.


## Disposable public/full evidence

The existing learner verifier accepted the reference through both public and
full scopes. All 11 required compiling wrong variants failed public Check with
FailureCategory=test and useful actual/want diagnostics; no compile error,
timeout or hidden grader detection substitutes for this observation. Execution
used temporary learner payloads and fresh disposable checker copies, never the
authored source checkout or a real learner home. The one-off batch record is
`.tmp/pedagogical-c2/public-feedback.jsonl`. Persisted public-feedback harness
rows belong to the subsequent separate tooling batch, per docs/DEBUGGING.md;
that follow-up remains required before claiming the full rollout complete.
