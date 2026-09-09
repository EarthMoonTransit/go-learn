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

## Reserved profiles

`fuzz-seed-tests` and `benchmark-contract` define bounded metadata only. They are
not promotion-eligible until their runners and non-flaky evidence contracts are
accepted in a later ADR.

ADR 0018 clarifies the learner-test relevance screen: a task-package selector
reference in a declaration or function counts as API usage, including constructing a
task type followed by receiver-method calls or public-field assertions. Imports
alone do not count. Passing this screen is not completion: reference success
and every required mutant's detection remain mandatory.

## Race-required verification

ADR 0019 applies `-race` to every test invocation for v2 tasks with `race_required: true`, including public Check and defect models. A compiled race-test failure is a detected behavior defect. Timeouts, compile errors and unsupported race infrastructure never count as detection. Collections now use `task-check/v3`; regenerate older collections instead of relabeling them. Task and progress schemas remain unchanged.
