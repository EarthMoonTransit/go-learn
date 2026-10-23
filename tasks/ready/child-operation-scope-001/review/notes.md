# I7 author precheck

Approved contract section SHA256 665624d4b40cd062ef8c1ee5012f54ea13e1d1881e7975a1e6b8db13f3ae005d.
TASK is unchanged from that approved section: 135 checker words, five requirements.
Earlier author metadata enabled targeted disposable checks; current sole-reviewer actual acceptance is recorded below.

| Requirement | Public assertion | Ordinary wrong specimen | Diagnostic |
| --- | --- | --- | --- |
| Empty result and all ordered callback pairs | TestPublicScopedSequence | phantom-empty, stop-after-error, skip-first, double-call, reverse-order, drop-value, replace-error | result count, callback order/count, indexed Value and Err identity |
| Incoming parent value/deadline/cancellation | sequence deadline/canceled rows and TestPublicParentCancellationWaitsForCleanup | background-parent, nil-child, skip-canceled | lost parent value/deadline/cancellation, nil context, skipped callbacks |
| Fresh sibling scope, independent from parent | live sequence rows | return-parent, reuse-child, chain-scopes | fresh child; child canceled before invocation |
| Release previous before next; all children released on return | sequence sibling/success/error rows | defer-batch-cancels, no-release, release-only-on-error, cancel-before-call | previous child still live; remains live after return; canceled before invocation |
| Synchronous cleanup and continuation after parent cancellation | TestPublicParentCancellationWaitsForCleanup | return-before-cleanup, skip-canceled | returned before cleanup; next callback entered early; joined result count |
| Separate calls | TestPublicSeparateCalls plus grader matrix | reuse-child also caught per call; cached state is actual review concern | ordered results and fresh live scopes |

Eighteen required mutants are finite ordinary implementation errors, not compile/panic/race/timeout probes. All are expected to produce useful public assertions. Grader extends counts 0–4, every error position, live/canceled parents and deadlines using the same declared policy. Nil and empty results are equivalent for empty input. Error values are compared by identity. Tests never compare context interfaces by ==; Done identity is inspected only while parent/scopes are live and Done has been obtained before release. Pre-canceled children may share closed Done channels.

The held callback has independent rescue and cleanup gates. Assertions run before releasing those gates; then the fixture cancels its parent, releases both, settles and collects the wrapper result. Callback events use bounded buffered channels (32 versus at most 8 invocations in finite mutants); sequence accounting is mutex-protected. No mutex is held during synctest.Wait. A nil context records a diagnostic and follows rescue. No Fatal can bypass release. The asynchronous return defect finishes after rescue; tests do not need leaked or permanently blocked workers as evidence. Cleanup refers to the callback's synchronous completion, not hidden asynchronous work that the callback might invent.

## Actual method review

The learner writes the private helper signature with context first, chooses remaining parameters/result shape, forwards the actual original parent into it, creates a fresh child there, forwards child into operation, cancels after callback return, then collects each pair. No scope loop/helper/cancel/forwarding logic is supplied. The application differs from the single-call propagate-context first by actual learner-authored helper API and two call boundaries; it differs from cancel-worker by sequential sibling ownership/release instead of repairing the parent for a supplied worker.

Correct controls: loop-created WithCancelCause children with context-first invocation helper, and helper-local defer cancellation. Both must pass public/full/race. The separate inline-loop control intentionally satisfies behavioral requirements but omits the explicitly required helper: preserve its passing behavioral results and reject its method, without inventing an AST guard. No helper name, private storage or constructor spelling restriction applies. Public tests are not a universal resource-leak proof. Only derivation and propagation-api independent applications are proposed; no signature-only, cancellation/goroutine/slice/error/struct credit.

## Provenance and acceptance

Original ADR 0008, exact creation record/digest, external sources[]. Preparation remains unchanged at 0dcd72255ddf347197cc94f6d3ba40da6e8cf36c. Current I7 base is root-managed. Authored code executes only via canonical task-check disposable workspaces, including temporary copied public/control fixtures. Source/evidence/method acceptance belongs to the sole reviewer; promotion and shared route changes belong to root.

## Independent actual acceptance

Sole independent reviewer /root/reviewer approved actual source, executed evidence, provenance, public feedback and finite method on2026-10-22. Current decision is accepted for official promotion; final batch local gates, final diff review, exactHEADvalidate and merge remain separate.

Actual context-first private helper derives a fresh sibling per operation from parent, forwards that child synchronously and releases it before next callback. Exact ordered results/errors, parent/value/deadline inheritance and cleanup evidence pass; inline behavior-passing control is rejected for missing authored helper.
