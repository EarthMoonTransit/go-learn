# Review notes

## Decision

Returned to draft. The learner is asked to demonstrate
`testing.table-driven`, but Contract v2 observes only whether submitted tests
pass the reference and kill mutants. Separate ad-hoc tests can satisfy that
behavior, so named cases, a loop, exact partitions, and actionable diagnostics
are not proved.

## Required repair

Either add a trusted structural evidence gate or stop claiming table-driven
testing as the primary skill. Boundary mutants should distinguish exact and
exclusive limits. The next review must map each requirement to executable
evidence and complete exact Exercism metadata attribution plus durable
clean-room comparison evidence.
