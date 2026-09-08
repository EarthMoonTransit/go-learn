# Quality gates

Task Contract v2 выбирает проверяемую матрицу через конечный
`verification.profile`; произвольные command strings запрещены. Для
`mutation-tests` mutant coverage означает способность learner-authored tests
отличить reference от каждой обязательной ошибки. Для `workspace-tests` перед
запуском отдельно проверяются allowlisted skeleton и точное множество editable
paths. Полные контракты описаны в `docs/TASK_CONTRACT_V2.md`.

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
