# Отладка учебного цикла

Цель — быстро подтвердить и исправить конкретный сбой, затем один раз выполнить
обязательную приёмку окончательного diff. Правила `AGENTS.md` и ADR 0006
сохраняются. Не запускать authored task code из source checkout и не использовать
реальный learner home для экспериментов.

## Короткий цикл

1. Зафиксировать task ID, действие (GoLand tests / Check / Complete), входной код
   и наблюдаемое сообщение. Сравнить требования условия с public/grader tests.
2. Воспроизвести сбой на одной задаче через learner service с временным home или
   `task-check -id`. Если сбой касается public feedback, полного `task-check`
   недостаточно: нужен вызов public verifier с конкретным ошибочным вариантом.
3. Сразу сохранить воспроизведение в подходящем regression test. Запускать только
   этот тест при итерациях; новый временный acceptance-скрипт нужен лишь когда
   существующий тестовый harness не покрывает сценарий.
4. Исправить минимальный набор файлов. Проверить правильный вариант и понятную
   диагностику ошибочного. Для изменений task content отдельно проверить
   совместимость сохранённых workspaces; не предлагать reset по умолчанию.
5. После стабилизации diff выполнить acceptance из `docs/QUALITY_GATES.md`:
   fast → race → vet → integration; один полный `task-check -all -output`.
   Validator и report-check используют этот же collection через `-checks`.
   При изменении content сначала обновить отчёт из единственного collection,
   затем выполнить остальные gates; не запускать банк повторно для отчёта.
6. Опубликовать PR. Локальную приёмку и CI одного неизменного commit можно
   выполнять одновременно после успешной целевой регрессии. Сливать только
   когда обе завершились успешно и `validate` зелёный на точном head commit.
   После новых изменений проверить затронутое поведение и актуальность evidence.
   Не повторять уже успешные проверки без изменения inputs или новой причины.

Для текущей регрессии из корня репозитория:

```bash
go test -tags=taskintegration ./internal/learner -run '^TestRealOpeningLessonRunsThroughLearnerWorkspace$' -count=1 -v
```

Команда запускает repository test harness; task code выполняется в disposable
copies. Harness проверяет starter, четыре независимые ошибки, public/full успех
reference, учёт попыток и переход к следующей задаче. Для обычной проблемы одной
задачи не нужны повторный browser smoke и полный банк на каждой итерации. Они
нужны, когда меняется соответствующее поведение или это требует acceptance.

## Отладка finite fuzz

```bash
go test -tags=taskintegration ./internal/tasking -run '^TestNativeFuzzTaskCheckAndPublicFeedback$' -count=1 -v
```

Synthetic регрессия выполняет task code только через disposable Engine.
Сначала различать `native-eligibility` (source frame, process или complete
observation) и public assertion failure. Public feedback проверяется на всех
correct variants; paired roundtrip может пройти Check и оставить defect на
Complete. После изменения production verifier source, actual runtime или
grammar contract старый collection не переименовывается: после freeze inputs
создаётся один новый v4 файл для всех acceptance consumers. Реальный learner
home не используется.

Metadata-only `go env` тоже проходит через isolated environment. Перед первой
командой environment создаёт telemetry mode file `off` в своём user config.
Одна переменная `GOTELEMETRY=off` не задаёт режим: без файла Go использует
`local` и может оставить sidecar, который мешает немедленному cleanup.
`TestIsolatedGoTelemetryIsOffBeforeFirstCommand` проверяет actual mode, отсутствие
counter/sidecar files и удаление disposable root; real user config не меняется.

## Где уходит время

В CI run `34168388464` исправления literal-bindings интеграционные тесты заняли
7 минут 13 секунд, полный банк — 1 минуту 13 секунд, весь job — около 10 минут
31 секунды. Это измерение одного запуска, не обещание времени следующих.
Источник: https://github.com/EarthMoonTransit/go-learn/actions/runs/34168388464.
Фикс потребовал отдельных tooling/content PR из-за совместимости task digest.

Сначала измерять конкретный медленный шаг. Параллельная приёмка локально и в CI
сокращает ожидание; targeted regression сокращает цикл исправления. Оптимизацию
runner/cache и разделение CI jobs делать отдельным tooling batch с замерами и
проверкой сохранения изоляции. Не ускорять работу отключением race, integration,
grader или process gate. Пока такой batch не принят, полный CI остаётся нужен.

## Измеренная оптимизация tooling (2026-10-12)

Профиль fast learner suite показал повторный разбор taxonomy и provenance
policy внутри каждого полного чтения каталога. Теперь один вызов каталога
использует один loader общих metadata. Следующий вызов создаёт новый loader;
правки taxonomy, registry, identity, license, schema и task files не скрываются.
Все package checks и digest checks выполняются для каждой задачи.

Интеграционные матрицы используют явный `CheckSession` на top-level test:
холодная сборка стандартной библиотеки не повторяется для каждого Check/Complete.
Каждый variant всё ещё запускается в новой disposable workspace с `-count=1`.
CLI integration сохраняет одноразовый путь; public/grader, mutation, determinism,
vet и race assertions не удалены. Сессия закрывается в `t.Cleanup`, включая
ошибочные сценарии. Подробнее — дополнение к ADR 0006.

Локальные измерения на Apple M4 Max, darwin/arm64, Go 1.27.0, база `7a1c153`
и те же 151 learner tasks (это наблюдения, не wall-clock acceptance thresholds):

| Сценарий | До | После |
| --- | ---: | ---: |
| Полный learner catalog, median из 3 запусков | 1,301 s | 0,639 s |
| Тот же catalog под race, median из 3 запусков | 14,257 s | 3,152 s |
| Allocated bytes на catalog без race | 1,077 GB | 0,540 GB |
| Opening-lesson integration, один запуск | 22,12 s | 7,34 s |

Для воспроизведения benchmark на обеих ревизиях используйте один и тот же
`BenchmarkReadyLearnerCatalog` из `internal/tasking/learner_catalog_test.go`:

```bash
go test ./internal/tasking -run '^$' -bench '^BenchmarkReadyLearnerCatalog$' -benchtime=1x -count=3 -benchmem
go test -race ./internal/tasking -run '^$' -bench '^BenchmarkReadyLearnerCatalog$' -benchtime=1x -count=3 -benchmem
```

Benchmark читает исходники как данные. Проверка
`TestLearnerSessionKeepsFreshWorkspacesAndInputs` дополнительно фиксирует пять
реальных subprocess runs в пяти разных workspace для starter public,
reference full (два determinism runs и vet) и последующего public. Она помещает
ломающий test в старые disposable copies, проверяет отсутствие grader в новом
public workspace, повторную валидацию изменённых inputs и удаление session root.
Timeouts, workflow gates и содержимое банка этой оптимизацией не меняются.

## Компиляция schemas каталога (2026-10-18)

Metadata каждого package раньше заново компилировали одни и те же JSON schemas.
Теперь `learnerPackageLoader` сохраняет compiled schema с первого использования
до конца одной операции чтения каталога. Для каждого schema name используется
отдельный Compiler с прежними settings, resource names и разрешением refs.
Следующая операция создаёт новый loader и видит изменения schemas, включая
ошибку и последующее исправление через тот же Service. Атомарный filesystem
snapshot этим не обещается.

Каждый документ по-прежнему проходит strict YAML decoding, JSON normalization
и schema validation. Результаты package validation не кэшируются. Lookup,
lifecycle inventory, API, provenance, review, symlink и content-digest checks
сохраняются для каждой задачи; standalone validation остаётся свежей.
Fixture с двумя настоящими packages проверяет четыре компиляции и восемь
document validations, затем отклонение неверного следующего package.

Измерения того же `BenchmarkReadyLearnerCatalog` на Apple M4 Max,
darwin/arm64, Go 1.27.0: одинаковые 169 learner tasks, база `1e362157`,
median из трёх запусков командами выше. MB здесь — 1 000 000 bytes.

| Сценарий | До | После |
| --- | ---: | ---: |
| Полный каталог без race | 0,686 s | 0,329 s |
| Полный каталог под race | 3,305 s | 1,673 s |
| Allocated bytes без race | 650,6 MB | 141,0 MB |
| Allocated bytes под race | 704,1 MB | 193,3 MB |

Все 169 экспортируемых `LearnerTaskInfo` до/после совпадают при чтении одних
и тех же inputs. Это локальные наблюдения, не timing acceptance thresholds.
Race, integration, grader, disposable workspaces и обязательные workflow gates
не сокращаются.

## Разбор taxonomy v4 (2026-10-18)

Приёмка миграции v4 обнаружила повторный YAML-разбор live taxonomy и frozen v3
при одиночных metadata reads. Private read-only consumers теперь сохраняют
последний успешный parse в двух slots по роли и точным bytes. Перед lookup каждый
вызов заново читает файл через прежние path/size guards. Migration, scopes,
reviewed bytes и все semantic checks проверяются каждый раз; результаты
validation и repository roots не сохраняются. Обычный `readStrict` продолжает
создавать независимые mutable maps и nodes.

На Apple M4 Max, darwin/arm64, Go 1.27.0, одна серия из 20 вызовов
`LoadReadyLearnerTask(".", "zero-results-001")` на одинаковых v4 inputs дала
0,388 → 0,239 s без race и 5,173 → 2,698 s под race. Allocated bytes под race
снизились с 327,3 до 213,8 MB. База — `b89f80c`; это наблюдение одного запуска,
не timing threshold. Для полного каталога выигрыш небольшой: его filesystem
и Go AST guards сохраняются. Все 169 экспортируемых `LearnerTaskInfo` совпали.

Регрессии прогревают slots перед изменением bytes, YAML errors, missing files,
symlinks, nonregular files и превышением 4 MiB. Отдельно проверяются свежие
semantic checks, конкурентная замена slots под race и невозможность испортить
shared parse через mutable публичные outputs. Таймауты и workflow не меняются.

## CI integration suite budget

Runs 34253711927 and 34254229744 both hit Go's default 10-minute package
timeout during the growing learner integration suite. They were executing
disposable verifier subprocesses, with no preceding assertion failure. Local
acceptance completed. CI now gives the integration suite an explicit bounded
20-minute package budget and caps the whole job at 35 minutes. Per-task command
timeouts, isolation, required test cases and collection reuse are unchanged.
This prevents an aggregate timeout from truncating otherwise bounded checks;
it is not a runtime-speed improvement. Cache/session optimization remains
a separate measured change.

## CI race suite budget (2026-10-18)

Run `39136659242` stopped the learner race package at Go's default 10-minute
alarm (600.036 s). The current test was 12 seconds old and its active subtest
2 seconds old, running metadata YAML loading; there was no preceding assertion
failure or data-race diagnostic. The same source passed local race in 370.528 s.
CI now gives the complete race suite a bounded 20-minute package budget.
The job remains bounded at 50 minutes and tagged integration at 30 minutes.
Per-task deadlines, assertions, race instrumentation and collection reuse are
unchanged. This aggregate budget does not establish that a run passes: exact-head
`validate` must still complete successfully. The workflow/documentation-only
change does not invalidate the current checker-input collection.

Source: https://github.com/EarthMoonTransit/go-learn/actions/runs/39136659242.

## T0a read-only fixture budget

T0a state regressions use a bounded disposable eight-task catalog with the real
task manifests and old addresses. Omitted introductions are explicit fixture-only
preparation assumptions, never learner evidence. The full 151-row mapping and
its two formal blockers are tested separately against the repository. Repeated
whole-bank metadata loads in each history/UI fixture exceeded the race package's
10-minute CI budget; bounding those fixtures preserves their Service/CLI paths
and every storage/history/workspace assertion without increasing the timeout.

## Cold race-fixture compilation

Runs 34284014940 and 34284226173 detected the race-only mutant, but the new
integration fixture's first public check exceeded its inherited 30-second command
budget while compiling with `-race`. The later reference and race gates passed.
That synthetic fixture now allows 60 seconds per command, including its ordinary
behavior control. Real task limits and every defect/public-feedback assertion are
unchanged. A timeout remains an infrastructure failure, never proof of detection.
Re-run the focused fixture before the mandatory suite; reuse the existing full
collection when only this test harness changes and bank inputs remain identical.

## Reviewed public-feedback matrix

`TestReviewedDefectsHavePublicFeedback` exercises inclusive boolean OR, exact
whitespace matching and exact map-key lookup using temporary learner homes.
Each reviewed defect must fail `Check` with its input/actual/expected diagnostic;
the reference must then pass the same public-only path. Hidden grader success
cannot substitute for this regression. Canonical task digests remain unchanged.

Run the bounded loop with:

```bash
go test -tags=taskintegration ./internal/learner -run '^TestReviewedDefectsHavePublicFeedback$' -count=1 -v
```

Add a matrix row when fixing a public-success/full-failure gap that would
otherwise remain masked by hidden mutant coverage. Keep content edits and this
tooling regression in separate reviewable batches; complete both before claiming
the feedback regression is protected.

The ownership/interface feedback matrix also covers cleared dictionaries,
overflow and overlapping interface cases, category filtering, sentinel identity,
zero-reserve aliases, capped-view input preservation and writer closure on errors
or short writes. Each repaired public requirement has a concrete defect and a
useful diagnostic checked through `Service.Check` in a temporary learner home.
Saved tolerated-output review fixtures must pass both public Check and full
Complete in that temporary home, protecting allowed numerical alternatives.
No fixture completion is counted as real learner progress.

C10 regressions cover two-package setters, exact padded labels and Error text
when a cause implements fmt.Formatter. The test installer copies either a
fixed-task file into subject or all package-relative files from a workspace patch.
Each defect must fail public Check, then the canonical reference must pass. An
omission experiment removing the three public assertions makes all three
regressions fail because Check incorrectly passes.

As the bank approaches 150, CI allows 30 minutes for the integration package
tests and 50 minutes for the complete job. The preceding 104-package tooling
run 34275377981 used 18m19s for integration alone, leaving little room within
the previous 20-minute bound. Per-task execution limits are unchanged; no
absolute learner performance criterion or weakened acceptance gate is added.

The public-feedback matrix also protects bounded-copy sentinel identity and successful-stream ownership, flushing a caller-owned pending buffer with empty input, and discarding a partially decoded JSON object. Each row runs only learner `Check`, verifies a useful diagnostic and verifies the reference in the same disposable learner workspace.

Testing-component public regressions cover Helper-before-acquire, preservation of the exact close-error argument to Errorf, and the initial empty writer Snapshot. The Snapshot fixture is a narrow tooling-owned modification of the reference installed only in a disposable learner workspace, preserving the accepted content package.

Concurrency public regressions verify that an unused finite stream cannot retain a blocked worker, a readable nonclosed cancellation channel is handled, and both unsynchronized counter updates and reads receive public DATA RACE diagnostics. These checks execute only disposable learner payloads, with a passing reference counterpart for every row.

The C2 feedback regression uses named task/variant subtests for three Formatter
requirements and all eleven Resource defects. It checks useful public diagnostics
and public reference success independently of hidden grader coverage. Target only
these cases while iterating:

```bash
go test -tags=taskintegration ./internal/learner -run 'TestReviewedDefectsHavePublicFeedback/(formatter-contract-001|resource-finalize-001)' -count=1 -v
```

Worker's P13 revision is pinned separately by ADR 0028. The bounded fixture matrix
runs the proposed starter with each of the eight existing controlled defect
bindings in a temporary repository. Inline work is rejected by a named assertion
after safe release; premature/shared completion remain intentionally uncovered.
It supplements historical workspace preservation and real learner verification;
command timeout is never accepted as evidence of defect detection.

### C3 public feedback and valid alternatives

The named `TestReviewedDefectsHavePublicFeedback/<id>/<mutant>` rows execute
public checks only for all 41 C3 defects, require useful assertion diagnostics,
and check a passing reference counterpart. Hidden grader detection, command
timeouts and compilation failures cannot replace these public diagnostics. The
existing Worker matrix preserves its two intentionally uncovered starter variants.

The tolerance matrix installs review fixtures as learner data and runs public
and full checks for both MergeTwo helper styles, buffered Forward output and an
equivalent WithDeadline wrapper. Each top-level matrix owns an explicit disposable
CheckSession; each verification uses a fresh temporary module and each case a
private learner home. Canonical task digests must stay unchanged.

Seven disposable assertion-removal experiments were run for held launch,
right-close merge progress, forwarding before EOF, retained input receivers,
nil-success result/release and earlier-parent child ownership. Every experiment
failed the outer row on `Passed:true`. The cleanup-join row instead requires its
named join diagnostic alongside other result assertions; no join-only
false-accept claim is made. Keep these proof boundaries when modifying fixtures.
