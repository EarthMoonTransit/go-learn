# Review notes

## Decision

Returned to draft. Value-receiver mutation is the intended mechanism, but the
private `counterAdder` path is not observable through the exported API. A
learner can bypass it by updating `Counter` directly in `Add`. The value-receiver
and copy-before-interface mutants also have the same observable behavior.

## Required repair

Put the receiver defect on an observable API boundary or remove the private
interface and method-set requirement. Replace one equivalent mutant and add an
explicit independence requirement only if a test proves it. Recalibrate the
one-line repair's difficulty and time, and complete task-specific attribution
and clean-room evidence.
