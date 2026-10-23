# Author requirement mapping and actual method boundary

Three independent application candidates only: dynamic interface type/value, Stringer and typed-nil. Learner owns both return conversions, conditional nil interface and two formatting methods. Empty starter bodies supply no behavior. This differs from supplied Greeter dispatch tests, static nested String representations and typed-nil input rejection. Original provenance, no source donors; all existing firsts preserved.

| Requirement | Public assertion | Required wrong | Diagnostic |
|---|---|---|---|
| exact Label dynamic value snapshot | TestSnapshotAndLiveBindings; TestAbsentCounterIsNilInterface | label-pointer | snapshot dynamic type actual/wanted |
| original live Counter pointer and external changes | TestSnapshotAndLiveBindings; TestSeparateBindingsStayIndependent | counter-copy; omit-live | live pointer / nonnil live binding |
| absent counter returns nil interface | TestAbsentCounterIsNilInterface | typed-nil-return | nonnil interface dynamic type; no method call |
| exact stable signed/text representation through real fmt dispatch | TestDirectStableStrings and binding helper | trim-label; absolute-count | quoted formatted text actual/wanted |
| preserve receiver/input and repeated results | TestDirectStableStrings and repeated binding formats | mutate-counter | formatted text / changed receiver |

Every wrong shape/nil is guarded before dependent operations. The typed-nil result is compared with nil and diagnosed with %T only; no nil Counter method is called. Grader extends finite valid text/count pairs and rebinding the same original pointer. Tests never use Parallel; no concurrency obligation or race credit.

Valid alternative uses strconv and a local initially nil interface. Both interface identity and after-assignment mutation outcomes must pass, not merely equal initial text. Method review inspects direct nonrecursive field formatting and actual interface assignments. The finite guarded-reentry Stringer control uses a private two-level guard/helper and correct observed output; it is method wrong despite behavior passes. Both public/full and existing vet must actually pass before any behavior-pass claim. Unguarded recursion is neither executed nor credited; vet is never weakened.

The historical author precheck authorized targeted task-check only. Required ordinary public diagnostics, the correct alternative and method controls were then retained in copied fixtures and accepted by the sole reviewer below. No new checker/AST or supplied-mechanism credit is claimed.

## Independent actual acceptance

The sole reviewer /root/reviewer accepted actual source and current canonical/public/control evidence on2026-10-23; see independent-review.json. Two current canonical checks detect7/7 each; combined26fixtures retain14ordinary public diagnoses and actual accepted/rejected method controls. Author/preflight records above retain their original historical role; Official task-promote transitions are complete; final batch gates remain required.
