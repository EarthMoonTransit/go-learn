# Task packages

Lifecycle task package определяется каталогом:

```text
tasks/
├── drafts/      # draft
├── review/      # in-review
├── ready/       # ready
├── rejected/    # rejected
├── retired/     # retired
└── _template/
```

Полный Phase 3A package имеет структуру:

```text
tasks/<state>/<task-id>/
├── go.mod
├── task.yaml
├── lifecycle.yaml
├── TASK.md
├── ATTRIBUTION.md
├── starter/
├── tests/
│   ├── public/
│   └── grader/
├── solution/
│   ├── reference/
│   └── mutants/<mutant-id>/
├── review/
│   ├── review.yaml
│   └── notes.md
├── hints/HINTS.md
└── provenance.yaml
```

`task.yaml` остаётся schema v1; точное lifecycle-состояние хранится отдельно и
совпадает с каталогом. Каждый package — отдельный stdlib-only Go module
`task.local/<task-id>`. Public/grader каталоги содержат reusable suite и binding
к starter; reference и каждый mutant содержат собственный binding к обеим
suites. Три исторические ready-задачи остаются legacy v1 и не переписываются.

Создание и проверка выполняются командами из корня репозитория:

```bash
go run ./cmd/task-init -root . -slug <slug> -primary <skill> \
  -source <source-id> -source-path <path>
go run ./cmd/task-check -root . -id <task-id>
go run ./cmd/task-promote -root . -id <task-id> -to in-review \
  -actor <name> -reason <reason>
go run ./cmd/task-report -root .
```

`task-check` никогда не запускает authored code из source checkout: он делает
disposable copy. Правила переходов и isolation зафиксированы в
`docs/adr/0004-task-lifecycle-and-isolation.md`.
