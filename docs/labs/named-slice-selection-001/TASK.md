# Preserve named slices

## Context
Select values without losing container types.

## API
Below has type parameters S,E and signature `(input S, limit E) S`.

## Requirements
- Key admits exactly int/string and their defined types using union and approximation.
- S admits ordinary/defined slices of E. Preserve static S and its methods in the result.
- Return fresh storage containing values strictly below limit, preserving order/duplicates. Input remains unchanged; nil yields nil.
- Constraints justify generic comparison and slice operations. Implement without fixed concrete conversions.
- Run ordinary/named/alias clients and semantic checks. Explain rejected element kinds, arrays, mismatched S/E and Key values.
- Submit source, actual output and type-set reasoning; distinguish compile-time admission from runtime behavior.

## Constraints
- At most eight elements; integers in [-9,9], strings at most eight bytes, ordered lexicographically.
- Standard library only; follow the reviewed disposable procedure.
