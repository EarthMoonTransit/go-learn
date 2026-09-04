# Pilot v2 independent red-team review

Phase 3C reviewed the 12 existing Task Contract v2 pilot packages. Five
independent read-only roles covered Go correctness, tests and mutation quality,
pedagogy, provenance and licensing, and learner experience. They did not edit
the repository or execute authored or upstream code. The primary agent resolved
the reports, made all changes, and used `task-promote` for every transition.

No task, source, taxonomy entry, UI, or replacement seed was created.

## Independent verdicts before remediation

| Review role | Ready | Return to draft | Reject |
| --- | ---: | ---: | ---: |
| Go correctness | 4 | 8 | 0 |
| Tests and mutation | 1 | 11 | 0 |
| Pedagogy | 2 | 9 | 1 |
| Provenance and license | 0 | 12 | 0 |
| Learner experience | 5 | 7 | 0 |

The provenance verdict was intentionally strict and preceded remediation. It
found valid pinned source identities, artifact hashes, and compatible licenses,
but insufficient task-specific attribution and durable clean-room evidence.

## Consolidated decisions

| Task | Decision | Principal evidence |
| --- | --- | --- |
| `detached-append-001` | ready | Aliasing is observable in both directions; the duplicate mutant was replaced; provenance and attribution were completed. |
| `map-take-001` | ready | Stored zero, deletion, missing and nil maps, and unrelated entries are observable; comma-ok is now a prerequisite rather than a new assessed skill. |
| `clamp-tests-001` | returned-to-draft | Killing boundary mutants does not prove table-driven structure, named cases, or diagnostic quality. |
| `counter-receiver-001` | returned-to-draft | The private receiver/interface lesson can be bypassed, and two mutants are behaviorally equivalent. |
| `effective-rate-001` | returned-to-draft | The unbounded integer domain is unsound through `float64`; zero capacity is promised but untested. |
| `error-chain-query-001` | returned-to-draft | Wrapping, typed errors, `Is`, and `As` cross several prerequisite layers; the exported error has an unsafe zero state. |
| `formatter-contract-001` | returned-to-draft | The starter already satisfies the interface; two mutants assess string assembly rather than the primary skill. |
| `resource-finalize-001` | returned-to-draft | Cleanup, short writes, wrapping, and multi-error policy are too broad; operation context is not proved and hints leak strategy. |
| `rune-offset-tests-001` | returned-to-draft | Consecutive invalid bytes, truncated UTF-8, valid U+FFFD, and mutation-feedback behavior are missing. |
| `stream-normalize-001` | returned-to-draft | The grader does not prove incremental output, Unicode whitespace, or nil-error short writes. |
| `typed-nil-resource-001` | rejected | A universal reflection predicate is outside the typed-nil interface skill and only partly covered by tests. |
| `utf8-prefix-001` | returned-to-draft | The task assesses explicit UTF-8 decoding rather than `range`; negative limits and valid U+FFFD are not covered. |

## Repairs accepted in this phase

`detached-append-001` replaced `alias-after-restore`, which was semantically the
same as direct append reuse, with `reuse-when-empty`. Review notes now state that
nilness is not observable and relate each mutant to backing-array ownership.

`map-take-001` removed the unobservable requirement to call `delete` only after
successful lookup. `maps.operations-comma-ok` is now a prerequisite, not a new
or secondary assessed skill. The hint states the actual Go rule for deletion.

Both tasks now name their exact Exercism metadata artifact, pinned commit, MIT
license, authors and contributors, authoring exclusions, and later static
clean-room comparison. The comparison found no copied test literals or
substantive lexical overlap. Statements, Problem Specifications cases, Go code,
tests, exemplars, and solutions remained outside authoring.

## Main defect patterns

- requirements existed without a corresponding observable grader condition;
- mutant count hid equivalent mutants or mutants unrelated to the primary skill;
- prerequisite and assessed skills were conflated;
- broad tasks crossed several language mechanisms at once;
- public feedback did not always expose the learner's principal action;
- attribution existed structurally but lacked task-specific human evidence;
- difficulty and time were not calibrated from learner trials;
- original-task policy conflicts with the accepted provenance ADR.

Detailed per-task rationale remains versioned in each package's
`review/notes.md` and `review/review.yaml`.

## Contract v2 defects exposed by real tasks

1. There is no requirement-to-skill-to-test-to-mutant evidence matrix.
2. Mutation execution cannot prove table-driven structure, named subtests, or
   diagnostic quality.
3. Mutant relevance and semantic distinctness are not gates.
4. A skill cannot be represented as both an assumed prerequisite and a transfer
   encounter; the checker rejects prerequisite/secondary overlap.
5. Public and grader kill responsibility is not reported separately.
6. Workspace path sets are not checked pairwise for destination collisions, and
   workspace mutant APIs are not fingerprinted.
7. Learner-authored tests run in a disposable workspace, not a hostile-code
   sandbox; source-less compiled dependencies or OS isolation are still needed.
8. Check lists do not prove that every copied test package is executed.
9. TASK Markdown API and behavioral clauses are not machine-linked to code.
10. Hints have no staged disclosure or leakage evidence model.
11. Provenance hashes, attribution contents, clean-room methods, internal seed
    lineage, release manifests, and notice closure are not fully machine-bound.
12. AGENTS.md's original-source rule conflicts with ADR 0008's project identity
    and empty-source model.
13. Mutation-test submissions have no learner-facing execution and feedback
    contract; `editable_paths: []` is also ambiguous outside authoring tooling.
14. Difficulty and estimated time have no calibration evidence.

These defects are findings only. Contract, schema, validator, and taxonomy were
not changed in this content-review phase.
