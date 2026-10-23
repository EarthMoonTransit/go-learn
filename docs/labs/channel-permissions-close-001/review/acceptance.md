# Author observations submitted for independent review

The sole reviewer approved the exact source, manifest and fixed recipe before
execution. The author then ran that unchanged recipe exactly once: 11 fresh
private offline copies and 40 commands, all matching their reviewed expectations.
The command process exited 0. Every disposable root was removed; no command timed
out or reported an uncaught runtime failure. Maximum command duration was
2.492308 seconds. Local version: `go version go1.27.0 darwin/arm64`.

Actual normalized argv, private environment, source and prediction snapshots,
stdout/stderr, exits, durations and cleanup are retained in `execution.json`.
Raw evidence and the author summary remain in the batch `channel-execution.json`
and `channel-author-evidence.json`. All 54 snapshots have verified source hashes;
normalization was independently reconstructed by the author as a data-only check.
All 24 content files, the manifest and the approved execution script are unchanged.

## Predictions compared with actual observations

| Requirement or control | Actual observation |
| --- | --- |
| Empty and buffered closed-channel observations | Reference tests passed nil, empty, queued-zero, mixed and four-value inputs. The exact demo preserved `(0,true,true)` for queued zero, followed by two `(0,false,true)` closed-empty readings. Both actual post-close attempt flags were true. |
| Producer send-only permission | Native build rejected receive in the learner-selected producer helper: `cannot receive from send-only channel`. Deleting that inserted operation restored build and tests. |
| Receiver receive-only permission | Separate native builds rejected send and close in the learner-selected receiver helper. Each single-statement repair restored build and tests without widening types. |
| Missing initial close | Ordinary assertions reported `Ready:false` where closed-empty `Ready:true` was required, and both false panic flags. Nonblocking operations kept the specimen finite; restoring the reference passed. |
| Premature buffer discard | Ordinary indexed reading assertions showed lost buffered values, including queued zero. Reference repair passed. |
| Zero mistaken for closure | Ordinary indexed assertion showed `OK:false` for queued zero instead of `OK:true`. Reference repair passed. |
| Omitted owner attempts | Both false flag assertions failed ordinarily; reference repair passed. |
| Wide helper parameters | Tests and demo passed; all three illegal-direction probe operations compiled under the wider parameters. This is a required source/method rejection, not an automatic failure. |
| Hardcoded true flags | Tests and demo passed. The source omits the actual owner operations, so method review must reject it. |
| Correct alternate | Named directional types, different helper names, a larger sufficient buffer and indexed result assembly passed tests and exact demo. |

The separate reference predictions and per-specimen prediction events precede the
relevant commands in retained evidence. Their predicted outcome families match
these actual results. Compiler probes were built only; their inserted illegal or
wide-permission operations were never executed.

## Actual source/method handoff

The reference authors `publish(dst chan<- int, ...)` and `take(src <-chan int)`;
the starter supplies neither helper signature. `Observe` creates the sufficient
buffer, keeps its own bidirectional handle, calls the producer to fill and close,
and performs the required number of actual nonblocking comma-ok receives.
The queued value and `OK`/`Ready` fields distinguish data, closed-empty, and
open-empty states. Owner-side send and second close occur in two separate supplied
recovery calls. There are no goroutines, timers or unbounded operations.

The named-type alternate is permitted. Broad parameter controls and hardcoded
flags deliberately demonstrate why runtime success cannot replace source/method
assessment. Directional types restrict available operations; they do not establish
exclusive ownership or change buffering. Supplied records, recovery and formatting
receive no learner authorship credit. No panic-recovery, lifecycle or component
credit is claimed.

This is the author's interpretation and evidence handoff, not independent
acceptance, automatic coverage or final readiness. The sole reviewer must assess
the retained actual method and evidence. Shared route/backlog/integration and
release commit binding remain root-owned.

## Subsequent independent acceptance

The sole reviewer accepted actual source/method/native evidence on2026-10-23.
See independent-review.json: all24assets,40commands and54snapshots matched;
eight expected nonzero exits were actual starter/ordinary/compiler failures.
Two manual applications are accepted as content evidence. Final batch gates and
delivery remain required; no automatic coverage or learner mastery is implied.
