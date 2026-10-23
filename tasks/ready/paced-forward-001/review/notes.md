# Author evidence: repeatedly disabled stream cases

Approved I8 automatic contract and PACE-TOKEN-01 require one select-loop worker with actual local nil channel transitions. This application disables still-open input/release/output repeatedly as pending state changes, unlike the merge first's one-way EOF disabling. Only channels.nil-disabling application is proposed. Actual source/method/evidence independently approved2026-10-23.

| Required behavior | Public assertion | Wrong specimen | Useful feedback |
| --- | --- | --- | --- |
| Prompt nonnil private unbuffered output | beginPace, TestPaceIndependentOverlappingCalls | nil-output, wait-eof-before-return, buffered-output, shared-output | output is nil; did not return before EOF; capacity actual/want; overlapping calls share output |
| One unsent value held | TestPaceRepeatedPermissions, TestPacePreservesSecondTokenWhileSendHeld | read-ahead, reverse-batch | paused input remaining actual/want; held-send remaining |
| Permission before each delivery | TestPaceRepeatedPermissions | ignore-permission | output actual before permission |
| Preserve idle token | TestPacePreservesIdleToken | consume-idle-token | idle token count actual/want |
| Preserve second token while permitted send is held | TestPacePreservesSecondTokenWhileSendHeld | consume-extra-token | held-send token count actual/want; preserved second token missing |
| Exact ordered values including zero/duplicates | repeated permissions and complete stream | skip-zero, duplicate-value, reverse-batch | value or complete stream actual/want |
| Release EOF permits all remaining values | TestPaceReleaseClosure | drop-on-release-close | release EOF closed output; complete stream actual/want |
| Input EOF preserves pending and then closes output | TestPaceInputEOFWithPending, TestPaceEmptyInputEOF, finish | read-ahead, never-close | pending value actual/want; output remains open |
| Independent overlapping calls | TestPaceIndependentOverlappingCalls | shared-output | private channel identity, B remains paused, both exact streams |

All fixture input/release channels are test-owned and finite. finish always closes both before retrieving any late constructor result, then makes forty nonblocking receive/Wait attempts. Duplicate mutant has at most16 values. Every required wrong worker terminates after those external closures and drains; never-close leaves an open channel but no active worker. No inaccessible blocked internal send, panic, timeout or leak is detection credit. Synctest Wait is not used while any worker holds a lock with contenders. The shared-output cache resets before closure and never crosses a completed test bubble.

Correct alternative encodes state with persistent receiving/granting/sending channels and sets them nil at each transition. Behavior-passing method-wrong control uses sequential range/receive/send and no select/nil disabling. Public/full acceptance of that control is intentional; actual method review rejects it. Do not claim architecture can be fully inferred from outputs.

The held-output regression intentionally supplies the second token while there is no ready receiver, then calls synctest.Wait before observing token retention. It makes no claim about a co-ready select priority. Removing exactly this test in a public-only copied fixture must let consume-extra-token survive; record the omission experiment separately. All authored execution uses official task-check disposable projects only.

## Independent acceptance

The sole reviewer checked26task files,2controls and19report hashes. Canonical13/13, race, vet and determinism pass. All13public diagnoses are ordinary assertions; all copied roots were removed. The omission experiment actually leaves consume-extra-token surviving while reference and race pass. Official promotion is accepted; final batch delivery gates remain separate.
