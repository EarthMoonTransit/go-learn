# Branch paths: author evidence, independent decision pending

The first recursive-total task divides a flat slice and combines sums. This
application recursively explores every nonnil child and constructs independent
root-to-leaf paths with branch history. Only functions.recursion application.

Public requirements map: nil/leaf/nil children and branch multiplicity use
TestPathShapeAndOrder count assertions (nil-as-leaf, skip-first/last-child,
duplicate-leaf, skip-negative, ignore-leaf); exact full DFS path labels/order use
full ordered path assertions (reverse-children, omit-parent-label,
carry-sibling-prefix, shared-path-backing). Input preservation is checked after
every call (mutate-label). TestPathsOwnTheirStorage modifies returned path prefixes
and checks siblings, separate calls and original nodes; shared-path-backing also
fails here. Eleven public-only wrong-reference copies each produced an ordinary
named assertion, without compiler/panic/race/timeout evidence. Canonical reference,
all11mutants,determinism,vet pass. Public-only reference and public/full recursive
alternate and explicit-stack control all pass:16fresh removed fixture repositories.

Finite method inspection: actual reference walk recursively visits each nonnil
child; nil return and no-live-child leaf emission terminate; subtree height
strictly decreases. Prefix accumulation directly determines emitted copied paths.
Bottom-up alternate recursively gathers every child's subpaths and prepends label.
Explicit-stack control passes behavior but has no authored recursion and must be
rejected by source/method review. No automatic recursion/AST claim. Tree limits
exclude cycles/DAGs; tests make no timing or performance assertion. Active existing
prerequisites only; route placement must follow pointer/struct preparation.

Original task-init provenance, empty external sources, actual digest retained;
release commit binding deferred. Actual independent content/method review and
all batch gates still required before ready acceptance.

Sole actual review approved2026-10-23 at05:44:48UTC with no findings. Official lifecycle transitions and final batch delivery remain separate.
