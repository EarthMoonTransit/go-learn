# Progressive hints

1. First distinguish nil from an error interface containing a typed nil pointer.
2. Use a comma-ok assertion before inspecting fields.
3. Test Error independently with custom fields so validation-specific strings cannot hide a formatting defect.
