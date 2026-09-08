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
