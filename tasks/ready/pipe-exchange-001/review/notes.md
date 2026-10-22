# Finite-v1 review

Contract and actual content/method/execution independently accepted by /root/reviewer2026-10-22. Final batch gates remain separate.

| Requirement | Public assertion | Defect/method |
| --- | --- | --- |
| callback invocation once, ordered byte transfer and empty EOF | TestEmptyAndCompleteTransfer; finishExchange callback counts | duplicate-producer/duplicate-consumer: callback calls and byte diagnostics |
| synchronous writer waits for reader | TestWriteWaitsForReader held readGate | actual io.Pipe method review; no behavioral result alone certifies implementation mechanism |
| producer error closes writer and reaches consumer unchanged | TestProducerFailureReachesReader | drop-writer-close-error: reader terminal error; missing-writer-close: completion diagnostic before rescue |
| consumer error and early nil close reader, releasing unfinished write | TestEarlyConsumerErrorAndJoin; TestEarlyConsumerNilAndJoin | drop-reader-close-error: blocked Write error; missing-reader-close: completion diagnostic before rescue |
| join producer before Exchange returns | both early consumer tests hold producer after partial write | return-before-join: returned while producer callback was still held; gate released and worker joined by quiescence |
| both exact callback errors, including independent partial-transfer errors | checkExchangeErrors in every test | drop-producer-result/drop-consumer-result: actual and exact expected error |
| actual ownership and endpoint interface flexibility | reference source; separate correct hidden-Closer wrapper control | Close assertion is only failure rescue, never a passing-type requirement; no panic/timeout detection credited |

Every behavioral mutant must fail the public-only copied fixture with a useful assertion. Timeout, panic, compile error or cleanup failure alone does not establish detection. Method checks inspect actual code; signatures alone earn no mechanism credit.

Independent actual method and distinct-action acceptance: Actual synchronous Pipe, writer and reader CloseWithError paths, callback-error identity and producer join accepted. Correct wrappers hiding io.Closer pass. Failure cleanup follows named assertion; no timeout/panic credit.
