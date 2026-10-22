# Choose constants for distinct contexts

Create limits/constants.go; its starter supplies no declarations.

- Declare Scale as exact 2^80 without premature machine representation. Derive Mask as (Scale >> 72) - 1, usable without conversion as uint8, uint16 and an array length.
- Declare Marker with value Mask as a typed uint8 constant. Preserve actual constant-expression use; a variable is insufficient.
- Predict and explain Scale's inferred-variable defaulting, Mask's int8 representability, Mask+1 assigned to uint8, and Marker implicitly assigned to uint16.
- Reproduce those four compiler failures using PROCEDURE.md. Compare explicit uint16(Marker) with the rejected implicit assignment.
- Build, test and run the unchanged positive consumers. Preserve their exact output, 255 255 2 followed by newline.
- Retain source and actual results for every reviewed wrong declaration. An untyped Marker passing the forbidden uint16 context is a contract failure despite command success.

This manual application earns no iota, automatic coverage or learner-mastery credit.
