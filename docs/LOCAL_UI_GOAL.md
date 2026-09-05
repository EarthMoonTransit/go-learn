# Bounded delivery: local UI and five strings/bytes lessons

Status: accepted scope; implementation pending. Baseline: master
`9d5e9305012a86417c0df8bfdceaa6e23725658c`, 17 learner-visible ready v2 tasks,
three legacy ready packages, eight drafts and one rejected package.
This inventory is repository evidence, not a fresh read of private learner data.
The historical human checkpoint in ROADMAP.md is not automated calibration.

## Delivery order

| ID | Separate delivery | Acceptance |
| --- | --- | --- |
| K00 | Baseline and access audit | Clean checkout or isolated worktree, exact remote state, toolchain/CI/reviewer access, inventory; existing PR9 untouched. |
| D01 | Documentation | Current UI priority and historical trial labels agree across roadmap, vision, CLI and README. |
| T01 | GoLand workspace compatibility | Reproduce IDE metadata rejection; narrowly exclude root IDE metadata from submission/export while retaining filesystem guards across all three submission kinds. ADR, compatibility and rollback documented; no progress schema change. |
| T02 | Local UI tooling | `learn ui` serves the complete loop below through the existing service/checker; coordinated UI/CLI writes; independent boundary review and green checks. |
| K01 | Independent UI acceptance | Disposable-home integration across submission kinds, browser smoke, concurrency/security/shutdown coverage and explicit IDE availability limits. Acceptance evidence recorded with T02. |
| C01 | Exactly five content tasks | One distinct strings/bytes block appended after the existing 17 tasks; correct provenance, public/full verification, defect models, independent agent review and task-promote lifecycle transitions. |

Every D01/T01/T02/C01 uses its own codex/* branch and PR. Merge only after
`validate` succeeds on the current PR SHA and blocking review findings are
resolved; inspect post-merge `validate` before dependent work. Tooling/content
acceptance follows AGENTS.md and reuses one current full CheckCollection for
validator/report checks. No full-foundation coverage claim at 22 tasks.

## First UI contract

Route → next task → formatted statement → start/resume workspace → open in
GoLand or copy path → public failure/pass → full check and completion → review
→ next. GoLand remains the editor. Short CLI commands remain supported.

Use Go net/http, html/template, embed and small vanilla JavaScript unless a
reproducible limitation requires a separately justified change. Serve foreground
on 127.0.0.1 only; print address and effective root/home. Preserve global
root/home and GO_LEARN_HOME conventions. Do not change global launchers.

UI and CLI share the learner service, checker, home and progress rules. Read-only
requests must not select tasks or change timers. Conflicting operations and
repeated clicks cannot apply a result to a different workspace. Preserve ADR
0009 session semantics and progress v1; hold coordination only during operations,
not for the server lifetime. Explain busy/stale locks without unsafe recovery.

Provide progressive hints, used hints, workspace path/copy, optional GoLand
launch and post-completion taxonomy-based review. Do not present review as an
automatic analysis of learner code. Reference reveal remains explicit through
CLI; no solution endpoint. No fabricated mastery or repetition counters.

Require exact Host, unpredictable session authorization, same-origin protection,
POST mutations, bounded requests/jobs and no-store private responses. Escape
statements/diagnostics; do not expose raw internal errors or checker internals.
No arbitrary filesystem, command, root/home API or shell-based IDE launch.
Shutdown stops new work and makes interrupted verification unambiguously failed;
no promise of persistent jobs across restart.

## Acceptance evidence

All development/acceptance uses explicitly temporary learner homes, never a copy
of the user's real home. Cover implementation, clamp test submission and existing
workspace patches; restart persistence; UI/CLI overlap; duplicate clicks/tabs;
compiler errors and safe full failure; private artifacts; cross-origin, Host,
traversal and escaping; IDE metadata and unavailable GoLand; shutdown during
verification. Authored code executes only through disposable verification.

Use browser automation for a real browser smoke; distinguish it from HTTP tests.
Use a fake launcher/metadata fixtures when desktop GoLand is unavailable and
state that limitation. Independent agent review is not human calibration.

## Five-task content scope

1. Immutable string bytes: implementation.
2. Exact separator/prefix operation: debugging.
3. Typed formatting: debugging.
4. String construction from known values/collections: implementation.
5. Independent byte-boundary verification in another API: scaffolded test-writing.

Preserve three intervening tasks between byte introduction and independent
practice, without claiming mastery. Review the saved party-robot and
welcome-to-tech-palace seeds for suitability; use honest original provenance
for original mechanics. No new adapter/harvest, taxonomy, schema or reserved
profile changes. Keep the original 17 IDs and order. The appended route must
not hit an empty block before the five new lessons.

## Scope and continuation

No theory course, entry diagnostic, embedded editor, accounts, cloud, backend,
progress v2, mastery migration or dynamic repetition in this goal. Those remain
separate later work. No second content batch. Target inventory is 17 → 22 only
when all five tasks are merged and available; drafts/unmerged work do not count.

Keep a gitignored engineering journal outside learner home with baseline, batch
ID, branch/commit/PR, check identity, review findings and exact next step. Do not
publish private learner material, tokens or hidden diagnostics. On a blocked
mandatory gate retain explicit unfinished state rather than weaken acceptance.
The final report records delivered launch commands, master SHAs, inventory,
PR/check evidence, actual browser/IDE scope, untouched learner data, unfinished
IDs and the next engineering step. Learner feedback after five real solutions
calibrates later work and is not required while the user sleeps.
