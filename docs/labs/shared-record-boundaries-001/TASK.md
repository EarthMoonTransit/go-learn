# Relocate shared data and choose test boundaries
## Context
The supplied store/codec graph contains an import cycle.
## Requirements
- Predict and reproduce the cycle; identify both ownership edges.
- Move Entry to record, preserve store.Entry as an alias, and move Preview composition to cmd/preview. Libraries depend only on record, never on each other; record depends on neither.
- Preserve supplied Load/Encode behavior and exact command output. Do not duplicate Entry or export normalize.
- Author a private normalize test and an exported Load/Entry consumer test. Choose each package declaration from its observation boundary and explain it.
- Retain pre/post files, native diagnostics, repaired graph, passing tests and command output; demonstrate external private-access rejection.
## Constraints
- Edit only the six listed files; preserve supplied algorithms, module, public checks and fixtures. Use the finite offline procedure and source/method review. Manual evidence adds no automatic completion.
