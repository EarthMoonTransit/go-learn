# Append, copy and ownership

Before modifying a slice, decide whether the caller expects shared storage or an
independent result. Ownership is an API promise. A fresh slice variable alone
does not create fresh elements, and passing a slice to a function does not hand
over exclusive access to its backing array.

Append returns the resulting slice value; retain that result. If capacity is
sufficient, append reuses the existing array. When it is insufficient, append
allocates a different array and copies existing elements. The language does not
promise a universal capacity-growth multiplier.

Copy transfers elements into an existing destination and returns how many it
copied. The count is limited by the shorter length, not the destination capacity.
Copy supports overlapping source and destination. This example examines count
and independence without implementing a reserve or clone exercise:

```go
package example

import "fmt"

func Example() {
    source := []byte{'x', 'y', 'z'}
    target := []byte{'?', '?'}
    copied := copy(target, source)
    source[0] = '!'
    fmt.Println(copied, string(target), string(source))
    // Output: 2 xy !yz
}
```

Here target already owns two elements. Copy fills those elements and leaves no
connection from target back to source. For slices whose elements themselves
contain references, copying the elements is still shallow: maps, pointers and
nested slice backing arrays may remain shared. Independence must be defined at
the level the contract actually promises.

A useful ownership check changes one side after the call and inspects the other.
Another check snapshots input before the call to detect unintended mutation.
Append checks should use known starting capacities so the expected relationship
is determined by the contract, not a guessed runtime allocation pattern.

Common mistakes are appending without saving the returned slice, confusing zero
length with zero capacity, and testing only the final values. Those values can
look correct while a later caller operation corrupts shared storage. In Practice,
read the requested ownership boundary, work in GoLand and use public Check.
Return here to reason about the storage relationship instead of copying a
particular implementation pattern.
