# I9 pipe author precheck

Contract approved in stream-pipe-contract-review.json. Approved TASK wording unchanged; explicit constraint bullet yields158 words. Author metadata enables targeted disposable checks, not independent acceptance. Exchange and bounded-copy remain unchanged first preparations pinned5781a38; this task owns the coupled bounded-consumer termination policy.

| Requirement | Public assertions | Wrong variants / ordinary diagnostic |
| --- | --- | --- |
| Limit applied before copying; zero cap | CapAndProducerJoin; producer accepted prefix; dst bytes/count/no zero Write | unlimited-copy, zero-means-unlimited, limit-after-copy |
| Short EOF success; exact cap does not probe terminal | EOFAndProducerTerminal short/exact rows | copy-exact-count: copyErr EOF vs nil |
| Writer error reaches copy; callback result preserved | EOFAndProducerTerminal partial terminal marker | drop-writer-close-error, drop-producer-error, swallow-copy-error |
| Consumer outcome releases unfinished writer | CapAndProducerJoin and DestinationFailurePropagation | missing-reader-close: named producer/completion failure before rescue; drop-reader-close-error: exact Write error mismatch |
| Writer closure permits read completion | empty/short/producer-error rows | missing-writer-close: named completion failure before rescue |
| Destination partial failure/short nil error | DestinationFailurePropagation marker and ErrShortWrite | accept-short-write, conflate-errors: independent exact error mismatch |
| Producer joined | callback held after Write interrupted, before its own terminal result | return-before-join: returned before callback release |
| Caller stream ownership and once callback | relayValues/finishRelay all rows | close-dst, repeat-producer |

Public compares exact bytes, accepted output count, exact callback/copy/Write error identities, and pipe total accepted <=cap without specifying read chunk sizes. The source may deliver more bytes to the pipe reader than dst accepts. At cap, copyErr nil coexists with a nonnil joined producerErr because no subsequent read is performed. EOF before cap is successful. Grader expands binary/cap16/zero output/source EOF cases.

## Finite rescue and method boundary

Callbacks always capture the actual interface passed by the implementation in a bounded channel. The test first makes a useful assertion, releases its independent callback gate, then only on failed progress attempts io.Closer on captured writer and closes it. Closing that actual endpoint releases pending pipe reads/writes; finishRelay then joins through the call result. Hidden-Closer correct wrapper passes without rescue and receives no type restriction. Broken inaccessible hidden-closer omissions are method-only, never required timeout evidence. Known missing-writer-close and missing-reader-close specimens use short/empty and longer-than-cap cases respectively, with exact accessible endpoints. No Panic/Fatal bypasses release; no time thresholds or new runner.

Reference uses actual Pipe, callback worker, limited reader before io.Copy, CloseWithError at both ownership boundaries and exact joined callback result. io.LimitedReader and wrapper-hiding-Closer controls must pass full/public checks. Actual method review verifies this composition and absence of full-input staging independently from API/result assertions. Only pipe-lifecycle and copy-limit-composition applications are proposed; no supplied io interface, error, goroutine or component credit.

Original ADR0008 provenance sources=[] and task-init creation; canonical/public-only fixtures and controls execute only through task-check disposable copies. Independent actual method/evidence review, promotion and root batch gates remain pending.

Sole actual source/method/evidence review approved2026-10-23 without findings;44accepted fixtures with32ordinary public defects across both owned units. Finalbatchdelivery separate.
