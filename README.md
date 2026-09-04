# Go Practice Bank

Кураторский банк практических задач по Go. Репозиторий находится в
**ФАЗЕ 3D**. Taxonomy v3 и полный pedagogical triage Exercism зафиксированы;
добавлены lifecycle, Task Contract v2, multi-source provenance, изолированные
fixed/mutation/workspace verification profiles и локальный learner CLI.
Массовая генерация и web UI отсутствуют.

Актуальный личный маршрут, критерии готовности и ближайшие партии работ:
[docs/ROADMAP.md](docs/ROADMAP.md).

## Быстрые команды

```bash
go test ./...
go test -race ./...
go vet ./...
go test -tags=taskintegration ./...
go run ./cmd/learn list
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

Learner CLI сам создаёт отдельный workspace и не изменяет canonical package:

```bash
go run ./cmd/learn list
go run ./cmd/learn show map-take-001
go run ./cmd/learn start map-take-001
# измените файлы в напечатанном workspace
go run ./cmd/learn check
go run ./cmd/learn complete
```

Полный список команд, правила локального progress и граница hidden material:
[`docs/LEARNER_CLI.md`](docs/LEARNER_CLI.md).

Первое живое прохождение: `map-take-001`, затем `detached-append-001`.
На перерыв используйте `learn pause`, по возвращении — `learn resume <task-id>`.
[Протокол и feedback](docs/LEARNER_CLI.md#first-human-trial).
