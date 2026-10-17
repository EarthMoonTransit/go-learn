# Hints

## Hint 1

A generic function places its type parameter list after its name. Input elements and callback results need separate parameters; the supplied record already accepts both.

## Hint 2

The callback determines the result type during inference. Neither element type needs comparison, conversion to an interface value, or reflection.

## Hint 3

Track the input order, callback count, and the two fields independently. Nil and empty inputs should not visit an element.
