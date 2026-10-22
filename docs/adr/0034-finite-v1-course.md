# ADR 0034: Finite v1 course and explicit reviewed assessment

- Status: accepted owner decision; implementation through reviewed PRs
- Date: 2026-10-20
- Supersedes: open-ended delivery priority; ADR 0031's recursion machine-proof
  prerequisite for v1 hybrid authoring; implementation of ADR 0033 during v1
- Preserves: taxonomy meanings, numeric minima, schemas, lifecycle, source
  policy, quality gates, learner state and full-bank release requirements

## Owner decision and scope

The owner approved automated tests plus checklist review in chat for outcomes
where method or tool use matters. The 2026-10-22 owner authorization replaces the earlier pause-after-five limit:
continue autonomous reviewed batches of normally 5–8 tasks (maximum 10) through
finite-v1 readiness, with one independent reviewer and exact-head green validate
before each merge. Follow the updated delivery plan's introduction priority,
native evidence deadlines, budgets and stop conditions. P21 remains deferred.

V1 is a learning-route readiness milestone for the frozen 141 foundation
skills, distinct from the old 350–500-task full-bank release. Each skill needs
two materially different reviewed encounters, including independent application.
Six integrated component assessments complete the route. Combined credit needs
actual assessed outcomes: supplied mechanisms, mentions and prerequisites do
not count. Materials readiness does not establish learner mastery.

At `21d89d59939c3ee1f7ea891c136798decce9e9da`, the allocation is 171 existing
lessons, 33 missing first encounters, 92 independent-application slots and six
components: at most 302 learning units. The 49 apparently qualifying existing
skill pairs are metadata candidates requiring semantic review. Existing suitable
components or genuinely combined outcomes can reduce new work.
The [backlog](../plans/v1-backlog.json) is planning data, not task metadata,
source admission, coverage evidence or a new validator.

The four specialist sync outcomes and deeper bank requirements remain visible
deferred work. V1 deliberately reduces initial depth; two encounters do not
satisfy the old minima, three distinct task types or required modes.
No skill meaning, minimum or historical report changes.
`validate -coverage` and `-foundation-coverage` retain their current meanings.
This supersedes ADR 0031's equation of final goal acceptance with full coverage
only for the distinct v1 readiness claim. Full-bank release/export gates remain
binding; v1 readiness cannot turn numeric failure into a pass.

## Assessment without new engines

Keep fixed-tests, mutation-tests, workspace-tests and the accepted finite hex
fuzz profile. No new profiles, source grammars, rewriting, trace protocols or
general benchmark/module runners enter v1. Infrastructure fixes must address a
concrete blocker in named current tasks.

Behavior, errors and boundary values retain executable checks. Where method
also matters, publish a finite checklist before authoring. The learner supplies
code and actual tool results; review records exact evidence, passed criteria and
remaining gaps. A behavior pass alone does not satisfy an outstanding method
criterion. Automatic completion semantics are unchanged; no silent method credit.

- Recursion: test behavior and inspect real decomposition, termination and
  combination. Do not build machine proof rejecting every iterative equivalent.
- Reproduction: require actual deterministic reproduction, repair, a meaningful
  retained regression and causal explanation. Existing independent repair/test
  tasks can support this; no bespoke Cause getter or trace engine is required.
  A supplied failure does not earn discovery credit.
- Benchmarks: inspect native benchmark code and repeated results, setup/work
  boundaries, representative inputs, allocations and semantic equivalence.
  Do not use machine-specific timing thresholds as the grade.
- Modules/tools: inspect files and native command evidence from a small isolated
  project. Documentation does not add editable go.mod/go.work to current profiles.

Manual units remain outside `tasks/ready/` and automatic route/coverage counts.
They require a precise contract, verified reproducible reference procedure,
checks where applicable, realistic wrong examples, useful feedback, provenance
and dated evidence review. Missing execution/review means unfinished work.
Existing source-candidate `manual-artifact-review` keeps its original meaning.
No new UI, scheduler or completion store is necessary for a checklist/chat review.
Source-checkout task execution remains forbidden. On 2026-10-22 the owner
explicitly authorized the finite manual procedure in the delivery plan and its
narrow AGENTS rule 9 exception: reviewed build/run, modules/workspaces and
benchmark commands may execute outside task-check only for manual units outside
tasks/ready, in fresh disposable copies with clean offline environments, a
60-second command limit, retained real results and independent review. This
adds no runner, automatic coverage or acceptance without executed specimens.
Automatic task profiles and all of their existing gates are unchanged.

## Compatibility and migration

No schema, taxonomy, task-data or learner-data migration is activated.
Preserve all 171 task IDs, addresses, role/skill tuples, stored solutions and
completion interpretation. New tasks use official lifecycle transitions,
append-only addresses and reviewed pedagogical insertion after prerequisites.
Manual evidence never retroactively modifies automatic credit.

Retain ADR 0033/private P21 evidence as deferred design. ADR 0032's working hex
profile is unchanged. ADR 0031's migration, meanings and full-bank requirements
remain binding; only its special recursion machine-proof precondition is
superseded for v1 hybrid authoring.

## Acceptance

The [delivery plan](../plans/finite-v1.md) defines the bounded run and stop.
This documentation change may reuse a fresh unchanged-input collection.
Independent review and green exact-head GitHub validate are required.
Subsequent automatic tasks retain all public/grader/reference/mutant, useful
feedback, source, isolation, determinism and applicable race gates. Checklists
cannot waive executable failures or `ready` requirements.
Preparation content retains ADR 0022 acceptance. Final v1 readiness requires
the complete outcome matrix, six components, prerequisite/compatibility review,
all required automatic/manual evidence and a truthful remaining-bank report.
