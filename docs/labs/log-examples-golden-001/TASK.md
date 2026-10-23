# Document and review complete log output

Use the supplied Log API; do not change its v1 behavior.

- Author discoverable examples showing append/render order, rejected append without state change, and Counts map output. Choose intentional exact or unordered matching; exercise the API.
- Independently choose at least 24 entries covering all levels, empty and duplicate messages, quotes, newline, tab and Unicode. Author and review the complete expected rendering; do not derive the oracle from Render.
- Write TestLogGolden comparing all bytes with a focused first-difference diagnostic. Ordinary runs must never rewrite the baseline.
- Demonstrate useful rejection of entry loss, order, final-newline, rejection-policy and count defects; verify native example discovery.
- Review the proposed counts-footer evolution against every retained entry. Explicitly update only that approved delta; retain the actual diff and decision. Reject a candidate that loses an entry.
- Follow PROCEDURE.md in disposable copies; submit examples, baseline, tests, hashes and decisions.
