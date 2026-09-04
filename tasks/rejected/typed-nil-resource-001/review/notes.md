# Review notes

## Decision

Rejected. The declared skill concerns typed nil pointers inside interfaces, but
the task asks for a universal reflection-based nil predicate. Reflection is an
undeclared mechanism, and the contract promises all nil-capable dynamic values
while the grader covers only pointers and slices. The final hint prescribes the
mechanism, and the API teaches detection after the fact rather than preventing
the interface-boundary defect.

This seed is not suitable for a local patch. A future task would need a new API
design focused on returning a true nil interface or an explicit concrete-pointer
contract. No replacement is created in Phase 3C.
