# Hints

Separate the values stored directly in the struct from the object reached through its pointer. For equality, consider each field and what pointer equality observes. For comparability, inspect field types rather than one instance.
