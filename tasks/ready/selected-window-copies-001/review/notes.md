# Author requirement mapping and method boundaries

Two independent applications only: preallocation and subslice retention. Learner owns the count/collection plan and selected-window allocations; the starter supplies only Window and a nil stub. Existing whole-input reserved-copy and one-prefix detached-prefix firsts remain unchanged. Original provenance sources empty.

| Requirement | Public assertion | Required wrong | Useful diagnostic |
|---|---|---|---|
| select nonempty windows in requested order, no zero prefix | TestSelectNonemptyInRequestOrder; TestEmptySelections | reserved-length-then-append; include-empty | selected result count actual/wanted |
| inner values/order | TestSelectNonemptyInRequestOrder | reverse-span | window contents actual/wanted |
| exact outer and inner length/capacity | windowsCheck | extra-outer-capacity; extra-inner-capacity | outer/window capacity actual/wanted |
| preserve input and independent ownership both directions | windowsCheck; TestIndependentRepeatedCalls | shared-input-window | input mutation changed window / output changed input |
| independent duplicate/overlapping results | windowsCheck | shared-clone-windows | window mutation changed peer |

All lengths are guarded before indexing. Public cases include nonzero starts, empty windows, repeated/overlapping intervals, signed/zero values and input spare capacity; grader adds finite combinations and tiny middle windows from256 elements. Nil versus nonnil zero output is deliberately ungraded. No GC timing, unsafe/address probing, performance/allocator threshold or growth-factor assumption.

Actual source review separately checks the count before collection, length/index versus zero-length-capacity/append strategy, and allocation from only the selected interval. A correct alternative uses indexed outer writes and append of only selected elements followed by cap trimming; allocator rounding is allowed. Correct visible cap is never treated as proof of right-sized backing storage.

The whole-input-per-window clone then cap-trim control can pass all value/cap/alias assertions while retaining a large independent backing allocation; actual method rejects it. The grow-then-copy control can pass final exact cap after building a temporary outer result by append; actual method rejects missing planned collection. Both must actually pass public and full profiles before a behavior-pass claim. These passing controls are not required mutants and get no credit. No new mechanism/checker/AST.

The historical author precheck authorized targeted disposable task-check only. Required wrongs now have useful public-only ordinary diagnostics; full canonical, alternative and controls are retained and independently accepted below.

## Independent actual acceptance

The sole reviewer /root/reviewer accepted actual source and current canonical/public/control evidence on2026-10-23; see independent-review.json. Two current canonical checks detect7/7 each; combined26fixtures retain14ordinary public diagnoses and actual accepted/rejected method controls. Author/preflight records above retain their original historical role; Official task-promote transitions are complete; final batch gates remain required.
