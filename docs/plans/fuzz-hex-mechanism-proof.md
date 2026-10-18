# Private hex mechanism evidence for ADR 0032

This is a reviewed engineering experiment, not an activated verification
profile or a delivered task. Runs used Go `go1.27.0` and disposable isolation
workspaces. No authored learner test ran in a source checkout. Retained raw
evidence is under ignored `.tmp/fuzz-native-spike/hex-v0/`; the receipts below
bind sources, tools, requests, rewrites, native output and results. The generated
native binaries were removed with their disposable workspaces, so their saved
binary hashes cannot now be independently re-read. A later implementation PR
must supply committed reproducible regressions and its own acceptance evidence.

## Source and scope

The [P20 source decision](../../catalog/intake/2026-10-18/p20-hex-source-review.json)
selects registry source `golang-go`, exact commit
`2c31595783c74564c1299a5401dc6e606f379cc7`, artifact
`src/encoding/hex/hex_test.go`, SHA-256
`b59de7b29f9b3d0bd75d661666d57fed17d87574b79f5a299a92f9c0f88123bb`.
It is a concept/API seed, not an upstream native-fuzz exercise or substantive
adaptation. The experiment copied 16 of this project's own synthetic codec
tool files, then extended them for hex. It copied no upstream code, tests,
vectors or educational text. Its copy manifest hash is
`bfe63dc4b008eac8a9cb9f58e1b3a8141b15d40884c4ab2de5da743d690faa29`.
Future P20 provenance must retain that external concept-seed role; this proof
does not convert the task into an original with an empty source list.

The private subject path is `stageb.local/subject`; production must use the
generated task module identity. The experiment fixes target `FuzzHex` and
observer protocol integer version 3. That private protocol number is distinct
from the proposed production `task-check/v4` execution contract.

## Results and immutable receipt hashes

| Relative evidence path | Observed result | SHA-256 |
|---|---|---|
| `analysis-results-v0/receipt.json` | 61 expected compile/admission outcomes: all originals compile, 28 admitted and their rewrites compile, 33 rejected. No native learner execution in this stage. | `994c0d97d9ed8973b9cf61114a6dd898913f7fedae37ba137b126f13f0b89a1e` |
| `trace-controls-v0/receipt.json` | 67 offline malformed/protocol controls match expectations. | `0dbdb5173087da6931eb65df6880f7cf60cc1a2bf23403ae553b683dd1f9059a` |
| `native-results-v0/receipt.json` | All 305 native controls match expectations; unfiltered run, disposable environment removed. | `461af6b91c5aae0213612a40467871b36f2fc9990334f88ffc8b5b3b637fb072` |
| `native-results-v0/hash-audit.json` | 6457 retained linkage/hash checks, zero failures. Removed binaries excluded from rehashing. | `21764812f559d211f3780c80d34a364a3c3d0d692dce17d33d79bf81b4f29506` |
| `native-results-v0/offline-replay.json` | Revalidated all 305 saved helper responses and native outcomes; no native/compiler/learner execution. | `5e4ee168b667d970cdd53b2dd0249a047bacc5ed40c4544ee4605db1eb90b168` |
| `analysis-results-v0/trace-size-bound-v1.json` | Conservative 13,620,355-byte/5314-record upper bound, below 16 MiB/8192 records. | `51ee1438c95a93a910a643634bbee3d57c858447d022853054282cac7f3a9705` |
| `analysis-results-v0/sizing-v1-binding.json` | Binds corrected size tool, observer and bound. Supersedes the old assumption text without relabeling the old receipt. | `522fa3b9447f11e296351fd34bf176cfd6001c772671867cbd8e591421ee477f` |

The first size receipt retained a copied `<=4098` description while its
calculation already used 8192. The immutable v1 receipt corrects that description;
the numerical upper bound did not change. The superseded receipt is preserved.

Native driver SHA-256:
`fb2b20d7ca51be89a0718b13031baf74983915a5875d7f8826fb70508b8f2564`.
Strict trace validator SHA-256:
`0b8f17d943e1c4eefba5be701d2cafff3910f32a9700e518c3ab118d4c9f1a05`.
Observer SHA-256:
`19bfa608c723b0c0b904a896d9fad34fe08e0ad63a4ec0cac08256292f5c34f8`.

## What the controls distinguish

The native run includes 168 positive repetitions: 28 accepted forms against
three correct implementations, twice each. Forms include direct/named native
callbacks, aliases, helper/tuple returns, clones, per-index lowercase wire
checks and the independent `hex.AppendEncode` oracle. Correct implementations
use the library, a manual loop and an append-based alternative, with both nil
and non-nil empty decoder outputs. Value-equivalent correct alternatives pass.

Seven learner seeds include nil and empty forms, zero, high-bit, order,
repetition and a lowercase hex letter. The `0x0a` letter seed is below 128, independently
testing letter handling without replacing high-bit coverage. Eleven core value
defects include encode/decode zero and high-byte loss, reordering, truncation,
false decode status, paired uppercase and paired XOR codecs. A roundtrip-only
harness survives paired wrong codecs. Removing the only letter-producing seed
lets paired uppercase survive; the independent lowercase wire oracle detects
it when that learner seed is present. Diagnostic-only teacher probes never
rescue a surviving learner submission.

Missing Decode, dead helpers, early Fatal/Error-return, weak status/value-only
checks and seed-class removal have separate controls. Encode failure asserted
before Decode is legitimate; no-assertion missing Decode is ineligible. Native
panic, unexpected Goexit, unexplained exit 1, race exit 98, timeout, I/O/capture
failure, excess calls/assertions and overlarge API outputs remain invalid with
no Detection even after a genuine earlier assertion.

The largest eligible trace observed was 10,075,208 bytes/5314 records: 64
callbacks, 512 SUT calls and 4096 assertions. Maximum assertion-free correct
and race-correct traces were 9,314,504 bytes/1218 records. A 65th assertion per
callback invalidated the run (10,087,102 bytes/5378 records) without Detection.
These are observations of bounded fixtures, not general Go resource guarantees.

An independent read-only reviewer (`codec_static_review`, 2026-10-18) approved
the exact native receipt after independently checking 6171 actual API returns,
106325 semantic checks and 2322 additional checks, with zero mismatches. The
review did not rerun learner-native code. Approval covers the private finite
mechanism only. It grants no ready-task, source-adaptation, profile-activation,
reproduction, numeric coverage or learner mastery credit. Public feedback,
schema/executor integration and existing-profile/history preservation remain
mandatory implementation gates in [ADR 0032](../adr/0032-finite-native-fuzz.md).
