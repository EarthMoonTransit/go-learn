# ADR 0026: Preserve Clamp workspaces across the proposed testing bridge

- Status: accepted for preparation tooling; E01/M1 content remains unreleased
- Date: 2026-10-17
- Base: master a2dfbd3dff1132d9ef46ca39b1b0aa756ba49c76

## Problem and bounded decision

The inactive pedagogical plan still has two formal testing prerequisites:
DecisionTable (03.30) and ConstantValueTests (03.26). E01/M1 proposes teaching the
basic Test API explicitly in existing Clamp (03.05), followed by a coordinated
`testing.basics-diagnostics` secondary in its manifest and route.

Editing the scaffold or hints changes TaskContentDigest. The learner
currently rejects an old workspace after that change and suggests reset, even
though the Clamp API, editable path and accepted submissions remain compatible.
Prepare that compatibility before the separate content batch, following the
bounded-pair approach of ADR 0014. Accept only this directional pair for
`clamp-tests-001`:

- Stored: `ed280dce8f21e6154adeb90a368972e45589af8788ddd6eedf030ce4a65b630b`
- Current: `8f1012bd993cac539813d6658a24ef7d12d5b887b52a3165835172eedc3e08c8`

The target is pinned by two proposed files in
`internal/learner/testdata/clamp-testing-bridge/`. They add a second named Test
example and explain actual/expected feedback. They are acceptance fixtures,
not released task content. Applying them over the exact base package must
produce the target digest. Any further authored-content edit requires a newly
reviewed pair; this is not a generic scaffold-migration policy.

The fixture reuses the admitted Clamp package and its existing provenance at
`tasks/ready/clamp-tests-001/provenance.yaml` in the base commit above. Its
Exercism metadata concept seed remains pinned to registry `exercism-go`, commit
`ae1477abfeeb9fbeaf971caee220e91f9e4df790`, path
`exercises/practice/counter/.meta/config.json`. No external statement, canonical
cases, code or solution was consulted or copied for this revision. This is reuse
of existing admission, not a fresh upstream artifact review.

## Preservation boundary

The exception permits validation of the old workspace; it never rewrites a
marker, prompt, learner tests, path, timer, hints, attempts or completion record.
The prompt, API declarations and every read-only workspace file stay byte-identical.
Existing workspaces retain their own tests; hint reads use the revised guidance
after its separate release. New workspaces
will receive the revised scaffold. Learners can keep extending their existing
tests without reset. Both revisions use the same Clamp API, valid input domain,
editable path and six behavior mutants. Test count, organization and diagnostic
style remain ungraded.

Schema, task/submission identity, generation, fixed-file and symlink checks stay
in force. Unknown hashes, other tasks and reverse migration are not authorized.
Rollback keeps old-content workspaces usable; a workspace created with the new
revision still requires that content and a compatible learner. No automatic
reset or downgrade is provided.

The fixture simulates the future M1 secondary to prove that reading an old
completed workspace does not add it to CompletionEvidence or SkillEncounters.
It grants no readiness equivalence or mastery and does not change scheduling.

## Acceptance and remaining work

Fast tests cover paused/completed workspaces, legacy/generation markers,
byte/path/permission preservation, historical evidence and exact directional
matching. Integration uses disposable copies to resume an old workspace, pass
public checks, reject its incomplete tests in full checks, discover both new
examples with actionable failures, and accept one combined test using plain
Fail calls. The real mutation contract remains authoritative.

Canonical tasks, route, taxonomy and the 151-row mapping are unchanged. The plan
stays inactive with both formal blockers and all 12 pending preparation records:
E01/M1 Clamp; E02 Prepend; four M1 prerequisite reviews (DetachedAppend,
PointerChain, DiagnosticHelper, NamedCases); P05 Formatter before Shape;
P06/P07 before Capture; P11/P12/P13 before Worker. Conditional ValueAverage
metadata remains a review question, not a required addition.

Next: review and release the pinned E01 scaffold together with its M1 manifest
and route secondary, refresh provenance/reports, and remove only E01/M1 from the
pending plan. Preserve historical completion evidence. Then address E02 and the
remaining bounded metadata/content batches. T0b activation remains separate.

## E01/M1 application, 2026-10-17

The following content batch applies the exact pinned starter/hints to canonical
Clamp and adds the coordinated manifest/route secondary. Its digest matches the
compatibility target above. The two formal testing prerequisites are resolved;
11 other preparation records and inactive status remain. No compatibility code,
API, prompt, address, role, scheduler or historical evidence rule changes.

The old starter/hints are now explicitly retained in
`internal/learner/testdata/clamp-before-testing-bridge/`, copied from master
`64926202e914cd57e8caf99d0f4947e7fa65d437`, path
`tasks/ready/clamp-tests-001`. Compatibility tests reconstruct that old package
before creating old workspaces, then apply the current content. This preserves
the historical test boundary after the canonical package advances. The earlier
status and remaining-work description above record the original tooling batch;
this application is a preparation content batch under ADR 0022, not foundation
release/export or activation of the new order.
