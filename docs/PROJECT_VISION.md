# Go Learn: видение продукта и путь к конечному результату

Статус документа: живое долгосрочное видение.
Актуальный приоритет с 8 сентября 2026 года — полное покрытие Go и
фундаментальной стандартной библиотеки по [ROADMAP.md](ROADMAP.md): сначала
150 готовых задач, затем полный foundation-банк (ориентир 450). Backend,
PostgreSQL, production-сервис, Docker и CI как учебные темы отложены; технический
CI проекта сохраняется. Разделы ниже описывают также долгосрочное продолжение
и не расширяют активный scope.

GoLand остаётся редактором; локальный UI уже реализован. Условия, подсказки и
разборы — на английском. Теория, входная диагностика и собственный редактор
исключены из ближайшей работы. [LOCAL_UI_GOAL.md](LOCAL_UI_GOAL.md) — история
завершённой ограниченной поставки, не текущий предел работы.

Историческая контрольная точка этого видения: Phase 3D, `master` на коммите
`2f09e573c315efb2cc6c7ca1b876db4127eb952f` от 4 сентября 2026 года.

Этот документ отвечает на четыре вопроса:

1. Что именно мы строим.
2. Чем продукт должен отличаться от обычных курсов и банков задач.
3. Как выглядит конечный результат.
4. В каком порядке мы дойдём до него, не превратив проект в бесконечное
   проектирование инфраструктуры.

## 1. Короткое определение

Go Learn — локальная practice-first система для глубокого изучения Go: от
механики языка и стандартной библиотеки до production backend, конкурентности,
распределённых взаимодействий и эксплуатации сервисов.

Это не видеокурс, не интерактивная книга, не LeetCode-клон и не детская игра с
XP. Ученик последовательно решает короткие, тщательно отобранные задания,
получает проверяемую обратную связь, возвращается к важным механизмам в разных
контекстах и постепенно учится писать, тестировать, читать и исправлять
production-код на Go.

## 2. Проблема, которую решает проект

Большинство существующих ресурсов дают хотя бы одну из следующих проблем:

- слишком много текста перед маленькой задачей;
- сюжетная или детская геймификация;
- случайный набор алгоритмических упражнений без маршрута;
- поверхностное прохождение синтаксиса без нюансов Go;
- задачи, где тесты проверяют только happy path;
- отсутствие практики debugging, test-writing, API design и code review;
- переход к фреймворкам до понимания стандартной библиотеки;
- перечисление Kafka, Redis, gRPC и Docker без понимания гарантий и failure
  modes;
- отсутствие связи между ошибками ученика и следующими заданиями;
- готовое решение от нейросети вместо самостоятельного программирования.

Go Learn должен заменить это циклом:

```text
короткая задача
    -> самостоятельная попытка
    -> публичная проверка
    -> прогрессивная подсказка при необходимости
    -> скрытая полная проверка
    -> разбор решения и механики языка
    -> следующая задача с учётом пробелов
```

## 3. Для кого продукт

Основной пользователь — взрослый разработчик или технический специалист,
который хочет:

- освоить Go через постоянное написание кода;
- понять не только «как работает», но и где код ломается;
- набить руку на идиоматичном Go;
- перейти к backend- и microservice-разработке;
- уверенно работать с тестами, ошибками, I/O, context и concurrency;
- научиться читать и изменять существующие кодовые базы;
- использовать нейросеть как наставника и reviewer, а не как генератор ответа.

Проект не предполагает, что пользователь обязан быть полным новичком. Будущая
диагностика должна позволять пропускать уже освоенные блоки.

## 4. Принципы продукта

### 4.1. Практика раньше лекции

Теория появляется тогда, когда ученик упёрся в конкретный пробел. Базовый цикл:

```text
задача -> ошибка или вопрос -> точечное объяснение -> повторная попытка
```

Документация и книги используются как справочник, а не как обязательный барьер
на 300 страниц перед первой программой.

### 4.2. Короткие и взрослые постановки

Обычное задание должно иметь:

- контекст максимум в два предложения;
- явный API или список редактируемых файлов;
- от четырёх до семи проверяемых требований;
- ясные ограничения;
- не более 160 слов в основном условии;
- никакой вымышленной истории, если она не влияет на инженерную задачу.

### 4.3. Изучается механизм Go, а не декорация

Каждая задача должна иметь один ясный primary skill. Secondary skills допустимы,
но не должны скрывать фактическую сложность. Нельзя выдавать косметическую смену
имен и предметной области за новое упражнение.

### 4.4. Разные виды инженерной работы

Банк не должен состоять только из функций `Solve`.

Нужны задания на:

- implementation;
- debugging;
- refactoring;
- API design;
- test-writing и test-repair;
- integration;
- code-review patch;
- fuzzing и benchmarks;
- чтение и изменение production-style кода.

Task Contract v2 отдельно описывает activity, вид submission и verification
profile. Это позволяет честно проверять не только реализацию, но и тесты или
патч нескольких файлов.

### 4.5. Ошибочные решения — часть спецификации

Для важной задачи недостаточно reference solution. В ней должны быть
реалистичные defect models или mutants: типичные неправильные решения, которые
обязаны обнаруживаться grader-ом.

Примеры:

- перепутаны байты и руны;
- возвращён aliased slice;
- проигнорирована ошибка;
- сравнивается текст ошибки вместо семантики;
- оставлен typed nil внутри interface;
- выбран неверный receiver;
- нарушена отмена context;
- осталась goroutine leak;
- тест проверяет happy path, но пропускает границу.

### 4.6. Локальность и воспроизводимость

Основной режим продукта — local-first и offline:

- код ученика остаётся на его машине;
- progress хранится локально;
- canonical tasks не изменяются;
- grader, mutants и reference не попадают в learner workspace;
- проверки выполняются в disposable workspace;
- внешние зависимости и сеть не требуются для foundation tasks;
- результаты инструментов детерминированы там, где это возможно.

Текущая изоляция — защита рабочего дерева и окружения, а не kernel-level
sandbox. Запуск кода неизвестных третьих лиц не является поддерживаемым сценарием
до отдельного sandbox решения.

### 4.7. Качество важнее количества

Нельзя заполнять coverage низкоценными задачами ради числа. Задача попадает в
`ready` только после технической, педагогической и provenance-проверки.

Массовая генерация одним запросом запрещена. Производство идёт небольшими
batch-ами с независимым review и периодической дедупликацией.

### 4.8. Нейросеть помогает, но не является источником истины

Модель может выступать как:

- исследователь источников;
- куратор кандидатов;
- автор черновика;
- red-team reviewer;
- наставник с прогрессивными подсказками;
- reviewer решения ученика;
- планировщик следующего задания.

Но модель не должна:

- автоматически объявлять собственную задачу качественной;
- заменять тесты убедительным объяснением;
- терять provenance;
- выдавать готовое решение до попытки;
- массово генерировать контент без независимой проверки;
- менять taxonomy или contract без воспроизводимого дефекта.

## 5. Как должен выглядеть конечный продукт

### 5.1. Foundation track

Полное погружение в язык и стандартную библиотеку:

- toolchain, modules и package boundaries;
- types, expressions, control flow и functions;
- arrays, slices, strings, bytes, runes, Unicode и maps;
- structs, methods, pointers, embedding и interfaces;
- error contracts, wrapping, `errors.Is`, `errors.As`, panic/recover boundaries;
- generics;
- `io.Reader`, `io.Writer`, files, buffering и encoding;
- time;
- unit tests, table-driven tests, subtests, helpers и test doubles;
- fuzzing и benchmarks;
- goroutines, channels, sync primitives и context;
- ownership, cancellation, termination, race conditions и leak prevention.

Текущая taxonomy v3 содержит 142 assignable skills. Полный foundation bank
должен удовлетворить её encounter requirements, а основной маршрут — дать
достаточную глубину без обязательного прохождения всех альтернатив.

### 5.2. Backend track

Переход от знания языка к созданию сервисов:

- `net/http`, routing, handlers и middleware;
- JSON, validation и единая модель ошибок API;
- `httptest` и HTTP clients;
- timeouts, cancellation и graceful shutdown;
- configuration и structured logging;
- PostgreSQL, SQL, pgx/pgxpool, migrations и transactions;
- isolation levels, locking, pagination и pool management;
- unit и integration tests;
- Testcontainers;
- worker pools, pipelines, batching и backpressure;
- Docker, Compose и CI.

Сначала используется стандартная библиотека. Фреймворк или router добавляется
только после понимания базовых контрактов.

### 5.3. Production и distributed track

Практика поведения системы при сбоях:

- metrics, tracing и profiling;
- health/readiness и lifecycle;
- retries, rate limiting и overload;
- Redis и cache semantics;
- gRPC и Protobuf;
- один message broker с глубоким пониманием модели;
- delivery guarantees, duplicates, ordering и idempotency;
- dead-letter handling и transactional outbox;
- vulnerability scanning;
- failure injection и recovery scenarios.

### 5.4. Сквозные проекты

Маршрут должен содержать небольшие законченные компоненты и несколько
постепенно усложняющихся capstone-проектов, а не один гигантский проект без
обратной связи.

Основной production capstone может быть Focus Session Service:

```text
domain model
    -> tested package API
    -> HTTP service
    -> PostgreSQL
    -> background workers
    -> observability
    -> gRPC или message broker
    -> idempotency и outbox
    -> deployment и failure drills
```

### 5.5. Интерфейс

CLI остаётся стабильным engine и диагностическим интерфейсом. Поверх него
появляется локальный взрослый UI:

- текущая задача и редактор;
- public check и complete;
- прогрессивные подсказки;
- roadmap и prerequisites;
- browser задач;
- review после решения;
- skill progress и due repetitions;
- keyboard-first управление;
- светлая и тёмная темы;
- отсутствие XP, streaks и игровых валют.

UI не должен дублировать scheduler, progress и verification logic.

### 5.6. Работа с моделями

Минимальная интеграция уже существует через `export-review-context`: ученик
может безопасно передать модели условие, свой код, публичный результат и
использованные подсказки без grader-а, mutants и reference solution.

Позднее допустим optional model adapter, но продукт должен полностью работать
без API-ключа и облака.

## 6. Количественные ориентиры

Это ориентиры планирования, а не цель ради красивого числа.

- Human-calibrated pilot: 8–12 ready tasks, покрывающих все основные submission
  profiles.
- Foundation alpha: около 60 качественных задач.
- Foundation beta: 120–180 задач и стабильный локальный UI.
- Foundation core route: ориентировочно 250–350 задач.
- Foundation full bank: не менее 409 primary encounters; вероятный размер
  450–650 задач с альтернативами и remedial practice.
- Backend/production core: ориентировочно 180–260 задач.
- Полный curated bank: ориентировочно 800–1,200 задач.
- Основной персонализированный маршрут не обязан включать весь банк; ожидаемый
  порядок величины — 450–650 заданий от основ до production backend.

Размеры корректируются по фактической taxonomy coverage, learner feedback и
семантической дедупликации.

## 7. Текущее состояние

На контрольной точке Phase 3D уже существуют:

- taxonomy v3: 28 domains и 142 assignable skills;
- source registry и pinned source commits;
- импорт и полный reviewed triage 177 Exercism Go candidates;
- pilot seed plan;
- Task Contract v2;
- provenance v2 и проектная Apache-2.0 policy;
- lifecycle `draft -> in-review -> ready/rejected`;
- deterministic validator и quality reports;
- fixed-tests, mutation-tests и workspace-tests verification engines;
- public/grader/reference/mutant separation;
- single-pass CheckCollection;
- CI с tests, race, vet и integration checks;
- локальный learner CLI и progress store;
- безопасный export review context.

Исторический content inventory на указанном выше Phase 3D commit
(актуальная сводка находится в ROADMAP.md):

- 5 ready packages;
- из них 3 legacy v1 packages, не показываемые learner CLI;
- 2 пользовательские Task Contract v2 задачи:
  - `map-take-001`;
  - `detached-append-001`;
- 9 v2 packages возвращены в draft после red-team;
- 1 v2 package отклонён;
- все 15 lifecycle packages проходят технические проверки, но draft-задачи не
  считаются готовым учебным контентом.

Это означает, что инфраструктурный вертикальный срез завершён. Следующая
неизвестность — не «может ли CLI запустить synthetic fixture», а «насколько
полезны, понятны и приятны реальные задачи живому ученику».

## 8. Исторический сценарий первого human trial

Этот сценарий относится к исходному банку из двух v2-задач. Сейчас начинайте
с `learn next` по последовательному маршруту; ночной goal не требует участия
пользователя. Историческая инструкция: пройти две v2-задачи без просмотра grader-а и reference solution.
Начать лучше с `map-take-001`, затем перейти к `detached-append-001`.

```bash
git switch master
git pull --ff-only

go run ./cmd/learn list
go run ./cmd/learn show map-take-001
go run ./cmd/learn start map-take-001
```

CLI напечатает путь к learner workspace. Изменяется только код внутри этого
workspace.

```bash
go run ./cmd/learn check
# после успешных public checks
go run ./cmd/learn complete
```

После задачи сохранить обратную связь:

```bash
go run ./cmd/learn feedback "clarity=4/5; difficulty=2/5; actual=18m; public-feedback=4/5; learned=comma-ok with stored zero; frustration=none"
go run ./cmd/learn export-review-context
```

Для второй задачи повторить цикл:

```bash
go run ./cmd/learn start detached-append-001
go run ./cmd/learn check
go run ./cmd/learn complete
go run ./cmd/learn feedback "clarity=...; difficulty=...; actual=...; public-feedback=...; learned=...; frustration=..."
go run ./cmd/learn export-review-context
```

Минимальный feedback contract:

- clarity: насколько понятно условие, 1–5;
- difficulty: субъективная сложность, 1–5;
- actual: фактическое активное время;
- public-feedback: помогли ли открытые проверки, 1–5;
- learned: что стало понятнее;
- frustration: что мешало — задача, формулировка, API, CLI или проверка;
- next: хотелось ли сразу открыть следующую задачу.

Не нужно пытаться угадать «правильный отзыв». Раздражение, скука и желание
обойти условие — важные продуктовые сигналы.

## 9. Ближайший план

### Milestone 1 — Human trial текущего вертикального среза

Цель: пройти две ready v2-задачи и проверить реальный learner loop.

Выход:

- два completion record;
- фактическое время;
- feedback;
- два review-context export;
- список дефектов CLI и постановок, если они проявились.

До завершения trial нельзя масштабировать текущий формат на сотни задач.

### Milestone 2 — Реальный pilot всех verification profiles

Цель: довести до `ready` минимум по одной качественной задаче на:

- fixed-tests implementation/debugging;
- mutation-tests test-writing/test-repair;
- workspace-tests integration;
- refactoring или API design.

Приоритетные draft-кандидаты для ремонта:

- `clamp-tests-001` или `rune-offset-tests-001` — живой test-submission flow;
- `stream-normalize-001` — живой workspace-patch flow;
- `counter-receiver-001` или `formatter-contract-001` — receiver/interface
  design;
- `utf8-prefix-001` — Unicode, только после исправления границы изучаемого
  механизма.

`typed-nil-resource-001` не восстанавливается косметически; typed-nil задача
проектируется заново с узким API и наблюдаемой семантикой interface value.

Выход: 8–12 human-testable ready tasks с разнообразными activities.

### Milestone 3 — Калибровка пилота

На основании реальных прохождений исправляются:

- wording;
- difficulty и estimated time;
- public diagnostics;
- hints;
- mutant relevance;
- unfair hidden requirements;
- scheduler ordering;
- CLI friction.

Инфраструктура меняется только при дефекте, воспроизведённом на реальной задаче
или learner workflow.

### Milestone 4 — Расширение источников и unified seed catalog

По одному подключаются permissive sources, которые закрывают доказанные gaps:

1. официальный Tour/website material;
2. официальный `golang/example`;
3. Learn Go with Tests;
4. Ardan Labs Go Training;
5. дополнительные sources только при уникальном coverage.

Для каждого источника выполняются pinned sync, harvest, license/provenance audit,
pedagogical triage и дедупликация. Затем создаётся единый foundation seed
catalog.

### Milestone 5 — Foundation alpha

Небольшими batch-ами создаются первые 60 ready tasks. После каждых 3–4 batch-ов
проводится cross-batch audit:

- semantic duplicates;
- повторяющиеся API;
- перекос task types;
- слабые mutants;
- excessive seed reuse;
- необоснованный coverage через secondary skills.

На 60 задачах добавляется первая mastery model и диагностический маршрут.

### Milestone 6 — Local UI beta

После стабилизации CLI и 60–100 задач создаётся локальный UI поверх существующего
engine. Контент и scheduler не переносятся во frontend.

### Milestone 7 — Foundation core и full bank

Производство продолжается batch-ами до полноценного foundation route и
taxonomy coverage. Параллельно измеряется:

- completion rate;
- actual time;
- hint usage;
- solution reveals;
- повторные ошибки;
- learner-rated clarity и usefulness.

### Milestone 8 — Backend и production tracks

Создаются отдельная backend taxonomy, source plan, pilot и production batches.
Сначала 12-task vertical slice, затем red-team, human trial и масштабирование.

### Milestone 9 — Release v1

Готовятся:

- reproducible public learner bundle;
- desktop/local UI package;
- progress migrations;
- public export без grader/reference/mutants;
- deterministic source and provenance manifests;
- THIRD_PARTY_NOTICES;
- installation, learner и contributor guides;
- clean-machine offline smoke test.

## 10. Правила принятия решений

1. Реальный learner defect важнее гипотетической архитектурной красоты.
2. Зелёные тесты не доказывают педагогическую ценность.
3. Красивое условие не компенсирует слабый grader.
4. Четыре одинаковых mutants не считаются четырьмя проверками.
5. Coverage не засчитывается для draft, rejected, source evidence или seed.
6. Новая задача не создаётся, если она семантически дублирует существующую.
7. Новая инфраструктура не строится без конкретного сценария, который текущий
   contract не способен выразить или проверить.
8. UI начинается после подтверждения learner loop, а не вместо него.
9. Модель предлагает изменения; repository contracts, tests и review решают,
   принимаются ли они.
10. После каждого существенного milestone пользователь снова решает задачи.

## 11. Definition of Done конечного продукта

Проект считается достигшим конечной цели, когда пользователь может:

- начать с диагностики текущих знаний;
- пройти адаптивный маршрут по фундаменту Go;
- несколько раз встретить каждый важный механизм в разных activity;
- писать не только реализации, но и tests, fixes и multi-file patches;
- понимать ошибки, interfaces, I/O, context и concurrency на уровне причин, а не
  рецептов;
- перейти к созданию, тестированию и эксплуатации backend services;
- пройти production failure scenarios и capstone projects;
- работать полностью локально без обязательной подписки или API-ключа;
- запросить помощь модели без утечки hidden verification material;
- видеть честный skill progress вместо vanity metrics;
- экспортировать и переносить progress;
- использовать стабильный CLI или взрослый локальный UI.

Со стороны content system должны выполняться:

- полное обязательное coverage маршрутов;
- отсутствие surviving required mutants;
- complete provenance для каждой adapted/combined task;
- human calibration для core tasks;
- deterministic validation и release export;
- отсутствие hidden artifacts в публичном learner bundle;
- документированная поддержка текущей версии Go.

## 12. Brief для аудита новой моделью

Этот блок можно целиком передать сильной модели вместе со ссылкой на
репозиторий. На 4 сентября 2026 года подходящим кандидатом является GPT-6 Astra,
но review protocol не зависит от конкретного поставщика или модели.

```text
Ты проводишь независимый read-only аудит репозитория
EarthMoonTransit/go-learn.

Сначала прочитай:

- docs/PROJECT_VISION.md
- README.md
- AGENTS.md
- CONTRIBUTING.md
- docs/TASK_CONTRACT_V2.md
- docs/PROVENANCE_V2.md
- docs/LEARNER_CLI.md
- docs/QUALITY_GATES.md
- reports/task-quality.md
- reports/pilot-v2-red-team.md
- reports/pilot-v2-promotion.md
- reports/learner-cli-acceptance.md

Затем исследуй реализацию CLI, learner store, verification runners,
task lifecycle, current ready/draft tasks и CI.

На первом проходе ничего не изменяй и не создавай массово.
Не доверяй README и отчётам без проверки кода.
Не предлагай полный rewrite только потому, что можешь придумать другую
архитектуру.

Ответь:

1. В каком фактическом состоянии находится проект относительно PROJECT_VISION.
2. Готов ли текущий learner loop к живому прохождению двух ready v2 задач.
3. Какие дефекты являются блокирующими, а какие можно проверить только
   реальным learner trial.
4. Есть ли утечки grader/reference/mutants или способы обойти workspace
   boundaries.
5. Корректно ли сохраняются progress, attempts, hints, solution reveal,
   completion и feedback.
6. Справедливы ли public/full diagnostics и scheduler при текущем маленьком
   наборе задач.
7. Какие 3–5 draft tasks разумнее всего исправить для проверки
   mutation-tests, workspace-tests и refactoring на живом ученике.
8. Какие части инфраструктуры сейчас избыточны и не должны развиваться до
   появления фактического дефекта.
9. Какие три следующих PR дадут максимальную продуктовую ценность при
   минимальном риске.
10. Что пользователь должен сделать руками до следующего content batch.

Для каждого finding укажи:

- severity: blocker / high / medium / low;
- конкретный файл и участок кода;
- воспроизводимый сценарий;
- почему это мешает цели продукта;
- минимальное исправление;
- что произойдёт, если ничего не менять.

В конце дай один из трёх вердиктов:

- READY_FOR_HUMAN_TRIAL;
- FIX_BEFORE_HUMAN_TRIAL;
- ARCHITECTURE_BLOCKED.

Если вердикт READY_FOR_HUMAN_TRIAL, не придумывай ещё одну обязательную
инфраструктурную фазу. Сформулируй точный протокол прохождения и сбора feedback.
```

## 13. Обновление документа

Документ обновляется только на границах существенных milestone:

- изменилось конечное product vision;
- изменились обязательные tracks;
- получен human trial evidence;
- принята новая taxonomy major version;
- изменился Task Contract;
- начался UI или backend track;
- утверждён release scope.

Мелкие технические изменения и отдельные batch-и сюда не добавляются. Их место —
ADR, reports и commit history.
