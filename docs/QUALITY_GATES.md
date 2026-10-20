# Quality gates

[ADR 0034](adr/0034-finite-v1-course.md) adds separate v1 route-readiness and explicit checklist review. All executable ready-task gates below remain mandatory. Manual units/checklists do not create automatic coverage or full-bank release/export eligibility; incomplete-bank preparation still follows ADR 0022.

Task Contract v2 выбирает проверяемую матрицу через конечный
`verification.profile`; произвольные command strings запрещены. Для
`mutation-tests` mutant coverage означает способность learner-authored tests
отличить reference от каждой обязательной ошибки. Для `workspace-tests` перед
запуском отдельно проверяются allowlisted skeleton и точное множество editable
paths. Полные контракты описаны в `docs/TASK_CONTRACT_V2.md`.

Для `fuzz-seed-tests` действует ADR 0032 и [полный finite source contract](FUZZ_SOURCE.md).
Public повторяет canonical и все declared correct variants. Full сначала
подтверждает public correctness, затем требует eligible assertion detection
каждого обязательного defect model во всех повторах. Starter оставляет один и
тот же eligible survivor во всех повторах. Seeds принадлежат learner; скрытый
corpus не добавляется. Panic/race/timeout/truncation/quota, даже после assertion,
не являются detection. Сравниваются native semantic events и полные observer
values/status; nil/empty нормализуются по значению. Для этого профиля обычные
правила starter-failure ниже заменены starter-correctness и insufficiency.

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
- task package остаётся pure Go: запрещены `import "C"`, cgo directives,
  native source/object files, native compiler/linker flags, external modules и
  `toolchain` directive;

Варианты собираются checker-owned способом: implementations не содержат tests,
а canonical public/grader tests нельзя заменить solution binding. Все команды
выполняются в disposable workspaces с минимальным allowlisted environment,
offline module resolution и отдельными home/cache/temp paths.

## Gate 5 — review и maintainability

### Public feedback при создании и исправлении задач

Для каждого базового наблюдаемого требования условия нужен public assertion
и реалистичный ошибочный вариант, нарушающий это требование. В `review/notes.md`
автор записывает соответствие: требование → public test → ошибочный вариант →
ожидаемая диагностика. Reviewer проверяет это соответствие при создании или
существенном изменении задачи. Формат task/review schemas не меняется; это
проверяемое evidence в notes, а не новый автоматически вычисляемый gate.

Public-проверка ошибочного варианта должна завершиться test failure и назвать
ошибку понятным образом (для значений — actual/expected, при необходимости тип).
Compile error, timeout и общий текст full verification не подтверждают это
покрытие. Правильное решение должно проходить public и full проверки.
Скрытые проверки могут расширять набор входов и граничных случаев для уже
заявленного поведения; успешный public run не гарантирует полный зачёт.
Для test-submission задач учитывается их mutation-tests профиль: не требуется
публиковать скрытые mutants или превращать public suite в полный grader.

Исправление обнаруженного public-success/full-failure пробела сопровождается
сохранённым автоматическим тестом: ошибочный вариант отклоняется именно public
проверкой с полезной диагностикой. Для `literal-bindings-001` это
`TestRealOpeningLessonRunsThroughLearnerWorkspace` в `internal/learner`:

| Требование | Public test | Ошибочный вариант | Диагностика |
| --- | --- | --- | --- |
| count: int 12 | TestNumericBindings | floating-count | actual float64; want int 12 |
| scale: float64 1.25 | TestNumericBindings | integer-scale | actual int; want float64 1.25 |
| label: string G | TestTextAndBooleanBindings | rune-label | actual int32; want string G |
| enabled: bool true | TestTextAndBooleanBindings | disabled | actual false; want bool true |

Удаление public assertions для label/enabled должно ломать эту регрессию, даже
если full grader всё ещё обнаруживает все четыре ошибки. Исполнение — только
через disposable verifier. Этот тест защищает конкретный урок; полноту
сопоставления требований остальных задач по-прежнему проверяет reviewer.

### Остальные требования

- `review/review.yaml` содержит reviewer, date, decision и minimum mutants;
- story, alignment, API, provenance, license, clean-room и leakage evidence
  подтверждены явно;
- `go fmt`, `go vet` и repository validation зелёные;
- public export исключает grader, reference, mutants и internal review;
- commands и reports детерминированы.

## Machine-readable status

Одиночный `task-check -id` выдаёт документ
`schemas/task-quality-report.schema.json`; `task-check -all` выдаёт versioned
`schemas/task-check-collection.schema.json`. Gate использует `pass`, `fail`
или `not-applicable`; race gate дополнительно использует
`infrastructure-unsupported` для явно неподдерживаемой GOOS/GOARCH. Такой
статус не является pass и делает technical status failing. Mutant summary
фиксирует `required`, `total`, `detected` и `survived`. Дополнительно фиксируются
technical/review status, promotion eligibility, failing gate IDs, starter,
reference, race и legacy compatibility.

Collection фиксирует task input digests и структурные execution counters без
timestamps. `task-report` детерминированно создаёт `reports/task-quality.yaml`
и Markdown из validated collection.
После появления десяти lifecycle packages aggregate gate требует минимум три
task types и не более 60% одного типа. Curriculum policy остаётся строже: не
более 50% primary-встреч одного типа на домен.

## Repository acceptance

Authoring infrastructure и промежуточные preparation content batches неполного
банка (точная граница с release зафиксирована ADR 0022):

```bash
go test ./...
go test -race ./...
go vet ./...
go test -tags=taskintegration ./...
go run ./cmd/task-check -root . -all -output .tmp/results/task-checks.yaml
go run ./cmd/validate -root . -phase1 -checks .tmp/results/task-checks.yaml
go run ./cmd/task-report -root . -check -checks .tmp/results/task-checks.yaml
go run ./cmd/coverage-report -root . -scope full -check -checks .tmp/results/task-checks.yaml
go run ./cmd/coverage-report -root . -scope foundation -check -checks .tmp/results/task-checks.yaml
```

### Scoped numeric reports

`coverage-report` требует явные `-scope full|foundation` и `-checks`; команда
не запускает task code. Она сначала проверяет ordinary repository validation,
schema/freshness полного collection и `Report.Passed` каждой задачи. Любая ошибка,
включая ordinary error с частичным Result, останавливает запись обоих файлов.
Fresh collection с failed package и согласованными task-quality artifacts также
отклоняется. Без eligible evidence новый numeric Result/CLI status unavailable;
старые raw counts доступны для диагностики. Новый foundation selector требует
`-checks` или явный `-run-commands`, создающий один effective collection.

Фиксированные файлы `reports/coverage-{full,foundation}.{md,yaml}` используют
собственную strict `schemas/coverage-report.schema.json`. Оба scope сохраняют
145 active rows с actual/required primary/total и дефицитами, выбранные sorted
IDs (145 или 141), foundation/specialist/full sums и independent ready target.
Каждый дефицит обрезается снизу нулём **до** суммирования: surplus другого skill
не компенсирует пропуск, total не компенсирует недостаток primary.
Full numeric pass требует всех minima **и** общего target 350..500; foundation
numeric pass требует только foundation minima. Меньший foundation target
не вводится. Генерация без `-check` обновляет файлы; `-check` сравнивает точные
bytes. Оба режима успешны при валидном неполном банке и отражают deficits.
`validate -coverage` и `validate -foundation-coverage` отдельно enforce числа;
оба флага вместе и каждый из них с `-phase1` — CLI usage error (exit 2).
Ordinary/full v3 validation сохраняется; новые foundation/report paths требуют v4.

Input identity — SHA-256 JSON encoding typed struct в порядке полей:
`report_contract`, `selected_scope`, `taxonomy_version`, `classification`,
`report_schema_sha256`, `check_contract`, `toolchain_policy_digest`, `tasks`.
Classification содержит sorted по skill ID пары `skill_id, scope`; tasks —
sorted по task ID тройки `id, state, input_digest` из полного validated collection.
Schema digest вычисляется по точным bytes. Timing, execution counters, timestamps,
absolute paths, formatting YAML collection и outputs в identity не входят.
Исходные task/contract/schema bytes связаны через task input digests и schema
digest; их форматирование остаётся input change. При этом
переставленный collection по-прежнему отклоняется обычным validator: сортировка
identity не ослабляет contract collection. Собственная report schema входит
только в v4 TaskInputDigest; v3 contract и TaskContentDigest не меняются.

Порядок обновления после изменения inputs: один `task-check -all -output`, затем
`task-report -checks` для task-quality, затем оба `coverage-report -checks`.
Validator, task-report check и оба coverage checks используют тот же collection.
Собственные coverage outputs не проверяются внутри ordinary ValidateRoot,
чтобы generation не зависела от собственной freshness. CI проверяет их отдельно.
`semantic_review_status: not_evaluated_by_numeric_coverage` означает, что counts
не устанавливают required modes, distinct task types, observable outcomes или
learner mastery. Release требует отдельной semantic acceptance.

Metadata-регрессии numeric guard/report используют ограниченный disposable
catalog: все задачи шести Theory pages, три точные migration references и
diversity control (сейчас 12 пакетов с исходными bytes, максимум 16). Полные
taxonomy, scopes, route, source metadata и ordinary validation сохраняются; synthetic
checks пересчитывают digests только этого fixture. Assertions для failed/current
evidence, partial Result, v3, обоих scopes и всех active rows сохраняются.
Это не evidence полного банка: acceptance collection исполняет все packages
через `task-check` и повторно используется всеми report/validator consumers.

Public release/export и итоговый foundation acceptance дополнительно используют
`-coverage`. Промежуточный ready package не объявляет весь банк выпущенным;
P2/P4C/P5 первой волны являются preparation batches. Ни `-phase1`, ни наличие
upstream tests не ослабляют gates уже существующих packages.

The tooling matrix `TestReviewedDefectsHavePublicFeedback` additionally locks
public-only failure and useful diagnostics for `boolean-policies-001`
(`exclusive-or`), `token-values-001` (`trim-space`), and `column-occupancy-001`
(`case-normalized`). These cases execute through `Service.Check` in temporary
learner homes, with reference success and unchanged canonical content checked.

The expanded public-feedback integration matrix checks reviewed ownership and
interface defects independently of hidden tests. Tolerance fixtures are valid
alternatives, not mutants: public and full learner verification must accept them.
When testing the regression itself, remove only the relevant public assertion in
a disposable repository copy, refresh that copy's digest, and require the outer
regression to fail because the public verifier unexpectedly passed. A compile
failure does not establish this property.

The public-feedback regression matrix must cover package-relative workspace
patches as well as single-file tasks. Verify both setter directions, exact text
preservation and the distinction between an error's Error text and custom
formatting when those behaviors are promised by the task.

The public-feedback matrix also protects bounded-copy sentinel identity and successful-stream ownership, flushing a caller-owned pending buffer with empty input, and discarding a partially decoded JSON object. Each row runs only learner `Check`, verifies a useful diagnostic and verifies the reference in the same disposable learner workspace.

Testing-component public regressions cover Helper-before-acquire, preservation of the exact close-error argument to Errorf, and the initial empty writer Snapshot. The Snapshot fixture is a narrow tooling-owned modification of the reference installed only in a disposable learner workspace, preserving the accepted content package.

Concurrency public regressions verify that an unused finite stream cannot retain a blocked worker, a readable nonclosed cancellation channel is handled, and both unsynchronized counter updates and reads receive public DATA RACE diagnostics. These checks execute only disposable learner payloads, with a passing reference counterpart for every row.
