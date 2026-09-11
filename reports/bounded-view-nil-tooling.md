# P4T / T15 — bounded-view nil public regression

The disposable learner Service.Check fixture rejects a nil-only spare-capacity
defect with len=0 cap=1 / expected len=0 cap=0 diagnostics. Reference public/full
checks pass and a non-nil empty zero-capacity alternative is accepted. The existing
contract does not require the result itself to be nil.

ADR0023 adds an exact directional content digest pair. Viewing an old workspace
preserves solution, marker and progress bytes. Unknown/reverse pairs stay invalid.
Before P4C the regression prepares the reviewed assertion only in its disposable
copy; after P4C's review marker it exercises the canonical assertion as supplied.
An omission experiment must then fail from unexpected public success.

Targeted integration passed; full tooling gates and independent review remain
required. No task content, ready IDs, route numbers or real learner state change.
P4C separately supplies the public assertion, declared mutant and provenance.
