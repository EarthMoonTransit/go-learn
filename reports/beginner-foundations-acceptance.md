# Beginner foundations batch 1

Date: 2026-09-05. Five original Task Contract v2 lessons.

| Task | Primary skill | Profile | Required defects detected |
| --- | --- | --- | ---: |
| literal-bindings-001 | types.declarations-inference | workspace-tests | 4/4 |
| zero-results-001 | types.zero-values | fixed-tests | 5/5 |
| access-logic-001 | expressions.boolean-logic | fixed-tests | 4/4 |
| division-results-001 | functions.signatures-results | fixed-tests | 4/4 |
| exclusive-loop-001 | control-flow.for | fixed-tests | 4/4 |

Independent read-only agent review approved correctness, observable skill scope,
prompt/test alignment, distinct defect models, provenance and non-duplication.
References pass; starters fail; every required defect compiles and is detected.
All transitions used task-promote. No existing ready package changed.

The first task uses supplied Snapshot plumbing to make variable types observable.
The division task explicitly assesses arithmetic as a secondary skill. The loop
task accepts any correct implementation of the stated interval contract.

Local acceptance passed: fast tests, race, vet, taskintegration, one full
20-package task-check collection, phase1 validation and deterministic report
check using that same collection. Reports were regenerated before repository
tests because the inventory had grown from 15 to 20 packages.

This is a partial content bank, not a foundation coverage release. Runtime tests
and agent review are not human calibration. Estimated times remain provisional.
The first toolchain prerequisite needs the guided preparation in the route
batch; no toolchain encounter or mastery is invented.
