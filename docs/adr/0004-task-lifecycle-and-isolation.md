# ADR 0004: lifecycle и изолированная проверка task packages

> Execution binding, environment и promotion semantics уточнены и частично
> заменены ADR 0005. Lifecycle state machine из этого ADR остаётся действующей.

- Статус: accepted
- Дата: 2026-09-04
- Решение: ФАЗА 3A

## Контекст

Schema v1 использует `quality_status` как смесь редакционного состояния и
результата проверки. Её нельзя менять без отдельной миграции, а три уже готовые
задачи должны остаться byte-identical. Authoring-инструментам при этом нужны
явные состояния, проверяемые переходы и безопасный запуск недоверенного кода.

## Решение

Lifecycle определяется расположением task package и `lifecycle.yaml`:

| Состояние | Каталог | `quality_status` v1 |
| --- | --- | --- |
| `draft` | `tasks/drafts/` | `draft` |
| `in-review` | `tasks/review/` | `review` |
| `ready` | `tasks/ready/` | `ready` |
| `rejected` | `tasks/rejected/` | `quarantined` |
| `retired` | `tasks/retired/` | `retired` |

Разрешены переходы `draft → in-review`, `in-review → draft`, `in-review →
ready`, `in-review → rejected`, `rejected → draft` и `ready → retired`.
`validated` остаётся значением schema v1 для совместимости, но не является
lifecycle-состоянием. Результат gates хранится отдельно в machine-readable
quality report. Каталог и `lifecycle.yaml` обязаны совпадать. Переход в `ready`
разрешён только из `in-review`, после полного `task-check` и одобряющего review.
Целевой каталог никогда не перезаписывается. Перемещение выполняется rename в
пределах одного filesystem; metadata подготавливаются atomic writes и
откатываются при ошибке rename.

Три существующих ready packages считаются legacy v1: валидатор продолжает их
поддерживать без добавления новых файлов. Все packages, созданные `task-init`,
следуют полному контракту ФАЗЫ 3A.

## Полный package contract

```text
<state>/<task-id>/
├── task.yaml
├── lifecycle.yaml
├── TASK.md
├── go.mod
├── starter/*.go
├── tests/public/*_test.go
├── tests/grader/*_test.go
├── solution/reference/*.go
├── solution/mutants/<mutant-id>/*.go
├── review/review.yaml
├── review/notes.md
├── hints/HINTS.md
├── provenance.yaml
└── ATTRIBUTION.md              # обязателен для адаптаций
```

Первоначальный контракт bindings заменён checker-owned assembly из ADR 0005.
Ни один test не импортирует каталоги из source checkout. Не-Go fixtures и
внешние зависимости в контракте 3A не разрешены; их поддержка потребует ADR.

## Isolation

Ни одна test-команда не запускается из source checkout. Инструмент отклоняет
symlinks и нерегулярные файлы, копирует только разрешённые артефакты в новый
temporary module и запускает Go там. Для legacy `validation_commands` целый
task package копируется в temp перед каждой командой; shell больше не используется.

Каждый запуск получает `GOWORK=off`, `GOPROXY=off`, `GOTOOLCHAIN=local`,
`GOFLAGS=-mod=readonly` и отдельные `GOCACHE`, `GOMODCACHE`, `TMPDIR`.
`go.mod` новой задачи обязан иметь module `task.local/<task-id>` и не может
содержать `require`, `replace`, `exclude`, `retract`, `tool` или `godebug`.
`go.work` запрещён. Это исключает workspace capture и сетевое разрешение
зависимостей; OS-level sandbox не заявляется и остаётся отдельной задачей.

## Проверочная матрица

- starter + public: компилируется, но test run обязан падать;
- starter + public + grader: компилируется, но test run обязан падать;
- reference + public + grader: проходит два одинаковых запуска;
- каждый mutant + public + grader: компилируется и обнаруживается тестами;
- concurrency tasks дополнительно проходят `go test -race` на reference.

## Последствия

`schemas/task.schema.json` и существующие ready packages не меняются. Новый
lifecycle и review получают отдельные versioned schemas. Public export обязан
исключать `tests/grader`, `solution/reference`, `solution/mutants` и внутренние
review artifacts.
