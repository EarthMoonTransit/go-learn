# ADR 0008: Provenance v2 and original content identity

- Status: accepted
- Date: 2026-09-04

## Context

Provenance v1 permits one external-looking source and requires original tasks to
refer to the repository's own future commit. That is circular before publication
and cannot accurately describe combined sources or per-artifact use.

## Decision

Provenance v2 records zero or more audited external sources. Every source repeats
the registry id, repository, exact commit, license and usage mode, then declares a
reviewed role and artifact-level path, SHA-256 and usage classification.

Original work uses `config/project-identity.yaml`, the project license, a creation
record and an immutable digest of authored task content. It does not invent an
external source entry or require a commit that cannot exist yet. A later release
manifest binds the task id and content digest to a VCS commit.

Original project code and educational content are licensed Apache-2.0. Adapted
third-party artifacts retain their source license. `THIRD_PARTY_NOTICES` reports
those terms and does not relicense them. The machine policy permits reviewed use
of registered MIT, BSD, Apache-2.0 and CC-BY material; non-compatible or ambiguous
material remains reference-only, concepts-only or blocked.

Reference-only and concepts-only records may document factual, conceptual or
comparison evidence. Their roles and artifact classifications cannot claim copied
or adapted code, tests, structure or text. Semantic clean-room review remains a
human gate.

## Consequences

- Combined source chains are explicit and versioned.
- Original scaffolds are valid before their first release commit.
- Artifact attribution survives future transformations.
- `provenance.yaml` plus a notice file does not erase upstream obligations.
