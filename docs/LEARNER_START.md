# Start in GoLand

The repository contains the course engine. Each exercise gets its own private
Go module. Read the exercise's TASK.md and edit only the files it names.

Before the first exercise, understand these two toolchain basics:

- A module is a group of packages identified by the module line in go.mod.
  A package groups Go files in a directory; their package declaration names it.
  The course supplies go.mod and the verification harness.
- Go compiles a package before running its tests. Compilation errors point to
  invalid code; a failing assertion means the code compiled but its behavior
  did not match a requirement. Here, learn check assembles and tests a disposable
  copy. learn complete runs the full verification.

This is guided preparation, not an automatically completed task or proof of
toolchain mastery. Later toolchain exercises will assess these skills directly.

## Required manual preparation

Before the first automatic lesson, complete [local modules, workspace and
build/run](labs/modules-build-001/TASK.md). It first assesses module/workspace
edits, then actual target selection, build, execution and tests. Use its authorized
disposable procedure and submit real files/command evidence for checklist review.
The supplied program is preparation; it does not earn code-writing credit.
This manual result is recorded separately and does not change automatic progress.

Continue these manual first encounters at their stated positions; submit actual
files/results for checklist review. They do not create automatic completion:

| Position | Manual unit | Assessed action |
| --- | --- | --- |
| After modules-build, before benchmark work | [Format and vet](labs/format-vet-001/TASK.md) | Apply stable formatting, interpret a vet finding, distinguish vet from correctness |
| After diagnostic-helper-001 and modules-build | [Documentation and target selection](labs/docs-environment-001/TASK.md) | Apply go doc, inspect go env and predict native/cross-target file selection before go list |
| After diagnostic-helper-001 and format-vet | [Native join benchmark](labs/join-benchmark-001/TASK.md) | First author timing boundaries, then analyze allocations, then choose representative repeated comparisons |
| After resource-finalize-001 and filesystem-record-tests-001 | [Reviewed golden file](labs/report-golden-001/TASK.md) | Author file comparison/diagnostics, preserve normal-run expectations and explicitly review a local update |
| After modules-build and formatter-contract-001; before greeter-dispatch-tests-001 and package-boundaries | [Consumer-owned text contract](labs/consumer-text-contract-001/TASK.md) | Author the minimal named consumer interface; verify a foreign provider and interpret excessive-method/concrete-coupling failures |
| After package-inventory-001, modules-build and consumer-text-contract | [Package boundaries](labs/package-boundaries-001/TASK.md) | Diagnose and remove the actual cycle; separately predict and repair internal parent-tree access |
| After named-cases-001 and launch-provided-job-001 | [Parallel subtests](labs/parallel-subtests-001/TASK.md) | Author private fixtures and verify the parent/subtest/Cleanup partial order |
| After parallel-subtests and concurrent-counter-read-001 | [Native race evidence](labs/native-race-001/TASK.md) | Map real conflicting stacks, repair both accesses and distinguish exercised paths from untested branches |
| After hex-seed-classes-001 and hex-wire-invariants-001 | [Retained fuzz regression](labs/fuzz-regression-001/TASK.md) | Reproduce the supplied seed, author a meaningful ordinary regression before repair, retain actual failure/pass and explain the cause |
| After consumer-text-contract, package-boundaries, package-inventory, external-window-tests and writer-double | [Injected configuration project](labs/injected-config-project-001/TASK.md) | Organize module/internal boundaries, author a consumer dependency contract and recording tests, then verify actual build/run/import behavior |
| After constant-kinds-001, numeric conversions and fixed-width-shifts-001 | [Constant contexts](labs/constant-contexts-001/TASK.md) | Choose actual typed/untyped declarations and explain native defaulting, representability and assignment outcomes |

The automatic `bounded-lru-cache-001` at its existing 07.20 address also serves
as the finite-v1 generic-container assessment after maps, callbacks, methods and
generics. Its sequential API and existing tests define the accepted scope.
The native join-benchmark unit also serves as the performance assessment; the
fuzz-regression unit serves as the regression assessment. Those component decisions
do not create additional independent benchmark/fuzz skill encounters.

Later hybrid introductions also require their published method checklist:
`recursive-total-001` needs actual recursive decomposition, and
`interval-table-tests-001` needs learner-selected named data cases, and
`owned-batch-channel-001` needs a private output with one producer/closer.
`drain-values-001` requires an actual channel range with a defined producer close
path; `key-summary-examples-001` requires discoverable, useful executable examples
with intentional output. A behavioral
pass alone does not establish those methods. Their placement in the existing
route follows the required automatic prerequisites.

For `owned-values-001`, also submit the buffering checklist: if n values must
be queued synchronously before any receiver exists, capacity must be at least n.
Explain why the reference chooses len(values), why zero input still closes a
non-nil channel, and why no producer remains blocked after return. This bounded
capacity choice is separate from the existing leak-prevention action. The old
six automatic mutants retain their scope; they do not promise useful automatic
diagnostics for an insufficient synchronous buffer. The supplemental checklist
is manual evidence, with no added automatic coverage or progress.

`two-stage-stream-001` requires two actual communicating range stages, separate
send/close owners and whole-chain termination under full draining. `once-result-001`
requires actual Once initialization and completed result publication, with cached
errors and an explicit non-reset lifecycle. `trace-context-001` requires a private
comparable key, safe typed accessor and request-scoped metadata. Submit the actual
code with these explanations; behavior-only success does not establish method.

For the finite-v1 cancellation first at `cancelable-forward-001`, submit a
checklist identifying the cancellation path at both blocking operations. Explain
which already queued or winning sends may still deliver, how output closes, why
no input receiver remains, and why closing the data input cannot release a stage
blocked sending to an abandoned downstream consumer. This is a single-stage
first; a whole-chain cancellation application follows separately.

For the finite-v1 WaitGroup first at `indexed-job-results-001`, submit your actual
WaitGroup-based implementation. Explain registration before launching each worker,
Done on every worker exit including error results, Wait before accessing the
complete result, and per-call state ownership. `join-all-tests-001` supplies the
mechanism and serves as guided preparation. Other correct synchronization choices
remain valid for the original automatic task; this additional manual method
assessment requires inspecting the actual Add/Done/Wait lifecycle. Existing
13-mutant/race behavior evidence does not automatically establish that primitive.
These supplements add no automatic progress or coverage.

`bounded-job-pool-001` also requires an actual bounded worker pool: identify the
shared job queue, result producers, the single result closer, and the draining
consumer. Explain how worker registration and completion prevent closing results
while a producer can still send, and how each indexed result survives an error.
Result arrival order is unspecified. This applies the earlier collection and
WaitGroup actions to a different bounded queue and channel lifecycle.

For `recover-worker-001`, inspect the actual deferred recovery inside the worker
and the published outcome after both normal return and panic. Explain why recovery
in the caller cannot catch that worker's panic and how one result and closure
remain guaranteed. The finite contract excludes panic(nil), Goexit and cancellation.

For `pipe-exchange-001`, show the actual synchronous Pipe, producer join and both
close paths. Explain which error each callback returns, how the consumer observes
a producer error, and how an early consumer return releases a blocked writer.
For `idle-stream-001`, show the owned Timer, each Reset and every return's Stop;
explain the interval start and input ownership without imposing select priority
when events are simultaneously ready. Ticker use remains a later application.
These actual-code checklists supplement the automatic behavioral checks.

## Your first workspace

Run learn next, then the printed learn start command. In the local UI, click
**Open lesson in GoLand**. It opens the shared `workspaces` folder as one project
and selects the lesson's editable Go file. **Copy project path** lets you open
that same folder manually. For the first exercise the editable file is
subject/bindings.go; subject/api.go is supplied plumbing and stays unchanged.
Next shows the pedagogical chapter and position beside the stable lesson address.
Workspace folders keep their existing section and lesson numbers.

Keep this project open. Each lesson you Start appears in its folder tree, ordered
by the existing section and lesson numbers. Open another lesson's file there or
use **Open lesson in GoLand** to jump to it in the same project. On first use,
GoLand may ask how to open the shared project; existing individual lesson windows
can be closed once you have switched to it. Older unnumbered folders also work.
In the CLI, the common folder is `<learner home>/workspaces` (the home is shown
when `learn ui` starts). Opening a file does not change the active lesson: use Start
or Resume in the UI, or `learn resume TASK-ID`, before checking that lesson.

Use the GoLand terminal with your checkout-based `learn` launcher. A standalone
binary needs `learn -root /path/to/go-learn <command>` when run outside the course
repository, including from the shared project root:

- learn roadmap: see the course blocks and your place.
- learn next: see the next lesson, its preparation and the start command.
- learn check: try your current code.
- learn hint: request one progressive hint.
- learn complete: submit for full verification.
- learn pause: stop the session clock before a break.
- learn resume TASK-ID: continue an existing workspace.
- learn review: read the post-completion review prompts.
- learn feedback "your feedback": save a local note.

Keep TASK.md and go.mod unchanged. New lessons can have different editable files;
the task always names them. Public and full checks use temporary copies, so
hidden material never needs to be opened in GoLand.

An earlier completion remains completed even when it appears later in the route.
Read-only commands do not start a lesson, consume a hint or change your timer.

The repaired firsts at array-copy-stages-001 and struct-copy-comparability-001
require real copy operations and learner-owned prediction assertions. Inspect the
actual array helper/call and every struct field's copying/comparability role; a
formula producing the right snapshots alone does not pass the method checklist.
Their placement precedes dependent lessons while every old address/order remains.

The bounded-record-read-001 IO assessment follows binary/JSON preparation. Inspect
actual private recovery state, independently constructed failures and parsing
boundaries; supplied error signatures and harness calls earn no design credit.
The injected-config-project manual unit is the packages/toolchain assessment and
also requires actual minimal consumer contract, explicit injection and authored
recording tests. Neither manual unit changes automatic completion or progress.

The concurrency applications buffer-window-stream, cancel-error-pipeline,
child-operation-scope and deadline-capped-batch follow their corresponding
first encounters. A behavioral pass still requires the published finite method
checklist: receive-then-send capacity reasoning, cancelable owned stages and
joined termination, explicit context-first helpers with per-call child cleanup,
and one shared budget with every owned context released.

The unchanged bounded-job-pool is the finite concurrent assessment's complete
work alternative: bounded execution admission, exact indexed results/errors,
and joined closure. Cancellation and queue-capacity policy are outside that
assessment. The unchanged owned-batch-channel separately applies worker launch
and termination; it does not establish the close-semantics application, assessed separately below.

Pace a finite stream follows merge-two-streams and the buffer-window application:
review actual local-nil input/release/output gating in addition to behavior.
The Context business-boundary repair follows trace-context; business arguments
stay explicit, while request metadata crosses the callback boundary.

Continue these manual applications with actual files/results and checklist review:

| Placement | Unit | Independent action |
| --- | --- | --- |
| After owned-values, closed-receive and select preparation | [Channel permissions and closure](labs/channel-permissions-close-001/TASK.md) | Choose helper directions, predict comma-ok buffered closure and observe real compiler/post-close behavior |
| After native join benchmark | [Mutable-input benchmark](labs/reset-sort-benchmark-001/TASK.md) | Own reset/timing/allocation boundaries and representative repeated comparisons with equivalent results |

The unchanged injected-config-project also applies its actual production import
graph and internal delegation with a separate client. Supplied names/layout,
module commands, cycle diagnosis and test-package choice earn no extra credit.
The unchanged cancel-error-pipeline separately applies whole-pipeline termination
on intermediate failure or parent cancellation; this is no second cancellation
encounter or new automatic task. Manual review never changes automatic progress.

The branching-path recursion application follows pointer/struct preparation and
requires actual recursive branch work, complete paths and independent storage.
The streamed job-outcome application follows pool/recovery preparation; review
per-job recovery and continued bounded-worker work alongside indexed output.
The limited Pipe relay follows PipeExchange; it owns the bounded consumer and
producer join across limit/EOF/destination-error endings. The keyed Once registry
follows once-result and mutex preparation; review stable per-key entry identity,
compound lookup/insert locking and actual Once execution separately from behavior.

| Manual placement | Unit | Independent action |
| --- | --- | --- |
| After generic type-set first encounter | [Named-slice selection](labs/named-slice-selection-001/TASK.md) | Design exact element/slice constraints and preserve the named result's static methods |
| After package-boundaries, external-window and injected-config-project | [Shared record boundaries](labs/shared-record-boundaries-001/TASK.md) | Relocate shared data/composition to break a real cycle and independently choose private/external test packages |

Retain actual native diagnostics, source snapshots, repairs and finite method
review. Manual evidence and behavior-passing method controls do not add automatic
coverage or learner completion. Supplied algorithms earn no authorship credit.

The classified-stage application follows the three error-chain first encounters.
It owns transparent/opaque error policy and classifies the returned boundary with
actual errors.Is/errors.As. The word-stream application follows Scanner and
buffer preparation; it owns word tokenization, bounds and a private buffered
writer with explicit error/flush ownership. The JSON record patch follows JSON
shape and package/receiver preparation; it owns exported fields, tags, private
state and partial-update behavior across a real package boundary. Periodic-budget
work follows duration arithmetic and idle-stream preparation, owning one Ticker,
one independent budget Timer and actual elapsed/remaining duration calculations.

| Manual placement | Unit | Independent action |
| --- | --- | --- |
| After modules-build | [Workspace/artifact consistency](labs/workspace-artifact-consistency-001/TASK.md) | Diagnose workspace-masked module replacement and prove source/run/binary freshness across explicit rebuild |
| After hex-seed-classes, hex-wire-invariants and fuzz-regression | [Run-compaction fuzz](labs/run-compaction-fuzz-001/TASK.md) | Author a lossy-transform fuzz target/oracle, reduce an actual failure and preserve native plus ordinary regressions before repair |

Retain finite actual source/method evidence alongside behavioral tests. Native
manual work never grants automatic coverage or learner completion; the bounded
native1x observations do not establish mutation search or automated minimization.

The atomic-count testing application follows interval-table-tests, named cases and diagnostic helpers.
It owns stateful partitions, isolated mutable fixtures, serial named subtests and
child-aware helper assertions. The dual-frame application follows binary-record
and text-level preparation. It owns both marshaling method pairs, exact wire
validation, atomic receiver updates and independent storage.

| Manual placement | Unit | Independent action |
| --- | --- | --- |
| After initialization-values | [Initialization snapshot](labs/package-init-snapshot-001/TASK.md) | Diagnose file-order dependence and replace late registration with an explicit variable dependency |
| After docs-environment and format-vet | [Selected tool diagnostics](labs/selected-tool-diagnostics-001/TASK.md) | Interpret actual source selection, formatting, vet and documentation before repairing API semantics |
| After parallel-subtests, cleanup-order and worker-completion | [Isolated stamp tests](labs/isolated-stamp-tests-001/TASK.md) | Introduce a private clock seam and own parallel fixtures whose cleanup follows descendant completion |
| After key-summary-examples and report-golden | [Log examples and golden](labs/log-examples-golden-001/TASK.md) | Author discoverable stateful examples and an independent complete golden baseline, then review an intentional update |

Native evidence and finite method review remain separate from automatic coverage
and learner progress. A clean tool command or passing weak oracle does not replace
the substantive source and observed-output acceptance recorded with each unit.
