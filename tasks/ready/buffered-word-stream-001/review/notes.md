# I10 scanner independent content acceptance

Original157-word approved contract; official task-init, original ADR0008 creation/digest and sources=[]. Sole independent source/method/evidence review is approved; see review/independent-review.json.

| Basic requirement | Public assertion | Wrong variants |
| --- | --- | --- |
| ScanWords, unchanged ordered/final words | WordOrderAndFinalToken compares exact bytes, empty writes0 | default-line-split, drop-terminal-words, drop-repeated-word, alter-word |
| Accept64 bytes, reject1024 | TokenBounds exact accepted bytes, nonnil oversize error and prefix | tiny-token-buffer, unbounded-token |
| Original reader error and accepted terminal word | ReaderErrorKeepsFinalWord compares identity and prefix | ignore-scan-error, return-before-flush-on-scan-error |
| Owned buffer below16byte output | OwnedBufferBeforeFinalFlush checks destination writes at terminal read | direct-unbuffered-writes |
| Fresh nested buffer, no caller Flush | FreshBufferAroundBufferedDestination checks raw sink/pending count then caller flushes | reuse-caller-buffer, flush-destination |
| Final Flush on all paths, write error precedence | FinalFlushErrorPriority compares original marker/ErrShortWrite and partial bytes | ignore-flush-error, scan-error-wins, ignore-short-write, omit-final-flush |
| Stop scanning after Write failure | NoReadsAfterObservedWriteFailure counts only reads after an actual failure | ignore-buffer-write-error |
| Leave streams open | Success/error Close counters | close-source, close-destination |

All18 required wrong variants are finite and diagnosed by ordinary public assertions. Sources are bounded<=16KiB; no blocking readers, panics, races, timeout or compiler failure credit. No65..1023 promise or exact oversize error text. UTF-8 fixture includes64-byte multibyte word, not64-rune confusion. Terminal bytes+reader-error preserves accepted ScanWords final token. Scanner may read ahead; no fixed read count. After-write-failure assertion only begins after actual underlying failure, allowing a larger correct buffer to defer that failure to final Flush.

## Actual method and allowed alternatives

Learner constructs/configures Scanner (ScanWords/token bounds), iterates accepted words directly into its own fresh bufio.Writer, stops after Write failure, attempts final Flush including scanning-error exits, chooses write/flush error before Scanner.Err, and never closes streams. Zero Writer.Reset(dst) is fresh; NewWriter alone may reuse dst and is wrong for ownership. A writer-only wrapper with32-byte or64KiB private buffer is valid; scanner maxima128/512 also valid. Tests avoid fixing the unspecified capacity gap. A large buffer is valid although current bounded input may fit; it still receives each word during iteration. Private bytes.Buffer staging intentionally passes finite behavior but fails actual bufio.Writer/streaming method; API use alone gives no credit. No AST extension, automatic method proof, component or interface bonus is claimed.

## Independent actual acceptance

Reviewer /root/reviewer accepted this final source and finite evidence on2026-10-23; exact immutable decision is retained in review/independent-review.json. Behavior-passing method controls are explicitly rejected by the required checklist. Official lifecycle transitions are complete; integration and all mandatory batch local/CI gates remain separate.
