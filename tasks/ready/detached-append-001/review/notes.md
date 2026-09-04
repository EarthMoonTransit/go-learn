# Review notes

`append` may reuse spare capacity, so two slice headers can expose the same
writable array. The public test demonstrates result-to-input mutation; the
grader also checks input-to-result mutation and the no-values path. Nilness is
not part of the contract: only content, non-mutation during construction, and
storage independence are observable.

The four required mutants model distinct faults: direct append into spare
capacity, loss of the original prefix, mutation during construction, and reuse
only when no values are appended. Independent reviewers found the reference
idiomatic without making allocation strategy observable. A Phase 3C static
clean-room comparison was performed only after authoring and found no copied
upstream text, cases, code, tests, or solutions.
