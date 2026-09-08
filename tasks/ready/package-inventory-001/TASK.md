# Implement exported fields and private state across packages

## Context

Complete the specified operations.

## API

```go
// stock: Item has exported Name string, Weight int; private onHand, sold int.
func New(name string, weight int) Item
func (i *Item) SetOnHand(n int)
func (i *Item) SetSold(n int)
func (i Item) OnHand() int
func (i Item) Sold() int
// subject
func Describe(item stock.Item) string
```

## Requirements

- Complete stock/item.go and subject/describe.go. New stores Name and Weight with both private counts zero.
- Each setter changes only its matching count; getters return current counts. Inputs are nonnegative integers.
- Describe returns "Name:Weight:OnHand:Sold" with decimal integers, using stock exports and methods.
- Preserve Name exactly, including empty text and spaces; separate Item values have independent state.

## Constraints

- Use only the Go standard library.
