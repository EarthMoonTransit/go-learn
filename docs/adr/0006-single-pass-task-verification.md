# ADR 0006: single-pass verification учебных задач

- Статус: accepted
- Дата: 2026-09-04
- Решение: ФАЗА 3A.1b

## Контекст

До этого решения `BuildAggregateReport` неявно вызывал `CheckAll`. В полном CI
validator, `task-check` и `task-report` создавали три независимых полных
проверки, а validator дополнительно выполнял legacy manifests. Каждый task
check создавал собственный `isolation.Environment`; каждый variant запускал
сначала compile-only, затем test command. Обычные unit tests повторяли ту же
subprocess-heavy механику на synthetic repositories.

## Решение

`CheckSession` является ресурсом одного явного invocation. Session загружает
toolchain policy один раз и владеет одним isolated root с общими `HOME`,
`GOCACHE`, `GOMODCACHE` и `TMPDIR`. `task-check -all` последовательно проверяет
все packages через одну session. Каждый starter, reference, mutant, race run и
legacy command всё равно получает новую workspace; source checkout не является
рабочей директорией процесса.

Один variant использует один `go test -json -count=1` вместо отдельной
compile-only команды. Наличие test-run events отличает успешно собранный suite
от compile failure. Проверка reference по-прежнему выполняется дважды и
сравнивает canonical events; семантика gates ADR 0005 не ослаблена.

Session не хранится в package-global state. Standalone `Check` создаёт и
закрывает собственную session. Ошибка `Close` возвращается отдельно либо
присоединяется к основной ошибке, но не заменяет её. Shared compiler cache —
оптимизация доверенного repository-authored content, не security sandbox.

## Вычисление и представление результата

`CheckAll` возвращает versioned `CheckCollection`.
`BuildAggregateReportFromChecks` принимает уже проверенную collection и не
имеет execution path. Локальный fallback `BuildAggregateReport` сохраняется: он
сам делает один `CheckAll`, после чего вызывает чистое преобразование. CI этот
fallback не использует.

Collection не содержит timestamps. Для каждой задачи записывается digest от
contract version, lifecycle state, относительных путей, permissions и байтов
всех файлов package, relevant schemas, taxonomy, source registry и
`config/toolchain.yaml`. Отдельно
записывается digest toolchain policy. Reader отклоняет неверную schema/version,
изменённые inputs, неизвестные, отсутствующие, дублированные или несортированные
task results и непротиворечивость execution counters.

Collection не является подписью или доказательством факта исполнения. В CI ей
доверяют только как результату обязательного предыдущего шага того же job.
Persistent cross-run cache и changed-task resolver не создаются.

## Tests и CI

Обычные unit tests используют маленький internal runner interface и fake
runner; они проверяют policy, schemas, digests, promotion и report transforms
без дочернего Go toolchain. Реальный contract собран под build tag
`taskintegration`; одна integration session покрывает valid, adversarial,
race, legacy и isolation scenarios с отдельными workspaces.

Race regression использует тот же single-pass runner и отдельный типизированный
mode: normal variants работают с cgo disabled, race variant — с cgo enabled.
Integration test делает skip только после явного определения unsupported
GOOS/GOARCH; linux/amd64 всегда обязан выполнить реальный race command.

CI строит binaries один раз. Затем один `task-check -all` пишет collection во
временный path, а validator и `task-report -check` повторно используют её.
`-checks` несовместим с `-run-commands`, чтобы скрытый второй execution pass был
невозможен.

## Последствия

Полный bank execution сокращён с четырёх проходов до одного. Быстрые test/race
suites больше не запускают nested Go commands. Детерминизм контролируется
структурными counters и byte comparison, а не нестабильным wall-clock
threshold. Три legacy ready packages и task schema остаются byte-identical.
