# Go Practice Bank

Кураторский банк практических задач по Go. Репозиторий находится в ранней
**ФАЗЕ 3A**. Taxonomy v3 и полный pedagogical triage Exercism зафиксированы;
добавлена infrastructure для draft/review lifecycle, изолированной проверки и
atomic promotion. В репозитории остаётся ранее созданный ручной batch из трёх
ready-задач; ФАЗА 3A не создаёт новых учебных задач. Массовая генерация и
пользовательский интерфейс отсутствуют.

## Быстрые команды

```bash
go test ./...
go run ./cmd/validate -root . -phase1 -run-commands
go run ./cmd/source-sync -root . -source exercism-go
go run ./cmd/harvest -root . -source exercism-go
go run ./cmd/catalog-report -root . -source exercism-go
go run ./cmd/task-init -root . -slug <slug> -primary <skill> -source <source-id> -source-path <path>
go run ./cmd/task-check -root . -id <task-id>
go run ./cmd/task-check -root . -all
go run ./cmd/task-promote -root . -id <task-id> -to in-review -actor <name> -reason <reason>
go run ./cmd/task-report -root .
go run ./cmd/task-report -root . -check
go run ./cmd/validate -root . -coverage -run-commands
```

Последняя команда станет зелёной только после наполнения банка в следующих фазах.

## Карта репозитория

```text
adapters/                 контракты будущих source-adapters
catalog/candidates/       метаданные сырых кандидатов (не готовые задачи)
cmd/source-sync/          pinned cache synchronization
cmd/harvest/              offline candidate extraction
cmd/catalog-report/       deterministic import/taxonomy reports
cmd/task-*/                authoring, isolation, promotion и quality reports
cmd/validate/             CLI-валидатор
curriculum/skills.yaml    taxonomy атомарных навыков и mastery-критерии
curriculum/migrations/    machine-readable migrations и immutable snapshots
curriculum/mappings/      versioned карты внешних concepts → skills
curriculum/evidence/      reviewed candidate-to-skill evidence
docs/                     нормативные документы проекта
docs/adr/                 решения и предложения по эволюции taxonomy
internal/isolation/       disposable Go workspaces без source-checkout execution
internal/tasking/         lifecycle, task gates, promotion и reports
internal/validation/      модель и проверки каталога
provenance/               правила хранения происхождения и attribution
schemas/                  task, concept-map и evidence schemas
sources/registry.yaml     реестр аудированных источников
tasks/                    drafts/review/ready/rejected/retired task packages
```

Детали trust boundary и cache: [`docs/SOURCE_SYNC.md`](docs/SOURCE_SYNC.md).

Формат task package и жизненный цикл описаны в
[`docs/TASK_SPEC.md`](docs/TASK_SPEC.md).

## Начать решать

Каждая задача — отдельный Go module. Работайте с копией, чтобы не изменять
canonical package в репозитории. Например:

```bash
scratch="$(mktemp -d)"
cp -R tasks/ready/variadic-sum-001 "$scratch/task"
cd "$scratch/task"
go test ./tests/public
```

Первый запуск public tests ожидаемо падает: исправляйте только файлы в
`starter/`, затем повторяйте команду до зелёного результата. Не открывайте
`tests/grader/` и `solution/`, если хотите решить задачу без подсказок.
