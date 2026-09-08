# ADR 0015: explicit foundation route numbers

Status: accepted. Date: 2026-09-08. Implements the accepted 150-lesson roadmap.

Route version 2 adds integer `number` to each block and step. Numbers are
positive, at most 999, unique and increasing within their scope; gaps are valid.
Array order remains prerequisite order. Numbers are authored identifiers for
navigation, not generated from array positions at display time. Existing numbers
must stay assigned when adding content; deliberate reassignment needs a separate
migration decision. Block IDs must be safe lowercase slugs.

The migration assigns 01–09 to existing foundation blocks, preserves all 22 task
IDs and their order and assigns existing lesson positions explicitly. English
titles no longer embed block numbers. Backend/production empty placeholders are
removed from this active route, remaining deferred in ROADMAP.md. Route v1
remains readable and unnumbered; explicit number fields
on v1 are rejected. No task/progress schema or task content changes here.

UI course lessons and task views gain an additive `number` string; UI block
titles and CLI list/show/next/roadmap show the same presentation. Task IDs remain
command/progress keys and are available in details. Ready tasks outside a route
remain visible in CLI list without an invented number. Read-only requests do not
write progress or select a lesson. Numbered workspace migration is a later ADR.

Regression tests pin existing assignments, reject missing/duplicate/out-of-order
numbers, preserve sequence and completion, and compare UI/CLI numbering. Rollback
restores route v1 and old code together; existing learner state is unchanged.
