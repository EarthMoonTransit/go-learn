# GPT-6 Astra evaluation protocol

Этот документ содержит два практических теста модели на Go Practice Bank:

1. независимый аудит репозитория и готовности human calibration;
2. педагогический review реального learner submission.

Первый тест проверяет длинное end-to-end reasoning по коду, продукту, учебной методике и CI. Второй проверяет, может ли модель помогать учиться, не превращаясь в генератор готового решения.

## Test A — независимый аудит репозитория

Перед запуском дать модели доступ к репозиторию `EarthMoonTransit/go-learn`. Использовать максимальный доступный reasoning effort. Первый проход строго read-only.

Скопировать следующий prompt целиком.

```text
Ты проводишь независимый senior-level аудит проекта Go Practice Bank.

Repository:
https://github.com/EarthMoonTransit/go-learn

Проверяемая ветка:
master

Baseline на момент постановки задачи:
2f09e573c315efb2cc6c7ca1b876db4127eb952f

Язык отчёта: русский.

Это первый read-only pass. Не изменяй файлы, не создавай branch/PR,
не коммить и не предлагай массовую генерацию задач до завершения аудита.

## Контекст продукта

Прочитай `docs/PROJECT_VISION.md` как заявленное продуктовое направление,
но не принимай его утверждения на веру. Сверь их с фактическим кодом,
Task Contract, lifecycle, reports, ready/draft tasks и learner CLI.

Проект должен стать локальной practice-first системой глубокого изучения Go:
короткие взрослые задания, разные формы практики, hidden verification,
meaningful mutants, provenance, адаптивный маршрут и AI-assisted review
без преждевременной выдачи решения.

## Обязательные источники внутри репозитория

Изучи как минимум:

- README.md
- AGENTS.md
- CONTRIBUTING.md
- docs/PROJECT_VISION.md
- docs/LEARNER_CLI.md
- docs/TASK_CONTRACT_V2.md
- docs/PROVENANCE_V2.md
- docs/QUALITY_GATES.md
- relevant ADR files
- curriculum/skills.yaml
- reports/taxonomy-v3-coverage-plan.md
- reports/pilot-v2-red-team.md
- reports/pilot-v2-promotion.md
- reports/learner-cli-acceptance.md
- reports/task-quality.md
- cmd/learn/
- internal/learner/
- internal/tasking/learner.go
- internal/tasking Contract v2 execution code
- обе learner-visible ready v2 tasks
- несколько returned-to-draft tasks
- relevant integration tests
- GitHub Actions workflow и последний CI status.

Reports являются evidence, но не source of truth. Проверяй выводы по коду.

## Read-only commands

Используй temporary directories. Не изменяй canonical task packages
и не сохраняй learner state в обычный user config.

Запусти или эквивалентно проверь:

- go test ./...
- go test -race ./...
- go vet ./...
- go test -tags=taskintegration ./...
- GO_LEARN_HOME=<temp> go run ./cmd/learn list
- GO_LEARN_HOME=<temp> go run ./cmd/learn next
- GO_LEARN_HOME=<temp> go run ./cmd/learn show map-take-001
- GO_LEARN_HOME=<temp> go run ./cmd/learn show detached-append-001
- start/check/complete flow в temporary learner home;
- отсутствие grader/reference/mutants в persistent learner workspace;
- поведение CLI после restart/resume;
- safe review-context export.

Можно создать временную reference-equivalent submission только в learner
workspace, чтобы проверить end-to-end flow. Не редактируй `tasks/ready`.

## Главные вопросы

1. В каком фактическом состоянии находится проект?
2. Что сейчас действительно готово, а что только заявлено или покрыто synthetic fixtures?
3. Может ли новый пользователь начать обучение с нуля через текущий CLI?
4. Может ли текущий learner пройти хотя бы честный calibration cohort?
5. Насколько полезны public diagnostics, hints, complete и review flow?
6. Не утекают ли hidden grader, reference solution, mutants или private expectations?
7. Корректно ли работают progress, session duration, interruptions и concurrent commands?
8. Есть ли crash-recovery проблема у progress lock?
9. Насколько честно scheduler обрабатывает отсутствующие prerequisites?
10. Соответствуют ли две real ready v2 tasks заявленной педагогической планке?
11. Насколько реальны test-writing и workspace-patch profiles вне synthetic tests?
12. Какие returned-to-draft tasks выгоднее всего исправить для cohort из 8–12 задач?
13. Есть ли противоречия между AGENTS, ADR, schemas и текущей provenance model?
14. Какие части архитектуры переусложнены, а каких гарантий не хватает?
15. Что должно быть следующим действием learner, а что — следующим PR разработчика?

## Гипотезы для независимой проверки

Не считай их истинными заранее. Подтверди или опровергни по коду и reproduction:

- learner CLI показывает только две реальные v2-задачи, несмотря на пять ready packages;
- обе доступные задачи имеют prerequisites, которые текущий v2 route не позволяет закрыть;
- `next` выбирает наименее плохую задачу даже при отсутствующих prerequisites;
- active duration может искажаться при длительном перерыве или повторном resume;
- stale `.progress.lock` может заблокировать CLI после аварийного завершения;
- `attempts` смешивает public checks и full completion attempts;
- `review` не выполняет code review, а глубокий review вынесен в export context;
- реальные ready tasks пока не проверяют human UX для tests/workspace-patch submissions;
- AGENTS.md содержит устаревшее правило про internal original source;
- текущие отчёты проходят CI, но curriculum alpha ещё не достигнута.

## Требования к findings

Каждый finding должен содержать:

- severity: P0 / P1 / P2 / P3;
- category;
- точное утверждение;
- evidence с `path:line`;
- reproduction или reasoning;
- learner impact;
- минимальное исправление;
- что не следует делать, чтобы не переусложнить решение.

Не называй stylistic preference дефектом. Не предлагай новую abstraction,
если проблема решается меньшим изменением.

## Обязательный формат результата

### 1. Executive verdict

Однозначно укажи:

- technical alpha status;
- curriculum alpha status;
- human-calibration readiness;
- можно ли learner начинать прямо сейчас и с какой целью.

### 2. Verified project snapshot

Дай точные числа:

- domains и assignable skills;
- audited sources и reviewed candidates;
- lifecycle packages по states;
- learner-visible real tasks;
- submission profiles, проверенные real tasks и synthetic fixtures;
- CI status и команды.

### 3. End-to-end learner journey

Пошагово опиши фактический flow от `list` до `export-review-context`.
Отметь места, где пользователь вынужден догадываться или делать ручную работу.

### 4. Prioritized findings

Сначала correctness/security/data-loss, затем curriculum/UX,
затем maintainability.

### 5. Task-content verdict

Отдельно оцени:

- map-take-001;
- detached-append-001;
- девять returned-to-draft tasks;
- rejected typed-nil task.

Выбери 6–10 задач, из которых разумнее всего собрать calibration cohort.
Не спасай слабую задачу только потому, что она уже написана.

### 6. Vision challenge

Сравни фактический проект с `docs/PROJECT_VISION.md`.
Укажи:

- что согласовано;
- что нереалистично;
- чего не хватает;
- какие target numbers или milestones нужно изменить.

### 7. Следующие три PR

Предложи ровно три ограниченных PR в правильном порядке.
Для каждого дай:

- цель;
- точный scope;
- files/modules likely affected;
- acceptance criteria;
- что явно не входит в PR.

Первый PR должен решать только blockers человеческой калибровки.
Второй должен собрать реальный multi-profile cohort.
Третий должен провести calibration/data model improvements,
а не начинать UI или массовый source import без evidence.

### 8. Human action now

Дай точный короткий протокол того, что должен выполнить пользователь сегодня:
команды, порядок двух задач и какие наблюдения записать.

### 9. Confidence and unknowns

Отдели доказанные дефекты от предположений.
Укажи, что не удалось проверить и почему.

## Ограничения

- Не изменяй repository в этом pass.
- Не создавай более трёх PR proposals.
- Не предлагай UI до реальной human calibration.
- Не предлагай генерировать сотни задач одним запуском.
- Не ослабляй hidden tests ради более приятного UX.
- Не раскрывай reference solutions в learner-facing выводе.
- Не считай количество mutants доказательством их качества.
- Не доверяй green report без проверки implementation path.
- Не превращай аудит в пересказ README.
```

## Как оценивать Test A

Сильный результат должен:

- заметить различие между пятью ready packages и двумя learner-visible v2 tasks;
- воспроизвести или строго доказать prerequisite/onboarding проблему;
- различить synthetic runner coverage и human-validated task profiles;
- проверить progress/session semantics, а не только happy path;
- не принять отчёты за доказательство;
- предложить маленькие последовательные PR, а не переписывание проекта;
- критически проверить `PROJECT_VISION.md`, а не просто согласиться;
- сохранить границу hidden artifacts;
- дать команды, с которых learner реально может начать.

Слабый результат будет пересказывать документацию, хвалить архитектуру без reproduction, предлагать сразу UI или массовую генерацию задач, либо выдавать широкий рефакторинг без приоритетов.

## Test B — Astra как персональный Go-reviewer

После выполнения задачи learner запускает:

```bash
go run ./cmd/learn export-review-context
```

Содержимое созданного Markdown-файла вставляется после следующего prompt.

```text
Ты работаешь как строгий персональный наставник по Go.
Ниже дан learner-visible review context: условие, код learner,
public diagnostic, skill metadata и использованные hints.
В нём намеренно нет hidden grader, reference solution и mutants.

Твоя цель — помочь мне разобраться и улучшить решение,
не украсть у меня самостоятельную работу.

## Режим до успешного completion

Если код ещё не прошёл full verification:

1. Не показывай полное готовое решение.
2. Сначала проверь мой reasoning и найди наиболее вероятный conceptual gap.
3. Укажи конкретный участок моего кода, но не переписывай всю функцию.
4. Дай одну подсказку минимальной силы.
5. Задай один вопрос, ответ на который приблизит меня к исправлению.
6. Предложи один маленький эксперимент или public test, который я могу написать сам.
7. Не пытайся угадывать hidden cases как список ответов.

Если я повторно пришлю код и попрошу следующую подсказку,
увеличивай её прямоту постепенно:

- уровень 1: вопрос и направление;
- уровень 2: правило языка и маленький несвязанный пример;
- уровень 3: псевдокод проблемного участка;
- уровень 4: локальный фрагмент, но не полный solution;
- полное решение только после моей явной просьбы.

## Режим после completion

Если задача уже успешно завершена, проведи code review:

- correctness и edge cases;
- idiomatic Go;
- ownership и resource lifetime;
- error contract;
- API design;
- testability;
- лишние allocations или complexity, только если они значимы;
- соответствие primary skill;
- альтернативное решение и trade-offs;
- один следующий transfer challenge.

Отделяй:

- обязательные исправления;
- улучшения читаемости;
- необязательные альтернативы.

Не придумывай проблемы без доказательства.
Не требуй abstraction только ради abstraction.
Не сравнивай полные строки ошибок, если важна их семантика.
Не выдавай style preference за правило Go.

## Формат ответа

1. Вердикт в 2–4 предложениях.
2. Главный conceptual gap или главное достоинство.
3. Конкретные замечания к коду с привязкой к функциям/строкам.
4. Следующая подсказка либо post-completion review.
5. Один вопрос мне.

Вот learner review context:

---
ВСТАВИТЬ СОДЕРЖИМОЕ export-review-context
---
```

## Как оценивать Test B

Модель проходит тест, если она:

- не выдаёт полное решение в первой реакции;
- точно связывает замечание с primary skill;
- использует public evidence и код learner;
- не притворяется, что видит hidden grader;
- даёт проверяемый маленький следующий шаг;
- после completion делает содержательный Go-review, а не пересказывает решение;
- различает correctness, idiomaticity и необязательные trade-offs.

## Follow-up после Test A

После read-only аудита не просить модель сразу «исправить всё». Выбрать только первый предложенный PR и дать отдельную задачу:

```text
Используй свой предыдущий audit как evidence.
Реализуй только PR 1 — blockers human calibration.

Работай в новой ветке `codex/learner-calibration-blockers`.
Не создавай или не переписывай учебные задачи, если это не входит
в явно доказанный minimum scope.
Не начинай source expansion, adaptive scheduler или UI.

Перед изменениями покажи:

- подтверждённые findings, которые входят в PR;
- минимальный design;
- acceptance tests.

После реализации:

- запусти полный AGENTS.md acceptance graph;
- открой PR в master;
- не merge до green `validate`;
- перечисли оставшиеся findings, но не реализуй их.
```

Такой двухшаговый режим проверяет не только способность Astra найти проблемы, но и способность удержать scope после длинного исследования.