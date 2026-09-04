# Review notes

## Decision

Returned to draft. Caller-defined formatting proves open interface dispatch,
but the starter already gives `Pair` the required method set. The learner does
not create implicit satisfaction, and two mutants cover string assembly rather
than the declared interface skill. Nil and typed-nil formatter behavior is also
undefined.

## Required repair

Reclassify this as focused interface dispatch or redesign the editable code so
method-set satisfaction is learner-owned. Replace unrelated mutants, remove
`interfaces.small-contracts` unless API design is observable, state the nil
precondition, shorten the estimate, and add exact source attribution plus
clean-room evidence.
