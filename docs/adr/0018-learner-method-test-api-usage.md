# ADR 0018: learner tests may use receiver methods and public fields

Status: accepted for implementation, 2026-09-08

## Context

The mutation-test source screen recognized only calls qualified directly by a
task-package import. A valid submission constructing `subject.Dictionary{}` and
calling `d.Add(...)` was rejected before execution. This prevents ordinary
method tests and field assertions despite their supported Go APIs.

## Decision

Recognize a reference to an imported task-package selector in a source declaration or function
as API usage. This covers free functions, constructor/type references, aliased
imports, method expressions, pointer receivers and public-field assertions.
An import alone remains insufficient. This is a structural relevance screen,
not an assertion-strength proof: reference success and detection of every
required mutant remain the behavioral gates. A type-only use cannot earn
completion merely by passing this source screen.

No new execution profile, dependency permission or bypass is introduced.
Unconditional failures, TestMain exits, payload boundaries and disposable
execution restrictions retain their existing checks.

## Compatibility and migration

Task Contract v2 fields, schemas and learner progress versions do not change.
Previously accepted direct-call tests remain accepted. Previously rejected
valid method/field tests are retried without rewriting submissions or resetting
progress. The missing-API diagnostic now says "reference" rather than "call".
No stored-data migration or content renumbering is needed. Synthetic source
regressions cover receiver/type forms and negative no-usage cases; the existing
mutation integration suite still enforces behavioral sufficiency. Content using
the new acceptance behavior ships only after this separate tooling batch.
