# ADR 0011: Editor declarations for test submissions

Status: accepted. Date: 2026-09-05.

The first ready test-writing lesson exposes an existing learner-export gap:
only the learner test file is copied into the workspace. The checker assembles
the function implementation in a disposable directory, so command verification
works but the editor cannot resolve the imported function declaration.

New test-submission workspaces receive a read-only subject/learn_api.go when the
public API can be represented using exported free-function signatures and
built-in types, including their supported collection, channel and function
combinations. The generated functions contain only a fixed panic explaining
that learn check/complete performs verification. Source bodies, comments,
constants, initializers and private declarations are never exported. Parameter
names are removed to avoid shadowing the panic built-in.

Named types, methods, generics, selector types, exported values, build-constrained
or platform-specific source files, duplicate declarations and other
unsupported public declarations retain the existing checker-only workflow.
If a built-in type name is shadowed, no stub is generated. The whole API must
fit this restricted representation; no partial declaration set is exported.
Broader API support requires a separately reviewed public-type contract.

The generated file is verified byte for byte if present and is excluded from
submission selection before execution. Extra implementation files remain
forbidden. Older workspaces without this optional editor aid remain valid;
their files, attempts, hints, completion and timer do not need migration.
New workspaces include it automatically. No task schema, progress version,
verification semantics or content digest changes.
