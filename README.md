# Go Practice Bank

Кураторский банк практических задач по Go. Репозиторий находится в
**ФАЗЕ 3D**. Taxonomy v3 и полный pedagogical triage Exercism зафиксированы;
добавлены lifecycle, Task Contract v2, multi-source provenance, изолированные
fixed/mutation/workspace verification profiles и локальный learner CLI.
Локальный UI включён в ближайшую поставку поверх общего learner engine;
GoLand остаётся редактором. Запуск из checkout: `go run ./cmd/learn ui`.
Инструкция и совместимость версий: [LEARNER_CLI.md](docs/LEARNER_CLI.md).
Ограниченный план:
[docs/LOCAL_UI_GOAL.md](docs/LOCAL_UI_GOAL.md). Массовая генерация не планируется.

Актуальный личный маршрут, критерии готовности и ближайшие партии работ:
[docs/ROADMAP.md](docs/ROADMAP.md).

## Быстрые команды

```bash
go test ./...
go test -race ./...
go vet ./...
go test -tags=taskintegration ./...
go run ./cmd/learn list
go run ./cmd/learn roadmap
go run ./cmd/learn next
go run ./cmd/source-sync -root . -source exercism-go
go run ./cmd/harvest -root . -source exercism-go
go run ./cmd/catalog-report -root . -source exercism-go
go run ./cmd/task-init -root . -slug <slug> -primary <skill>
go run ./cmd/task-check -root . -id <task-id>
mkdir -p .tmp/results
go run ./cmd/task-check -root . -all -output .tmp/results/task-checks.yaml
go run ./cmd/task-promote -root . -id <task-id> -to in-review -actor <name> -reason <reason>
go run ./cmd/task-report -root .
go run ./cmd/validate -root . -phase1 -checks .tmp/results/task-checks.yaml
go run ./cmd/task-report -root . -check -checks .tmp/results/task-checks.yaml
go run ./cmd/validate -root . -coverage -checks .tmp/results/task-checks.yaml
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
cmd/learn/                 локальный learner workflow без hidden artifacts
cmd/validate/             CLI-валидатор
curriculum/skills.yaml    taxonomy атомарных навыков и mastery-критерии
curriculum/migrations/    machine-readable migrations и immutable snapshots
curriculum/mappings/      versioned карты внешних concepts → skills
curriculum/evidence/      reviewed candidate-to-skill evidence
docs/                     нормативные документы проекта
docs/adr/                 решения и предложения по эволюции taxonomy
internal/isolation/       disposable Go workspaces без source-checkout execution
internal/learner/         локальные workspace, progress и scheduler
internal/tasking/         lifecycle, task gates, promotion и reports
internal/validation/      модель и проверки каталога
provenance/               правила хранения происхождения и attribution
schemas/                  task, concept-map и evidence schemas
sources/registry.yaml     реестр аудированных источников
tasks/                    drafts/review/ready/rejected/retired task packages
```

Конечное видение продукта, текущая контрольная точка, roadmap и brief для
независимого model audit зафиксированы в
[`docs/PROJECT_VISION.md`](docs/PROJECT_VISION.md).

Детали trust boundary и cache: [`docs/SOURCE_SYNC.md`](docs/SOURCE_SYNC.md).

Формат task package и жизненный цикл описаны в
[`docs/TASK_SPEC.md`](docs/TASK_SPEC.md).
Разделение activity/submission/verification и provenance v2 описаны в
[`docs/TASK_CONTRACT_V2.md`](docs/TASK_CONTRACT_V2.md) и
[`docs/PROVENANCE_V2.md`](docs/PROVENANCE_V2.md).

Изменения принимаются только через feature branch и зелёный PR check
`validate`. Локальный защитный hook и полный workflow описаны в
[`CONTRIBUTING.md`](CONTRIBUTING.md).

## Начать решать

В терминале GoLand используйте короткую команду learn, установленную локально:

    learn roadmap
    learn next
    learn start literal-bindings-001
    learn check
    learn complete

Откройте напечатанную папку Workspace как проект GoLand. В первой задаче
редактируется subject/bindings.go. У каждой следующей задачи её TASK.md
перечисляет редактируемые файлы. Вводная подготовка:
[docs/LEARNER_START.md](docs/LEARNER_START.md).

CLI сохраняет ранее пройденные задачи. На перерыв используйте learn pause,
по возвращении — learn resume TASK-ID. Полный контракт:
[docs/LEARNER_CLI.md](docs/LEARNER_CLI.md).

Для запуска без локальной команды из корня репозитория используйте
go run ./cmd/learn с теми же аргументами.
