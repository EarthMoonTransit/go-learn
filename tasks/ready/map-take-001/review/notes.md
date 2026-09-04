# Review notes

The comma-ok lookup is an explicit prerequisite and receives transfer practice;
it distinguishes a stored zero from an absent key. The primary mechanism is the
safe mutation contract: remove the requested entry, preserve unrelated entries,
and rely on `delete` being safe for missing keys and nil maps. The prompt does
not require a redundant conditional around `delete`.

Public feedback exposes zero-value ambiguity and removal. The grader separately
covers missing and nil maps plus unrelated-entry preservation. The four mutants
model unconditional success, clearing the whole map, leaving the requested
entry, and treating zero as absence. A Phase 3C static clean-room comparison was
performed only after authoring and found no copied upstream text, cases, code,
tests, or solutions.
