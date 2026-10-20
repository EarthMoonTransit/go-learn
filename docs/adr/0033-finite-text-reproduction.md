# ADR 0033: Finite text-boundary repair and reproduction

Update 2026-10-20: implementation is deferred by [ADR 0034](0034-finite-v1-course.md). Retain this design and evidence; do not continue its profile, prototype or collection-v5 work during v1.

- Status: independently accepted design; preparation starts after docs PR merge,
  no execution or activation approval
- Date: 2026-10-19
- Extends: ADR 0007, 0008, 0019, 0022 and 0032

## Decision and delivery boundary

Introduce a separate, closed `reproduction-tests` profile for an editable helper,
retained executable regression and typed causal report. The first grammar is
`tagged-text-v1`. This document and the normative
[mechanism design](../plans/reproduction-text-mechanism.md) are one decision.
They do not activate metadata, run a parser, add tasks or award coverage.
`hex-codec-v1` retains its lossless-codec meaning and excludes this assessment.
Independent exact-input re-review on 2026-10-19 closed five initial findings:
semantic versus invalid outcomes, absent-length checkpoints, report invocation,
honest v5 nonexecution/aggregation and aggregate allocation charges. The external
review receipt `p21-finite-design-rereview-epoch-7-v2.json` has SHA-256
`eb10abd249d62b88105341f87ba99f2f058989b45d4970af6d707be166557282`.
It approves design/preparation only; concrete prototype, launcher, execution,
production activation, content and merge gates remain separate.

The prospective introduction repairs text consumption in a count-prefixed,
tagged-record parser. The learner must execute a supplied smallest meaningful
case against the frozen faulty subject and their repair, retain that test, and
report independently checked causal facts. A supplied case earns no discovery,
randomized-search or minimization credit. This is one encounter and one task
type; it does not satisfy the whole reproduction skill or foundation depth.

The exact donor is registry `ardanlabs-gotraining`, commit
`4ae5a46bfc665909a7cb29e11c18775330c63401`, path
`topics/go/testing/fuzzing/exercises/exercise1/exercise1.go`, SHA-256
`16cbf64df2525b105e9c885090bcfec02bfaffef7452a06a81bc9266a5ad8b5e`.
The [source admission](../../catalog/intake/2026-10-18/p21-parser-source-review.json)
permits adaptation, subject to its retained license/notices and attribution.
Retain the count, tags, shared byte-string helper, decimal fields and selected
unchecked length expression. Rename records and add explicit local guard/error
policies; do not claim unmodified-upstream equivalence. Do not reuse the donor
test vectors, answer, legacy fuzzer or seed as a discovered failure. No
reference-only source supplies code, tests or educational text.

## Public parser and causal contract

The readonly package exposes `Decode([]byte) ([]Record, error)` and
`Record { Kind string; Label string; Value int }`. Its local policy is:

| Input | Result |
|---|---|
| Empty | Malformed. |
| Header byte C | Allocate C zero records; a header alone succeeds, including C=0. |
| Tag 0 | Advance current record, initially zero; reaching C is malformed. |
| Tag 1 / 2 | Set Kind / Label from a length byte N and exactly N payload bytes. Zero length and arbitrary byte strings are valid. Repeated fields overwrite. |
| Tag 3 | Consume exactly two ASCII decimal digits, set Value in 0..99. |
| Unknown tag, out-of-range current record, missing/truncated text or numeric input, nondigit | Malformed. |

Successful input may leave records unpopulated. Fields consume only their own
payload. On any malformed input, return no records and a nonnil error, never
panic. Nil and empty results are equivalent; error text/type/identity and
allocation, capacity and alias identity are unspecified. All visible input byte
values must remain unchanged, including on failure. The assessment domain is
at most 4096 input bytes; count and length bytes retain their entire 0..255 range.

The editable function is `takeText(data []byte) (string, []byte, bool)` in
package `subject`. Missing length or `N > len(data)-1` must return empty string,
no remainder and false; otherwise return exactly the next N bytes as a string,
the remaining bytes and true. Equality at the boundary is valid. The readonly
parser handles all other guards. It does not give learner repair credit for
teacher hardening of unrelated donor faults.
In particular, Decode rejects a text tag with no following length byte before
calling takeText. That path emits no helper checkpoint. Direct teacher helper
tests still exercise empty helper input and both text tags' missing-length
Decode errors, with no witness credit.

Readonly public types are `type Stage string`, constant
`TextExtent Stage = "text-extent"`, and:

```go
type Cause struct {
    Case []byte
    Stage Stage
    PayloadOffset int
    Declared int
    Available int
}
```

The external learner test package declares
`func ReproductionCause() subject.Cause`. The native run evaluates this exact
submitted function once; a fixture-side replacement or a source comment is not
evidence. Submitted code cannot invoke or alias the getter itself; a pure
private construction helper may be shared with tests. Offset is zero-based in
the complete Decode input and denotes the first byte after the length header.
Declared and Available count bytes. Stage
identifies an unchecked text extent, never a panic message or compiler PC.
Missing/default/wrong facts produce distinct public diagnostics.

The mandatory witness has C=1, initial tag 1 or 2, a present length byte,
and extent greater than remaining visible bytes. Minimize visible length,
then positive extent within that family: length three and extent one are the
two equally minimal tag alternatives. This is a structural lower bound, not
yet a runtime observation. The actual original bounds fault and same live
fixed case must subsequently be observed. Spare capacity does not alter
minimality or causal units.

A later text field/record may produce a permitted extra original fault if a
valid prefix reaches the same unchecked helper expression with an excessive
extent. It never replaces the mandatory witness. A wrapper may recover that
real, attributed original fault and return a synthetic no-records/error result
so the learner test continues. This continuation is explicitly different from
the raw original's behavior. Wrong-site faults, learner faults, missing finality,
race, timeout and trace corruption invalidate the whole run, even after a true
witness. The fixed subject has no successful fault-recovery path.

## Additive metadata and file assembly

Future activation adds `verification.reproduction`, required exactly for the
new profile and forbidden elsewhere. It is a closed object:

| Field | Meaning |
|---|---|
| `grammar` | Exactly `tagged-text-v1`. |
| `correct_variants` | Two to four unique existing mutant-syntax identifiers, naming trusted alternative correct helpers. |
| `required_rejections` | Closed mapping from every `required_mutants` id, and no others, to `behavior`, `witness`, `cause`, or `test-assertion`. |

Use `activity: debugging`, `submission.kind: workspace-patch`, exactly editable
paths `subject/repair.go` and `subject/regression_test.go`. All package selectors
equal `["./subject"]`; `fuzz_seconds: 0`, `determinism_runs: 2..3` and command
timeout 1..60 seconds retain their meanings. Race and vet apply to every
applicable native invocation. No shell command or inferred package selection.
For this finite profile require 1..16 required wrong-submission ids, each at
most 64 ASCII bytes; correct-variant ids have the same length bound.

```text
workspace/subject/scaffold.go       readonly Record and Decode
workspace/subject/cause.go          readonly Cause vocabulary
starter/subject/{repair.go,regression_test.go}
solution/reference/subject/{repair.go,regression_test.go}
solution/mutants/<id>/subject/{repair.go,regression_test.go}
verification/original/repair.go     frozen adapted buggy helper
verification/correct/<id>/repair.go trusted correct helper alternatives
verification/REPRODUCTION_SOURCE.md exact published finite contract
tests/public/subject/*_test.go      teacher behavior tests
tests/grader/subject/*_test.go       teacher additional behavior tests
```

`workspace_skeleton` equals the two readonly paths above. Every patch has both
editable paths, with no extra files. All directories use existing regular-file,
no-symlink, containment and pure-Go policies. No external modules or go.work.
Original and correct helpers expose precisely the private helper signature,
with no tests, globals, extra API or imports outside the helper grammar.
The original source identity is independently reviewed, not trusted because its
directory is called original. A content change requires source/site review.

Export readonly scaffold, Cause vocabulary, original helper as a separately
labelled reference artifact, and the complete public finite contract. Never put
the original in the compiled subject alongside the editable helper. Continue
to exclude grader, correct/reference answers, mutants and internal review.
Task-init, statements, export/reload and public/full service entrypoints must
share this contract; no hidden corpus or teacher test supplies learner liveness.

## Assessment and wrong submissions

Public checks require helper/parser behavior on teacher public cases, the actual
learner regression under original, learner-fixed and every correct-helper
binding, a minimal matched witness, correct Cause and complete successful
finality, for every repetition. Full adds teacher grader behavior cases; it
does not weaken public observation. Teacher helper tests may use same-package
access to takeText; the separately assembled learner regression remains external.
Teacher behavior executions never provide learner witness credit.
Admission precedes teacher behavior. Complete all planned behavior repetitions;
if they produce a semantic failure, replay/Cause obligations are explicitly
not-run with that dependency, rather than invented failures or successes.
Only passing behavior proceeds to the replay matrix. Within that matrix an
authentic, structurally complete wrong fixed-helper return is a public behavior
assertion failure, distinct from invalid observation. Preserve the real return
and continue the real child; never repair it by substitution. Any later native
or protocol invalidity dominates the earlier semantic failure.

The original assembly omits learner repair entirely. Every other replay uses
the identical submitted regression bytes; alternate correct variants replace
only the helper. Independence from capacity, allocation and error identity is
part of public fairness. The canonical repair is the learner's submitted one,
not a hidden reference that rescues an incorrect implementation.

The starter uses a safe, admitted, intentionally incomplete helper (for example,
always reporting failure), admitted tests and an unanswered Cause. It must
produce a complete public behavior assertion failure and complete full failure;
compilation/admission failure or runtime invalidity is not acceptable starter
evidence. The faulty adapted original is supplied separately for inspection and
reproduction. This keeps the repair exercise without making a panic the
starter-quality gate. The unchanged buggy helper is a mandatory *invalid-run
control*, not a required semantic mutant or evidence of detection.

For this profile `required_mutants` names whole wrong submissions. Each must be
admitted, compile and produce its declared semantic rejection in every full
repetition. `behavior` requires completed teacher assertion failure or a complete
replay with a failing independent public value/preservation assertion; `witness`
requires clean native replay without a matching minimal live case; `cause`
requires a real paired witness and a wrong report; `test-assertion` requires a
completed learner-native assertion failure. The mechanism defines dependency,
nonexecution and deterministic aggregation, including the exact prerequisites
for each rejection. Invalid native runs never increment detected, even when another
obligation already failed. Schema migration documents this profile-specific
meaning; other profiles keep their meanings.

Required controls include removing only the live minimal call, using only a
later-field fault, wrong case/stage/offset/extent/available, blanket error, lost
valid value/remainder, exact-boundary rejection and input mutation. Keep public
regressions with useful requirement-specific diagnostics and passing counterparts;
grader rejection alone is insufficient. Every basic prompt obligation maps to
public assertions in review notes. A public run does not disclose hidden grader
vectors or prove arbitrary unrestricted Go programs correct.

## Collection v5 and migration sequence

Advance the checker contract to `task-check/v5` only in the future activation
batch. The [mechanism design](../plans/reproduction-text-mechanism.md#collection-evidence)
specifies new reproduction evidence/counters and strict result categories.
Update quality-report and collection schemas, generated reports, freshness,
validator, coverage and all consumers together. Reject old v4 collections in
the v5 verifier; never relabel or upgrade saved evidence in place. No schema v1,
skill identifier, route address or old task content digest changes are proposed.

The sequence is binding:

1. Independently review this complete design; accept it in its own docs PR after
   exact-head green `validate`. This authorizes preparation, not execution.
2. Author a separate finite prototype and disposable launcher. Independently
   review their actual source manifests/diff and bounded proof matrix **before**
   executing raw donor, adapted parser, learner fixture or native test code.
3. Obtain actual raw/adapted/wrapped/fixed and negative-control evidence, with
   independent review. A failed attribution or grammar-fairness proof returns
   to design; do not substitute static predictions for observed faults.
4. Implement a dormant kernel separately, then profile/schema/service/collection
   activation with compatibility regressions. Each tooling PR gets ordered
   fast/race/vet/integration, one fresh bank and same-file consumers, independent
   final diff review and exact-head green validate before merge.
5. Author a separate content batch of at most ten tasks, normal task-promote,
   reference/public/grader/wrong-submission evidence, public feedback regressions,
   temporary-home export/reload/liveness, actual plan readiness, collection and CI.

The docs-only decision changes no checker inputs and may reuse the exact current
v4 collection only after an independent current-input closure proof. It must
not run an unchanged bank merely to regenerate a report. Any later verifier or
schema edit invalidates that reuse and requires its own ordered acceptance.

Existing task identities, statements, solutions, progress, history, workspaces
and current/Next/resume behavior must remain byte/behavior compatible in
temporary fixtures. Never inspect real learner home or private Git data.
Actual learner retention in version control remains an explicit **unassessed,
mandatory later outcome** of `fuzzing.reproduction`. A teacher commit or retained
workspace file is not learner VCS action. No new theory, foundation release,
350–500-task target reduction, mastery or aggregate semantic acceptance follows
from source admission, this design, its proof or a single future introduction.
