# Retain a reproduced fuzz failure

SplitPair accepts exactly one ASCII colon and preserves both substrings, including
empty sides. Zero or multiple colons return empty, empty, false. Accept arbitrary
string bytes. The supplied implementation is faulty.

- Predict and reproduce the supplied double-colon native seed; retain its assertion and explain the bytes.
- Author TestSplitPairRegression in regression_test.go before repairing split.go. Execute it against the broken implementation and retain its useful failure.
- Repair only split.go, retaining the ordinary regression, unchanged oracle and corpus. Check the same regression, all tests/seeds and bounded native fuzz invocation.
- Explain the defect, this two-byte case's minimality, and what the exercised results establish.
- Reject a log-only regression even if it passes. Demonstrate that the unchanged oracle still rejects the broken implementation.
- Submit source and actual command evidence for checklist review.

Use the authorized disposable procedure. The supplied target earns no authorship,
discovery or executed-minimization credit. No timing threshold is graded.
