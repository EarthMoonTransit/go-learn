# Content audit: C05–C07

Scope: thirty reviewed additions, ending at 82 learner-visible ready lessons
after C07 acceptance. Three legacy ready packages remain outside that count.
The additions are 23 substantive Exercism adaptations, six Learn Go with Tests
adaptations and one Ardan Labs adaptation; 15 implementation, nine test-writing
and six debugging tasks. The course total is 53 substantive / 24 original /
five concept-seed, with 41 implementation / 20 test-writing / 21 debugging.
49 distinct assessed skills describe course coverage, not learner mastery.

## Semantic review

Struct construction, transition copies, pointer updates and receiver reset have
different state contracts. Predicates and captured category/period filters
separate function arguments from closure capture. Fraction classification and
widened multiplication expose different float64/float32 boundaries.

Tail summation and final-element copying retain independent source operations.
Dictionary Search, Add and Update distinguish absence, insertion without overwrite
and replacement without insertion; they also practice sentinel and typed-error
forms. These are explicit practice/transfer encounters, not repeated introductions.

MethodFunctions requires usable value/pointer method expressions. AreaOf accepts
learner-defined shapes; Greeter composes arbitrary implementations. Nested
Stringer formatting forms a five-method component. Exact dynamic extraction and
type-switch descriptions differ in supported-type policy and interface priority.
Category totals distinguish unknown categories from known categories outside the
period. Withdrawal introduces failure atomicity of a shared receiver.

## Public feedback and test strength

Independent review corrected arithmetic-only method assessment and fixed-type
interface checks; dedicated defects now need the corresponding learner work.
Saved weak submissions for struct defaults, common-prefix comparisons, custom
Shape dispatch and manually constructed typed errors remain insufficient.
Reference submissions detect every required behavioral defect.

Reviewed public checks now protect repeated quantities, exact comparison,
original non-nil sentinel identity, absent-key map preservation, type-switch
overflow/priority and category filtering within the period. Numeric comparisons
consistently honor their declared tolerance and reject NaN. Saved +5e-10 valid
implementations pass public and full checks in disposable copies.

T06 public-only regressions for prior boolean/token/map-key feedback are merged.
The next separate tooling batch extends that matrix to the new repaired public
assertions; this audit does not claim that future tooling is accepted.

## Primary-domain distribution

Rows above 50% remain explicit balancing work before full foundation coverage.
A 5/3/2 batch does not by itself satisfy every domain. The partial course uses
phase1 acceptance and does not claim the final coverage gate.

| Domain | Implementation | Test-writing | Debugging | Highest share |
|---|---:|---:|---:|---:|
| arrays | 1 | 0 | 0 | 100% |
| control-flow | 4 | 1 | 5 | 50% |
| errors | 3 | 2 | 1 | 50% |
| expressions | 4 | 2 | 0 | 67% |
| functions | 2 | 3 | 2 | 43% |
| interfaces | 2 | 2 | 1 | 40% |
| maps | 3 | 2 | 0 | 60% |
| methods | 2 | 0 | 0 | 100% |
| pointers | 2 | 1 | 0 | 67% |
| slices | 3 | 1 | 5 | 56% |
| strings | 5 | 1 | 2 | 62% |
| structs | 1 | 2 | 0 | 67% |
| types | 5 | 0 | 3 | 62% |
| unicode | 4 | 3 | 2 | 44% |

Continue toward 150. Upcoming source selections include array/slice/constant/map
test-writing, while original gaps extend defer, ownership, wrapping, generics
and I/O. Preserve the source quota and make diversity deficits visible instead
of weakening requirements. No new real learner feedback was supplied during
these batches; no calibration, completion or mastery is inferred from fixtures.

Independent audit review is pending; task-level review is already approved.
