# Sequential personal route acceptance

Date: 2026-09-05. Tooling batch; no task or progress contract changes.

The versioned route defines the twelve opening lessons and visible placeholders
for later blocks. `learn roadmap` shows readiness and completion separately;
`learn next` selects the first unfinished lesson, preserves earlier completions,
uses `resume` for an existing workspace and stops at missing content.
The English preparation guide explains the GoLand workspace and toolchain
prerequisites without recording an invented encounter or mastery.

Independent read-only review identified two issues before acceptance: resume
guidance for an existing workspace and acceptance of deprecated skills. Both
were corrected and covered by regression tests.

An actual opening-lesson trial exposed an existing workspace submission bug:
the directory whitelist overwrote a declared `subject/` directory with false.
The narrow fix preserves declared directories while continuing to reject
undeclared directories and extra files. Independent review approved the fix;
reverting it makes the regression test fail.

Final local acceptance passed: fast `go test ./...`, race, vet, tagged integration,
one full 20-package `task-check` collection, phase1 validation and report check.
Validator and report check reused the same collection. The five lessons in the
second content batch are separate; this batch sees seven of twelve route tasks
as ready and correctly displays the remaining five as planned.

A tagged integration test and a CLI smoke trial used temporary learner homes to
exercise the actual first lesson: start, resume recommendation, failing starter,
successful completion and selection of the next lesson. Authored code ran only
through the disposable learner runner. The canonical task content was unchanged.
A read-only roadmap against the real learner home retained its progress file
byte for byte and reported the existing completion.

These are automated acceptance checks, not human calibration. Skill mastery,
spaced repetition and their progress migration remain future roadmap work.
