# Go Practice Bank

Кураторский банк практических задач по Go. Репозиторий находится в ранней
**ФАЗЕ 2C**. Taxonomy v3 зафиксирована отдельным ADR и полной v2→v3 migration, все concept labels
прикреплённого Exercism Go явно классифицированы, а candidate-to-skill evidence
хранится отдельно от task coverage. В репозитории остаётся ранее созданный
ручной batch из трёх задач; текущий taxonomy/provenance batch не создаёт новых
задач. Массовая генерация и пользовательский интерфейс отсутствуют.

## Быстрые команды

```bash
go test ./...
go run ./cmd/validate -root . -phase1 -run-commands
go run ./cmd/source-sync -root . -source exercism-go
go run ./cmd/harvest -root . -source exercism-go
go run ./cmd/catalog-report -root . -source exercism-go
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
cmd/validate/             CLI-валидатор
curriculum/skills.yaml    taxonomy атомарных навыков и mastery-критерии
curriculum/migrations/    machine-readable migrations и immutable snapshots
curriculum/mappings/      versioned карты внешних concepts → skills
curriculum/evidence/      reviewed candidate-to-skill evidence
docs/                     нормативные документы проекта
docs/adr/                 решения и предложения по эволюции taxonomy
internal/validation/      модель и проверки каталога
provenance/               правила хранения происхождения и attribution
schemas/                  task, concept-map и evidence schemas
sources/registry.yaml     реестр аудированных источников
tasks/                    изолированные ready task packages
```

Детали trust boundary и cache: [`docs/SOURCE_SYNC.md`](docs/SOURCE_SYNC.md).

Формат task package и жизненный цикл описаны в
[`docs/TASK_SPEC.md`](docs/TASK_SPEC.md).

## Начать решать

Каждая задача — отдельный Go module. Например:

```bash
cd tasks/ready/variadic-sum-001
go test ./tests/public
```

Первый запуск public tests ожидаемо падает: исправляйте только файлы в
`starter/`, затем повторяйте команду до зелёного результата. Не открывайте
`tests/grader/` и `solution/`, если хотите решить задачу без подсказок.
