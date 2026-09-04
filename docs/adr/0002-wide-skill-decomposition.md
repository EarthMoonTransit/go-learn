# ADR 0002: Decomposition of broad fundamental skills

- Status: superseded by ADR 0003
- Date: 2026-09-04
- Target version if accepted: taxonomy v3 or later

## Context

The Exercism audit found current ids that combine independently teachable and
testable mechanics. Changing their meaning in taxonomy v2 would silently alter
existing manifests, so this ADR records proposals without applying them.

## Proposal

| Area | Current broad skill or overlap | Candidate split/merge | Evidence examples |
|---|---|---|---|
| Slices | `slices.model`, `capacity-bounds`, `append-aliasing`, `growth-allocation` overlap | construction/nil/empty; len/cap/slicing; append/reallocation; aliasing/defensive copy; copy/overlap; preallocation | `card-tricks`, `anagram`, `list-ops` |
| Strings/bytes/runes/UTF-8 | construction and conversion are coupled; rune model and UTF-8 codec are coupled | string construction; string↔`[]byte`↔`[]rune`; strings package; formatting; rune model; UTF-8 encode/decode | `welcome-to-tech-palace`, `logs-logs-logs`, `micro-blog` |
| Methods/pointers | receiver semantics, method sets and receiver choice overlap | receiver declaration/mutation; method sets; method values/expressions; merge receiver-choice into receiver semantics | `jedliks-toys`, `high-scores`, `election-day` |
| Interfaces | dynamic representation is coupled with typed nil; assertion with switch | dynamic type/value; typed nil; assertion; type switch | `sorting-room`, `airport-robot`, `paasio` |
| Errors | creation/propagation and wrapping/inspection are paired | creation/sentinels; propagation/context; wrapping chains; `Is`/`As`; typed errors; panic/recover boundary | `the-farm`, `error-handling`, `nth-prime` |
| Testing | each current id joins two lifecycle/tooling mechanics | basics and diagnostics; table tests; subtests; helpers; cleanup; parallel tests; race tooling; examples; golden files | no direct Exercism curriculum label |
| Goroutines | ownership/results and leaks/panics are paired | state ownership; result/error collection; leak prevention; panic boundary | `parallel-letter-frequency` |
| Channels | operations/close, range/drain, select/timeout/nil, and pipeline/fan/cancel are paired | send/receive/blocking; close ownership; range/drain; directional channels; select readiness/default; timeout; nil disabling; pipeline; fan-out/in; cancellation | `robot-simulator`, `parallel-letter-frequency` |
| Sync | race semantics are implicit | add data races/happens-before; retain mutex/RWMutex/once/cond/pool/wait-group as separate mechanics | `bank-account`, `paasio` |
| Context | current decomposition is mostly sound | optionally add cancellation causes; do not change without non-Exercism evidence | no Exercism evidence |

## Alternatives

1. Keep all current skills permanently. This minimizes migration but leaves
   ambiguous coverage.
2. Change existing ids in place. Rejected because old task metadata would gain
   a different meaning without review.
3. Add replacement ids, migrate manifests explicitly, then retire old ids.
   This is the preferred path if the proposal is accepted.

## Required migration before acceptance

- inventory every ready task referencing an affected id;
- define old-to-new mappings with `one-to-one`, `split`, or `merge` status;
- review task evidence rather than assigning all split children automatically;
- update coverage baselines and sequencing;
- preserve retired ids until all manifests and reports have migrated;
- land taxonomy/tooling separately from content batches.

No operation in this ADR is applied by taxonomy v2.
