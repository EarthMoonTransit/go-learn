# Slice bounds and shared views

A slice describes a sequence backed by an array. Length tells you how many
elements are currently in the view. Capacity describes how far that view may
extend in its backing storage. Copying a slice value does not copy its elements.

For a two-index expression, the low bound is included and the high bound is
excluded. The length is high minus low. An index must be smaller than length;
a valid slice bound may reach an endpoint. Valid bounds depend on the operand:
for a slice, slicing can expose backing storage up to its capacity. Do not assume
that an index and a slice bound have identical rules.

Read the output before running this separate example:

```go
package example

import "fmt"

func Example() {
    data := []int{12, 24, 36, 48}
    middle := data[1:3]
    middle[1] = 99
    fmt.Println(len(middle), cap(middle))
    fmt.Println(data)
    // Output:
    // 2 3
    // [12 24 99 48]
}
```

The view starts at data's second element. An assignment through that view changes
the corresponding element of the backing array. This sharing is intentional;
it is different from producing an independent result. Shrinking length also does
not erase the underlying elements.

A full slice expression can additionally restrict a view's capacity. That is a
boundary on future extension, not an element copy. The linked exercise asks you
to apply this distinction; the example above deliberately uses only two bounds.
Consult the full-slice section of the specification to determine the expression
that fits the contract.

An empty slice has length zero. It may still be non-nil and may have capacity.
A nil slice has zero length and capacity. Do not infer representation from length
alone, or require a particular representation unless the task says so. Nil input
with valid zero bounds must be handled according to the stated result contract.

Common mistake: checking only element equality. Equal elements do not establish
capacity, sharing or absence of unintended mutation. When reasoning about a view,
ask which storage a write reaches and what an append is allowed to reuse.
