# Write a test that can explain a failure

A test connects an input to an observable requirement. Start with the contract,
choose an input that distinguishes correct behavior from a plausible mistake,
then compare actual and expected results. A passing test that would also accept
the mistake provides little evidence.

Go discovers test functions in files ending in `_test.go`. A normal test takes
`*testing.T`. `Errorf` records failure and continues; `Fatalf` stops the current
test function. Use Fatalf when later statements would be meaningless or unsafe
without the earlier condition. Report enough context to understand the failing
input without guessing from the test name.

This independent example tests a standard-library operation. It does not encode
the requirements or cases of the exercise linked below:

```go
package example

import (
    "strings"
    "testing"
)

func TestCut(t *testing.T) {
    before, after, found := strings.Cut("mode=fast", "=")
    if !found {
        t.Fatal("Cut with an existing separator: found=false; want true")
    }
    if before != "mode" || after != "fast" {
        t.Errorf("Cut: got (%q, %q); want (mode, fast)", before, after)
    }
}
```

The found flag is checked separately from the strings: different outputs may
happen to contain zero values. For your exercise, list useful partitions such
as an ordinary input, a boundary and a rejected input. Do not invent behavior
outside the stated contract. Prefer an explicit expected result over recomputing
the answer with a copy of the implementation under test.

In a test-writing task you submit tests, not the function implementation. The
mutation-tests verifier checks that your tests accept the reference and reject
required faulty implementations. A high line-coverage percentage alone does not
prove that your assertions detect those mistakes. Public Check gives feedback
without exposing hidden variants; Complete uses the full contract.

Common mistakes include logging a mismatch without failing, comparing only one
of several promised outputs, and treating every error as equivalent when the
contract requires a particular error. Use the linked task's API and requirements
to choose assertions. Return here after a failure to improve the assertion,
without revealing a hint or the reference solution.
