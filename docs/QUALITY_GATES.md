# Quality gates

Lifecycle использует состояния `draft`, `in-review`, `ready`, `rejected`,
`retired`. Разрешённые переходы определены ADR 0004, hardened semantics — ADR
0005.
`task-promote` запрещает прямой `draft → ready` и не перезаписывает target.

## Gate 1 — package, metadata и taxonomy

- все Phase 3A schemas и semantic checks проходят;
- id уникален, совпадает с directory и имеет стабильный numeric suffix;
- lifecycle, directory и `quality_status` согласованы;
- primary/secondary/prerequisites/new concepts не конфликтуют;
- все skills active/assignable, прямые prerequisites явно доступны;
- `TASK.md`, package paths и source path не допускают traversal;
- symlinks и нерегулярные файлы отклоняются;
- `go.mod` имеет module `task.local/<id>`, `go.work` и внешние зависимости
  запрещены.

## Gate 2 — legal и provenance

- source metadata точно совпадает с registry и exact commit;
- usage mode допускает transformation type;
- полная цепочка и transformation boundary записаны в `provenance.yaml`;
- attribution существует, notices сохранены;
- `reference-only` код или текст не скопированы и не перефразированы близко;
- reviewer отдельно подтверждает license и clean-room evidence;
- неоднозначность ведёт в `rejected`, а не маскируется.

Автоматический validator проверяет metadata, но не объявляет semantic detection
копирования полностью автоматизируемой.

## Gate 3 — instructional design

- один наблюдаемый primary skill и не более двух new concepts;
- prerequisites не образуют скрытый скачок;
- `TASK.md` не длиннее 160 слов, context не длиннее двух предложений;
- API точный, requirements содержат 4–7 проверяемых пунктов;
- constraints явные, сюжет и объяснение решения отсутствуют;
- reviewer подтверждает prompt/test alignment и отсутствие solution leakage;
- задача не является cosmetic reskin существующего решения.

## Gate 4 — executable correctness

- starter + public компилируется, но test run падает;
- starter + public + grader компилируется, но full run падает;
- reference + public + grader проходит два повторных запуска;
- наблюдаемые test events повторных reference runs одинаковы;
- каждый обязательный mutant компилируется и обнаруживается grader;
- public API starter/reference/mutants совпадает;
- tests не используют uncontrolled clock, randomness, network, environment или
  external processes; `time.Sleep` запрещён;
- concurrency reference дополнительно проходит `go test -race`.

Варианты собираются checker-owned способом: implementations не содержат tests,
а canonical public/grader tests нельзя заменить solution binding. Все команды
выполняются в disposable workspaces с минимальным allowlisted environment,
offline module resolution и отдельными home/cache/temp paths.

## Gate 5 — review и maintainability

- `review/review.yaml` содержит reviewer, date, decision и minimum mutants;
- story, alignment, API, provenance, license, clean-room и leakage evidence
  подтверждены явно;
- `go fmt`, `go vet` и repository validation зелёные;
- public export исключает grader, reference, mutants и internal review;
- commands и reports детерминированы.

## Machine-readable status

Одиночный `task-check -id` выдаёт документ
`schemas/task-quality-report.schema.json`; `task-check -all` выдаёт versioned
`schemas/task-check-collection.schema.json`. Каждый gate имеет `pass`, `fail`
или `not-applicable`; mutant summary фиксирует
`required`, `total`, `detected` и `survived`. Дополнительно фиксируются
technical/review status, promotion eligibility, failing gate IDs, starter,
reference, race и legacy compatibility.

Collection фиксирует task input digests и структурные execution counters без
timestamps. `task-report` детерминированно создаёт `reports/task-quality.yaml`
и Markdown из validated collection.
После появления десяти lifecycle packages aggregate gate требует минимум три
task types и не более 60% одного типа. Curriculum policy остаётся строже: не
более 50% primary-встреч одного типа на домен.

## Repository acceptance

Authoring infrastructure и неполные content batches:

```bash
go test ./...
go test -race ./...
go vet ./...
go test -tags=taskintegration ./...
go run ./cmd/task-check -root . -all -output .tmp/results/task-checks.yaml
go run ./cmd/validate -root . -phase1 -checks .tmp/results/task-checks.yaml
go run ./cmd/task-report -root . -check -checks .tmp/results/task-checks.yaml
```

Release coverage дополнительно использует `-coverage`. Ни `-phase1`, ни наличие
upstream tests не ослабляют gates уже существующих packages.
