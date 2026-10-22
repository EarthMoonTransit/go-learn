# Finite-v1 introduction batch I2

Base: `50412ef271c439cb8c7508341578a6ad03342353` (PR105 merge).
Branch: `codex/v1-introductions-02`. Sole independent reviewer: `/root/reviewer`.
Content/method/native reviews accepted 2026-10-22; final batch gates/merge pending.

| Unit | Actual evidence | Finite-v1 role |
| --- | --- | --- |
| drain-values-001,09.18 | Fixed-tests 9/9, race, public-only 9/9, manual-receive valid behavior control; actual range/close checklist | channels.range-drain first |
| key-summary-examples-001,08.08 | Native executable examples pass correct implementation and detect 8/8 defects; naming/Output/documentation review | testing.examples first |
| format-vet-001,manual | Three independent copies: formatting/idempotence, targeted vet diagnostic/repair, clean-vet wrong-result counterexample | toolchain.formatting-vetting first |
| join-benchmark-001,manual | Correctness equivalence,24 reference measurements; actual timing loop, allocation analysis and repeated representative comparison; three wrong specimens | Three distinct firsts: loops-timing,allocations,experimental-design |
| report-golden-001,manual | Seven copies: reference,four assertion defects,silent updater,explicit update; actual header-only diff reviewed before final ordinary pass | testing.golden first |
| bounded-lru-cache-001,unchanged07.20 | Existing complete generic API, reference/20 mutants and independent semantic review | component-generic-container |

Six units: two new automatic packages, three manual units and one unchanged
component reuse. Seven original first slots reduce22→15;77 applications remain.
One component reduces6→5. There are196 automatic route lessons,199 ready packages
(including three legacy),206 total checked packages and four cumulative manual
units. Manual outcomes do not enter automatic coverage or progress. Materials
readiness does not establish learner mastery.

## Public/method evidence and compatibility

Drain's initial independence test did not establish overlapping calls. Corrected
TestCallsKeepOwnResults starts A, holds it after a value, starts B, then completes
A while B stays open and finally completes B. A mutex-protected shared-pending
mutant fails public/full without a race or timeout. Removing this public assertion
in a disposable fixture leaves exactly that mutant alive 8/9 while reference/race
pass. The retained mutant and ordinary task-check protect the assertion without
a tooling extension. A manual receive loop passes behavior 9/9 but fails the
separate prescribed-range method. Examples use actual discoverable ExampleSummary
and ExampleSummary_empty; the supplied trivial empty example earns no new credit.

All transitions used task-promote. All 4043 previous task files,194 old route
objects/addresses/roles/skill tuples and old relative order are unchanged. New
addresses are appended; pedagogical insertion follows prerequisites. Manual
positions and LRU assessment are linked in the existing learner start index.
No schema, engine, checker profile, learner state or production tooling changed.

## Native evidence

The owner-approved finite procedure was applied to exact pre-reviewed manifests:
14 independent canonical disposable roots,52 commands,private offline local Go1.27,
60-second bounds,actual files/results and cleanup. Independent review checked
source snapshots, digests, normalization and actual method. See each lab's
review/acceptance.md and review/execution.json. All roots are removed.

The initial format recipe expected gofmt -d status0 on a diff. Actual Go1.27
returns1 for formatting differences (empty stderr),2 for errors; the receipt and
installed gofmt source confirmed this. Only that expected status was corrected,
independently reviewed and reexecuted. No infrastructure extension was needed.

Benchmark correctness precedes comparison. Twenty-four reference observations
show local variation; Builder is not a universal winner and is slower/more bytes
for the longASCII fixture here. Setup-inside and single-tiny execute but fail
method/design; missing separator fails output equivalence before timing. The
actual analysis separates B/op,allocation events,result lifetime and elapsed time.

Golden ordinary runs preserve expected bytes. The silent updater actually changes
them and is rejected despite process success. Explicit update changes only
RECORDS v1→v2; actual diff inspection precedes a passing ordinary run with no
further write. Receipt intervals distinguish initial ordinary v1→v1, explicit
v1→v2 and final ordinary v2→v2. No unexecuted instruction is counted as ready.

## Existing pairs and component reuse

Semantic review of all 49 baseline candidates is complete:44 accepted,5 concrete
repairs. Backlog rows pin exact ab880dea source paths, learner actions and
assessment. The five are arrays.value-semantics,structs.copy-equality,errors.typed,
interfaces.small-contracts and types.constants. Provided copying/declarations or
interfaces do not establish an independently authored outcome.

Repair allocation uses at most three extra units: combine separate array/struct
snapshot actions in one first encounter; add one small consumer-interface first;
at most one constants application. Typed-error and small-interface applications
fit the allocated IO/packages components with actual contracts; constants may fit
the already allocated native build application. None earns credit in advance.
Arithmetic after I2:196 automatic+4manual+15first+77applications+5components+3repair
=300≤302. Original first-slot deadline and extra semantic repairs remain separate.

The generic-container component reuses bounded-lru-cache-001 at unchanged source
ab880dea. Empty private-state scaffold requires the learner's generic K/V API,
typed misses/present zero/nil values, independent zero-state instances, recency,
eviction, updates, ordered stoppable Range and reusable Clear. Public generic,
composite/named/pointer keys and slice/struct values, independent grader oracle and
all 20 mutants were inspected. I2 stable collection passes the unchanged component reference and all20mutants.
Scope is sequential/non-reentrant, with no O(1),concurrency,GC-retention or external
resource-cleanup claim. Existing SHAD adaptation provenance at
6718601587ec77cfaa632bbd29a8ca0716fe81ec,four admitted lrucache artifacts and MIT
notice remain intact. No new borrowing or extra container package.

## Acceptance and accounting

I2 starts13:14UTC (16:14MSK),2026-10-22. One stable206-package CheckCollection passed206/206 in18m01s
(13:36:49–13:54:51UTC; childCPU34m09s). Quality and both numeric coverage reports
reuse it. Fast passed2m06s, race6m46s and vet0.24s. Integration and the final
validator/report-check consumers are running; final HEAD review and exact-head
validate/merge remain. Local checks/content/review
count toward the four-hour checkpoint; overlapping time is counted once.

Recipe diagnosis is conservatively≤3minutes in I2; rolling infrastructure≤13m03
including I1. No production tooling fixes. No sleep/user-wait/CI-only idle exclusions
claimed so far. After this second merge, update observed ordinary/manual/component
work estimates from these timestamps; do not project learner mastery or full-bank
minima from finite-v1 readiness. Next batch retains introduction priority.
