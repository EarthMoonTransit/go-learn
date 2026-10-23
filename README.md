# Go Practice Bank


Active roadmap: **226 ready Go lessons → finite v1 course**, using existing checks plus explicit checklist review. Autonomous reviewed batches are authorized; native manual units use the explicit isolated procedure. See [finite-v1](docs/plans/finite-v1.md) and [ROADMAP.md](docs/ROADMAP.md). Full-bank coverage and release remain separate.
Кураторский банк практических задач по Go: **226 уроков в педагогическом порядке**,
со стабильными адресами в разделах 01–09, включая самостоятельное написание кода,
тестов и debugging. Прежние адреса, прогресс и рабочие папки сохраняются.
87 уроков содержательно адаптируют разрешённые материалы Exercism, Learn Go with
Tests, Ardan Labs, SHAD и Birsi Go Training; происхождение и notices сохранены.

Локальный UI и короткие команды `learn` используют общий learner engine.
**Решения пишутся в GoLand**, в рабочей папке выбранного урока. UI показывает
маршрут, проверки, подсказки и отдельный прогресс освоения навыков.
Запуск из checkout: `go run ./cmd/learn ui` (или `learn ui` с настроенным launcher).
Инструкция: [LEARNER_CLI.md](docs/LEARNER_CLI.md). Историческая поставка UI:
[LOCAL_UI_GOAL.md](docs/LOCAL_UI_GOAL.md).

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
go run ./cmd/coverage-report -root . -scope full -checks .tmp/results/task-checks.yaml
go run ./cmd/coverage-report -root . -scope foundation -checks .tmp/results/task-checks.yaml
go run ./cmd/validate -root . -coverage -checks .tmp/results/task-checks.yaml
```

Последняя команда станет зелёной только после наполнения банка в следующих фазах.
`validate -foundation-coverage -checks <collection>` отдельно требует числовые
минимумы 141 foundation skills. Оба numeric отчёта сохраняют все 145 skills,
дефициты specialist и full, а также общий target 350..500 ready tasks.
Генерация и `coverage-report -check` допускают числовые дефициты; enforcement
остаётся в `validate`. Modes, distinct types и observable outcomes требуют
независимой semantic review. Подробный контракт — [QUALITY_GATES](docs/QUALITY_GATES.md).

## Карта репозитория

```text
adapters/                 контракты будущих source-adapters
catalog/candidates/       метаданные сырых кандидатов (не готовые задачи)
cmd/source-sync/          pinned cache synchronization
cmd/harvest/              offline candidate extraction
cmd/catalog-report/       deterministic import/taxonomy reports
cmd/taxonomy-report/      additive v4 migration report and read-only freshness check
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
