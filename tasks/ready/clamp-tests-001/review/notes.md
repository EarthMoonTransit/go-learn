# Review notes

## Prior rejection retained

The 2026-09-04 review returned this package to draft because mutation checks
cannot prove table-driven structure, named cases or diagnostic quality. The
lifecycle history is unchanged. Earlier acceptance reports are historical.

## Substantive redesign, 2026-09-05

This task now assesses expressions.comparison through verification practice.
The learner compares expected values for inclusive interval partitions. Test
writing is the submission activity; no testing, table or diagnostic skill is
credited. A discoverable test and assertion pattern is supplied. Table layout,
test names beyond normal discovery and failure-message style are not graded.

Six required behavior models distinguish missing lower/upper limiting, wrong
results exactly at each endpoint, an altered interior result and a zero-valued
singleton result. The exact-boundary models change output; merely replacing
< with <= while returning the same endpoint would be equivalent and is avoided.
The reference includes each partition; the starter passes the correct function
and leaves defects alive. Every correct test organization is accepted.

## Source boundary

The existing exact Exercism counter .meta/config.json record is preserved and
its SHA-256 was checked against the cached pinned artifact. Only that metadata
and root LICENSE were read during this redesign. The metadata supplies a
test-suite direction; it contains no interval API or cases. Attribution now
names its authors, contributors, path, commit and MIT notice. No upstream
exercise statement, canonical cases, implementation, tests or solution entered
the authoring inputs. Independent final review approved on 2026-09-05; all six required defects are detected in disposable verification.

## E01/M1 preparation revision, 2026-10-17

This revision applies the exact starter/hints pair prepared in PR #68
(commit 8acf555f5d6a72f2e2ea280c3f2184d5b036c4c1; ADR 0026). The second
named Test checks a below-range value. Comments and progressive hints explain
Test discovery, the supplied testing.T handle, actual/expected comparison and
compile errors versus assertion failures. No table, slice, callback or custom
reporter becomes a prerequisite. TASK.md, the API, valid-input domain, primary
expressions.comparison skill, practice role, address 03.05 and all six mutants
remain unchanged.

M1 adds testing.basics-diagnostics as a secondary in both task and route. This
credits scaffold-supported practice with discoverable tests and readable
failure feedback. It does not certify independent diagnostic design, require a
particular message, or imply mastery of the entire skill definition. Mutation
verification still assesses behavior detection; the integration regression also
accepts one combined Test using plain Fail calls. Table structure, test count,
case names beyond Go discovery and diagnostic style remain ungraded, preserving
the 2026-09-05 rejection boundary.

In the proposed order, Clamp now introduces the testing exposure needed by
DecisionTable 03.30 and ConstantValueTests 03.26. Both formal blockers disappear.
Only E01/M1 is removed from pending preparation; the remaining 11 actions and
inactive status are preserved. Historical completions and skill encounters are
not retagged. Completion evidence continues to describe the revision actually
completed; no readiness equivalence or scheduler activation is added.

### Public feedback and mutation alignment

For this test-submission profile, public verification runs the learner's tests
against the correct implementation. Full verification additionally runs them
against the six private behavior mutants. Starter public success is intentional;
its incomplete coverage must still fail full verification. The following
assertions are in the starter where noted and in the unchanged reference test
submission. Listing a reference assertion here does not publish a hidden mutant.

| Requirement | Learner assertion | Defective implementation | Expected assertion feedback |
| --- | --- | --- | --- |
| Below min returns min | starter TestClampBelowRange; reference TestClampBelow | missing-lower-bound | Clamp(-3,-1,8)=-3, want -1 |
| Above max returns max | reference TestClampAbove | missing-upper-bound | Clamp(14,0,10)=14, want 10 |
| Lower equality unchanged | reference TestClampLowerEquality | wrong-lower-equality | Clamp(2,2,8)=8, want 2 |
| Upper equality unchanged | reference TestClampUpperEquality | wrong-upper-equality | Clamp(8,2,8)=2, want 8 |
| Interior unchanged | starter TestClampInsideRange; reference TestClampInside | changed-interior | Clamp(5,0,10)=0, want 5 |
| Equal endpoints return that endpoint | reference TestClampSingleton | zero-singleton | Clamp(20,3,3)=0, want 3 |

`TestClampRevisionUsesExistingWorkspaceAndRealVerifier` uses the actual public
path with deliberately wrong expectations in each supplied example: it requires
the named test and exact actual/expected diagnostic. It also accepts the correct
combined plain-Fail alternative publicly and fully, and rejects the incomplete
historical starter privately. Execution is confined to disposable verifier
copies. No hidden mutant is exposed in learner feedback.

### Source and compatibility evidence

Reuse the admitted package at master 64926202e914cd57e8caf99d0f4947e7fa65d437,
path tasks/ready/clamp-tests-001. The registry/source chain, metadata artifact
hash and MIT notice are retained verbatim. This revision consulted no new
external statements, canonical cases, implementation or tests; it does not
claim a new upstream source review. The original creation record remains intact.
Only the authored-content digest changes to
8f1012bd993cac539813d6658a24ef7d12d5b887b52a3165835172eedc3e08c8.

Historical starter/hints are pinned byte-for-byte under
internal/learner/testdata/clamp-before-testing-bridge from the master commit
above. Tests create old workspaces from that explicit baseline before applying
the current revision. They preserve paused/completed records, old/new marker
formats, learner files, paths, permissions, attempts, hints and completion
skills. Every read-only exported file and editable flag remains identical.
The reviewed directional digest pair permits old-workspace continuation without
reset; unknown revisions and reverse migration remain rejected.

Independent agent content review approved this E01/M1 diff on 2026-10-17 with
no actionable findings. It separately recomputed the pinned digest, compared
historical/current fixture bytes and reviewed source-chain preservation, the
secondary's limited claim, public feedback, mutant contract and compatibility.
Targeted task-check detected all six required mutants; the reviewer did not
claim an independent rerun of full acceptance or a fresh upstream artifact audit.
