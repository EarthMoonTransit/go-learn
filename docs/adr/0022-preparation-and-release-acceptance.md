# ADR 0022: Preparation batches and release acceptance

- Status: accepted
- Date: 2026-09-10
- Basis: foundation-next §9 requires resolving existing phase-policy ambiguity

AGENTS rule 8 requires coverage for a release content batch; QUALITY_GATES and
ROADMAP use phase1 for incomplete-bank preparation. Preserve both requirements
with an explicit classification before executing the first wave.

P0 documentation, P1 tooling, P2 theory, P3 source review, P4T tooling,
P4C correction and P5 bounded content prepare the incomplete foundation bank.
A ready package means that package passed its technical, public-feedback and
review gates; it is not a declaration of complete foundation coverage or a
public release/export. These preparation batches use `validate -phase1`.

Any batch designated public release/export, and the final foundation acceptance,
must additionally pass `validate -coverage`. A failing coverage result cannot be
relabelled preparation after the fact to ship that release. This wave does not
publish/export the bank or claim the 142-skill minima are satisfied.

For tooling and content: fast tests → race → vet → integration, then one full
`task-check -all -output` for unchanged inputs. Refresh generated reports from
that collection when input digests change; validator and `task-report -check`
reuse it with `-checks`. Documentation-only changes may reuse successful exact
input evidence; their PR still requires independent review and exact-head CI.
Every package gate, ten-task bound, disposable execution and protected-master
process remains required. No schemas, minima, task data or learner state change.

Rollback restores prior documentation; it cannot authorize a release without
coverage. Historical milestone reports retain their original phase and claims.
