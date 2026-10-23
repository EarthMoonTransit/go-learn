# Stabilize a snapshot across package initialization

Prerequisites: imports/exports, initialization-values and modules-build.

A supplied snapshot depends on registration in another init function. Edit only registry.go and z_register.go.

- Predict seed.Base and assembly's map/snapshot before the first run. Inspect actual file selection and explain package dependencies, variable initialization and init execution.
- Reproduce the original snapshot defect. Rename only the snapshot file as prescribed, predict the changed outcome and observe it; restore its name.
- Replace the hidden registration-order dependency with explicit variable initialization. Preserve seed and snapshot code and public values.
- In both prescribed file orders require seed=4, port=8, snapshot=8 and exactly one seed init and capture per process. Repeated ordinary reads must not initialize again.
- Retain wrong-order-only and lazy-read controls, useful assertions, source snapshots and actual repairs.

Use the reviewed fresh disposable procedure. Native values and actual source review are required; renaming files alone is not a repair.

