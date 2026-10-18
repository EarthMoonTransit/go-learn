# Task Contract v2

Task v2 is selected by `schema_version: 2` in `task.yaml`. Lifecycle and review
metadata remain schema v1. Legacy ready packages and lifecycle task v1 packages
remain valid.

## Three independent dimensions

- `activity`: `implementation`, `debugging`, `refactoring`, `api-design`,
  `test-writing`, `test-repair`, `fuzz-test-writing`, `integration`,
  `performance`, or `code-review-patch`.
- `submission.kind`: `implementation`, `tests`, or `workspace-patch`.
- `verification.profile`: `fixed-tests`, `mutation-tests`, `workspace-tests`,
  `fuzz-seed-tests`, or `benchmark-contract`.

The schema rejects incompatible combinations. An activity does not imply a file
type, and a file type does not select verification implicitly.

## Declarative verification

`verification` names package paths, public/grader/reference package checks,
required mutant ids, optional race/vet gates and deterministic limits. Package
paths must be `./...` or an explicit relative Go package. Timeouts are at most 60
seconds; deterministic runs are 2–3; fuzz duration is bounded by 10 seconds. No
field is interpreted by a shell.

## Executable profiles

### fixed-tests

Submission is implementation code. Directories match the v1 fixed layout:

```text
starter/                    learner starter implementation
tests/public/               visible external-package tests
tests/grader/               hidden external-package tests
solution/reference/         reference implementation
solution/mutants/<id>/      required wrong implementations
```

Starter must fail public and full checks. Reference must pass repeated full
checks. Every declared mutant must compile, preserve the public API and fail due
to a test assertion.

### mutation-tests

Submission is Go test files:

```text
starter/                         insufficient learner tests
solution/reference/              reference learner test submission
verification/reference/          checker-owned correct implementation
verification/mutants/<id>/       checker-owned wrong implementations
```

Both test submissions must compile and pass the correct implementation. Starter
tests must leave at least one required mutant alive. Reference tests must detect
all required mutants. Learner tests use `package subject_test`, call the public
API, avoid uncontrolled dependencies, and cannot contain obvious unconditional
failure or `TestMain`/`os.Exit` bypasses.

### workspace-tests

Submission is a patch across declared paths:

```text
workspace/<path>                  immutable allowlisted skeleton
starter/<editable-path>           starter patch
solution/reference/<editable>     reference patch
solution/mutants/<id>/<editable>  wrong patches
tests/public/<package>/<test>      visible tests
tests/grader/<package>/<test>      hidden tests
```

Every file under `workspace/` must be listed by `workspace_skeleton`. Every patch
must contain exactly `editable_paths`. Tests are added only by the checker. The
same starter/reference/mutant and determinism gates apply across declared package
checks.

## Finite native fuzz profile

`fuzz-seed-tests` uses the closed `verification.fuzz` descriptor from ADR 0032:
`grammar: hex-codec-v1`, `target: FuzzHex`, and two to four declared
`correct_variants`. Its package selectors are exactly `["./subject"]`, active
fuzz duration is zero, and required mutants are nonempty. The layout extends
mutation submissions with `verification/correct/<id>` and the public
`verification/FUZZ_SOURCE.md` contract. Task-init copies the canonical contract;
the prompt links it, workspace export makes it readonly, and learner statements
include it inline. The complete finite Go language is documented in
[FUZZ_SOURCE.md](FUZZ_SOURCE.md).

Public checks repeat the canonical and every declared correct implementation.
Full checks require those assertions to pass before testing every required
defect model with the same learner seeds. Every repetition must produce an
eligible assertion failure to earn detection. Starter tests must retain one
same eligible surviving defect across every repetition. Semantic native events
and complete observed codec values must agree across runs. No hidden corpus is
added. Panic, race, timeout and incomplete observation invalidate execution,
including after an assertion.

## Reserved profiles

`benchmark-contract` defines bounded metadata only. It is not promotion-eligible
until its runner and evidence contract are accepted in a later ADR.

ADR 0018 clarifies the learner-test relevance screen: a task-package selector
reference in a declaration or function counts as API usage, including constructing a
task type followed by receiver-method calls or public-field assertions. Imports
alone do not count. Passing this screen is not completion: reference success
and every required mutant's detection remain mandatory.

## Race-required verification

ADR 0019 applies `-race` to every test invocation for v2 tasks with
`race_required: true`, including public Check and defect models. For the older
fixed, mutation and workspace profiles, a compiled race-test failure is a
detected behavior defect. Finite native fuzz instead requires eligible assertion
failure; a race is invalid execution. Timeouts, compile errors and unsupported
race infrastructure never count as detection. Collections use `task-check/v4`;
regenerate older collections instead of relabeling them. Existing task content
digests and progress formats retain their meanings.
