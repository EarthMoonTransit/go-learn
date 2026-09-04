# Go Practice Bank: продуктовая цель и путь к ней

Статус документа: canonical product direction  
Текущий снимок: `master` на `2f09e573c315efb2cc6c7ca1b876db4127eb952f`, 4 сентября 2026 года  
Технические решения фиксируются в ADR. Этот документ отвечает на вопросы **зачем существует проект, каким должен стать конечный продукт и в каком порядке мы к нему идём**.

## 1. Что мы строим

Go Practice Bank — локальная practice-first система для глубокого изучения Go. Это не видеокурс, не электронная книга, не сборник случайных kata и не ещё один LeetCode-клон.

Система должна провести человека от знакомства с языком до способности писать, тестировать, отлаживать и сопровождать production Go-код. Основной учебный цикл:

```text
короткая техническая задача
        ↓
самостоятельная попытка
        ↓
видимые проверки и прогрессивные подсказки
        ↓
скрытая проверка поведения
        ↓
разбор конкретных ошибок и идиом Go
        ↓
следующая задача с учётом обнаруженных пробелов
```

Теория появляется тогда, когда learner упёрся в конкретный механизм. Задачи должны заставлять не только реализовывать функции, но и находить дефекты, писать тесты, проектировать API, читать чужой код, работать с конкурентностью и собирать небольшие production-подобные компоненты.

## 2. Какую проблему решает проект

У существующих платформ обычно есть хотя бы один из следующих недостатков:

- длинные сюжетные условия и декоративная геймификация;
- много объяснений и мало самостоятельного кода;
- хорошая практика алгоритмов, но слабая практика именно Go;
- отсутствие последовательного пути от синтаксиса до production engineering;
- одинаковая программа для всех без учёта реальных ошибок learner;
- слабые тесты, которые проверяют happy path, но пропускают неверное понимание языка;
- непрозрачное происхождение учебных материалов;
- платная оболочка вокруг информации, доступной в документации и открытых репозиториях.

Go Practice Bank объединяет лучшее из открытых банков задач, официальной документации, production-примеров и оригинальных упражнений. Исходный материал проходит аудит, дедупликацию, переработку, mutation review и человеческую калибровку.

## 3. Для кого предназначен продукт

Первичный пользователь — разработчик, который хочет освоить Go через практику и в дальнейшем писать backend- и infrastructure-код. Он может быть новичком в Go, но продукт не должен разговаривать с ним как с ребёнком.

Система должна поддерживать три точки входа:

1. Полный новичок начинает с toolchain, типов, выражений и управления потоком.
2. Разработчик из другого языка проходит диагностику и пропускает очевидные основы.
3. Go-разработчик использует отдельные ветки для deliberate practice: ошибки, интерфейсы, тестирование, конкурентность, производительность или production backend.

## 4. Принципы продукта

### Practice first

Большая часть времени уходит на написание, чтение, тестирование и исправление кода. Теория короткая, точечная и привязана к обнаруженному пробелу.

### Взрослая подача

Нет персонажей, заклинаний, валют, XP и искусственных streaks. Условие описывает технический контекст, API, наблюдаемое поведение и ограничения.

### Go-specific depth

Задачи должны раскрывать не только общий алгоритм, но и механику Go: zero values, backing arrays, method sets, typed nil, error chains, `io.Reader`/`io.Writer`, ownership горутин и каналов, cancellation, data races, package boundaries и стандартную библиотеку.

### Несколько форм деятельности

Освоение навыка нельзя подтверждать четырьмя почти одинаковыми implementation-задачами. Используются implementation, debugging, refactoring, API design, test writing, workspace integration, fuzzing, benchmarks и code review.

### Observable quality

Каждое требование должно быть проверяемым. Mutants отражают реальные заблуждения и должны быть семантически различными. Количество tests или mutants само по себе не является показателем качества.

### Local first

Задачи, workspace, progress и grader работают локально. По умолчанию нет аккаунта, telemetry, cloud state и обязательной подписки.

### AI усиливает обучение, а не решает за learner

Модель используется как куратор, наставник и reviewer. Она не должна автоматически выдавать полный ответ до самостоятельной попытки. Взаимодействие строится через learner-visible review context без grader, reference solution и скрытых ожиданий.

### Provenance by construction

Для адаптированных материалов сохраняются source, exact commit, license, artifact hashes, роль источника и характер преобразования. Reference-only материалы не копируются и не маскируются под оригинальные.

### Evidence before infrastructure

Новая инфраструктура добавляется только для дефекта, который проявился на реальных задачах или в реальном learner flow. Проект не должен бесконечно совершенствовать фабрику, не обучая пользователя.

## 5. Каким должен быть конечный продукт

В конечном состоянии пользователь открывает минималистичное локальное приложение или CLI, проходит короткую диагностику и получает персональный маршрут.

Каждая учебная сессия содержит:

- одну ясную задачу;
- небольшой редактируемый workspace;
- public checks с полезной диагностикой;
- прогрессивные hints;
- full verification без утечки hidden cases;
- self-review и AI-review после решения;
- объяснение только тех механизмов, где обнаружен пробел;
- следующую задачу, выбранную с учётом prerequisites, разнообразия деятельности и отложенного повторения.

### Foundation track

Foundation должен охватить toolchain, типы, функции, выражения, control flow, arrays, slices, strings, Unicode, maps, structs, methods, pointers, embedding, interfaces, errors, generics, packages, I/O, encoding, time, testing, fuzzing, benchmarks, goroutines, channels, `sync` и `context`.

Плановые уровни:

- **Foundation Core** — основной маршрут примерно из 240–320 задач. Он даёт рабочее владение фундаментом без обязательного прохождения всех альтернатив.
- **Foundation Full** — полный deliberate-practice банк. Taxonomy v3 уже требует минимум 409 primary и 693 total encounters; практический размер полного foundation-банка, вероятно, составит 500–700 задач.
- **Foundation Extra** — дополнительные алгоритмические drills, альтернативные реализации, performance challenges и углубления.

Точные размеры определяются coverage, diversity и human calibration, а не заранее выбранным красивым числом.

### Backend and Production track

После фундамента пользователь проходит отдельный маршрут:

- `net/http`, routing, handlers, middleware и HTTP clients;
- validation, JSON, timeouts, cancellation и graceful shutdown;
- PostgreSQL, SQL, `pgx`, pools, migrations и transactions;
- unit, integration и testcontainers-based testing;
- worker pools, pipelines, batching и backpressure;
- logging, metrics, tracing, profiling и health/readiness;
- gRPC, Redis и один message broker;
- retries, idempotency, delivery guarantees и transactional outbox;
- Docker, Compose, CI и vulnerability scanning.

Backend/production core ориентировочно содержит 180–260 задач. Сквозной capstone постепенно превращает небольшой Focus Session Service в наблюдаемый и отказоустойчивый production-подобный сервис.

### Итоговый объём

Плановый основной путь foundation + backend — около 420–580 задач. Полный банк с альтернативами и remedial practice может вырасти до 800–1200 задач. Большой размер допустим только при отсутствии семантических дублей и при сохранении качества каждого задания.

## 6. Архитектура конечной системы

Проект состоит из шести слоёв.

### Источники и каталог

`source-sync`, adapters и harvest получают pinned snapshots разрешённых источников. Candidate catalog хранит метаданные, а не автоматически публикует учебный контент.

### Curriculum

Taxonomy описывает атомарные навыки, prerequisites, типичные заблуждения, mastery criteria и требуемое число встреч. Routes и checkpoints задают учебную последовательность.

### Task authoring and QA

Task Contract определяет activity, submission kind и verification profile. Lifecycle не позволяет отправить задачу в `ready` без технического, педагогического и provenance review.

### Verification engine

Checker запускает public, grader, reference, mutant, vet, fuzz, benchmark и race-профили в изолированных temporary workspaces. Детерминированные reports позволяют контролировать качество всего банка.

### Learner engine

Learner service управляет приватными workspace, progress, hints, completion, feedback, review export, prerequisites и scheduler. CLI и будущий UI используют один и тот же engine.

### AI review bridge

`export-review-context` формирует безопасный learner-visible пакет. Модель анализирует задачу, код, public diagnostics, skill metadata и уже использованные hints, но не получает grader или reference solution. В будущем возможна опциональная локально управляемая интеграция, однако ручной экспорт остаётся поддерживаемым и прозрачным режимом.

## 7. Определение успеха

Проект считается успешным не тогда, когда в каталоге лежит тысяча директорий, а когда learner после основного маршрута способен:

- писать идиоматичный Go без постоянного копирования шаблонов;
- объяснять ownership данных и жизненный цикл ресурсов;
- проектировать небольшие устойчивые package APIs;
- корректно использовать errors, interfaces, I/O и context;
- доказывать завершение goroutines и отсутствие data races;
- писать deterministic unit и integration tests;
- находить дефекты в чужом коде и объяснять причину;
- собирать HTTP/DB/background-worker сервис;
- понимать поведение системы при timeout, cancellation, duplicate delivery, shutdown и перегрузке;
- измерять производительность до оптимизации;
- читать production Go-код и делать ограниченные безопасные изменения.

Для релизной v1 дополнительно требуются:

- воспроизводимый install и offline run;
- стабильные migrations progress и task contracts;
- завершённый Foundation Core;
- learner CLI и локальный UI поверх единого engine;
- отсутствие hidden artifacts в public export;
- полный license/provenance manifest;
- зелёный CI и отсутствие surviving required mutants;
- human-calibrated difficulty и estimated time для core route.

## 8. Текущее состояние

Текущий `master` находится в Phase 3D.

Уже готовы:

- Taxonomy v3: 28 domains и 142 assignable skills;
- минимальный coverage plan: 409 primary и 693 total encounters;
- 12 аудированных sources;
- 177 reviewed Exercism candidates и seed plan;
- Task Contract v2 и multi-source Provenance v2;
- lifecycle, isolated verification, mutation checks и deterministic reports;
- локальный learner CLI;
- команды `list`, `show`, `start`, `resume`, `check`, `complete`, `hint`, `review`, `solution`, `next`, `progress`, `feedback`, `reset-task` и `export-review-context`;
- зелёный CI на текущем `master`.

В lifecycle находятся 15 task packages:

- 5 `ready`;
- 9 `draft`;
- 1 `rejected`.

Из пяти ready-задач три являются legacy v1 и намеренно не показываются learner CLI. Сейчас через CLI доступны только две реальные v2-задачи:

- `detached-append-001` — debugging backing-array aliasing;
- `map-take-001` — работа с map lookup, zero value и delete.

Следовательно, проект достиг **технической learner alpha**, но ещё не достиг **curriculum alpha**. CLI можно и нужно проверять на человеке, однако текущие две задачи не образуют последовательный курс с нуля.

## 9. Что требуется от learner прямо сейчас

Следующая фаза — не массовая генерация, а human smoke test.

Использовать отдельный learner home:

```bash
export GO_LEARN_HOME="$HOME/.go-learn-pilot"

go run ./cmd/learn list
go run ./cmd/learn next

go run ./cmd/learn show map-take-001
go run ./cmd/learn start map-take-001
# изменить код в напечатанном workspace
go run ./cmd/learn check
go run ./cmd/learn hint             # только при необходимости
go run ./cmd/learn complete
go run ./cmd/learn feedback "..."
go run ./cmd/learn export-review-context
```

Затем повторить для `detached-append-001`.

Во время прохождения фиксировать не только bugs, но и учебный опыт:

- понятно ли условие без скрытых тестов;
- известны ли prerequisites или задача начинается слишком поздно;
- ясно ли, какой файл менять;
- полезен ли public diagnostic;
- справедлива ли full verification;
- не раскрывает ли hint слишком много;
- соответствует ли difficulty фактической сложности;
- сколько было public checks и full attempts;
- что именно стало понятно после задачи;
- захотелось ли сразу перейти к следующей.

Это smoke test CLI и двух content packages. Он не является доказательством качества всего curriculum.

## 10. Ближайшие подтверждённые риски

Перед широкой human calibration необходимо проверить и при необходимости исправить следующие области.

### Нет корректной точки входа

Обе learner-visible задачи имеют prerequisites, которые нельзя закрыть текущим v2 route. `next` способен выбрать ближайшую задачу с отсутствующими prerequisites, но это не заменяет начальный маршрут. Нужны diagnostic assumptions либо несколько entry tasks.

### Недостаточно реальных submission profiles

Synthetic integration tests покрывают implementation, tests и workspace-patch, но обе реальные ready v2-задачи используют implementation submission. До масштабирования требуется хотя бы одна human-reviewed test-writing task и одна workspace task.

### Время активности пока приблизительное

Progress хранит одну активную сессию от `start`/`resume` до `check`/`complete`. Прерывание работы, повторный `resume` или возвращение на следующий день может исказить duration. Для калибровки времени нужны session history и явная pause/resume policy.

### Crash recovery progress lock

`.progress.lock` создаётся через exclusive file creation. После аварийного завершения процесса stale lock может потребовать ручного удаления. Нужен безопасный `doctor`/recovery workflow или настоящий process-aware lock.

### `review` ещё не является полным code review

Текущая команда показывает mastery prompts и список learner files. Глубокий разбор предполагает `export-review-context` и внешнюю модель. Это необходимо явно отражать в UX либо расширить review flow.

### Calibration data недостаточно детальна

Один счётчик `attempts` смешивает public checks и full completion attempts; хранится только последняя public diagnostic. Для содержательной калибровки полезны отдельные counters и bounded event history без сохранения grader data.

### Документационные противоречия

`AGENTS.md` всё ещё содержит старую формулировку о невозможности internal original source до появления canonical remote, хотя remote и Provenance v2 уже существуют. Правила необходимо синхронизировать с ADR 0008 и текущим contract.

Эти пункты не требуют останавливать первый smoke test. Они являются приоритетным входом для следующего ограниченного PR после получения фактической обратной связи.

## 11. Дорожная карта

### Milestone 1 — Human smoke calibration

Пройти две ready v2-задачи через CLI. Проверить onboarding, workspace, public diagnostics, hints, complete, feedback и review export. Результат — versioned calibration report с конкретными наблюдениями.

### Milestone 2 — Learner-flow hardening

Исправить только подтверждённые проблемы smoke test. Предварительный scope: entry route, activity sessions, stale lock recovery, distinct attempt counters, ясный review/export flow и синхронизация AGENTS/ADR.

Exit criterion: learner может начать, прервать, продолжить, проверить, завершить и отдать задачу модели на review без ручного поиска внутренних файлов.

### Milestone 3 — Calibration cohort из 8–12 задач

Вернуть в review лучшие из девяти drafts, не пытаясь спасти каждую задачу. Cohort должен включать:

- implementation и debugging;
- минимум две test-submission задачи;
- минимум одну workspace-patch задачу;
- interfaces/errors/Unicode/I/O наряду со slices и maps.

Каждая задача проходит новый independent review. Один content PR содержит не более 10 новых или существенно изменённых задач.

### Milestone 4 — Полная human calibration

Learner проходит cohort. Difficulty, time, hints, public diagnostics, grader fairness и AI-review flow корректируются по данным, а не по предположениям.

Exit criterion: минимум 8 качественных задач разных profiles пройдены человеком; нет P0 defects в CLI; calibration schema пригодна для следующих batches.

### Milestone 5 — Source expansion и unified seed catalog

Последовательно подключаются `golang/website`, `golang/example`, Learn Go with Tests и Ardan Labs. Каждый source получает pinned import, triage, duplicate clusters и provenance audit. После этого строится cross-source seed catalog и точный foundation route.

### Milestone 6 — Foundation Alpha: 60 задач

Создаются небольшие content batches по 6–10 задач. После каждых трёх-четырёх batches проводится semantic duplicate и learner-experience audit.

На 60 задачах появляются:

- начальный diagnostic;
- checkpoints;
- более содержательный deterministic scheduler;
- стабильный CLI workflow;
- основание для UI.

### Milestone 7 — Local UI Beta

Минималистичный keyboard-first UI строится поверх learner engine, а не дублирует его. Он поддерживает task view, editor, public/full check, hints, review export, roadmap и progress. Нет telemetry и cloud requirement.

### Milestone 8 — Foundation Core

Банк расширяется до полного core route. Coverage оценивается вместе с task-form diversity, spaced repetition и human calibration. Формальное закрытие skill ID без реального mastery не принимается.

### Milestone 9 — Backend/Production Track

Создаётся отдельная backend taxonomy, source plan, 12-task pilot и batches. Capstone развивается по стадиям и остаётся проверяемым после каждого шага.

### Milestone 10 — Stable v1

Фиксируются contracts, migrations, public export, license notices, installer, documentation и release criteria. Clean-machine offline smoke test подтверждает, что пользователь может установить продукт и пройти core route без доступа к private grader internals.

## 12. Правила дальнейшей разработки

1. Любое изменение идёт через feature branch и зелёный PR.
2. Tooling, taxonomy и content не смешиваются в одном большом batch без необходимости.
3. Не более 10 новых или существенно изменённых задач на content PR.
4. Ready-задача требует technical, pedagogical, mutation и provenance review.
5. После каждого небольшого набора задач проводится реальная learner trial.
6. Новый framework или abstraction добавляется только для воспроизведённого дефекта.
7. Нельзя повышать coverage low-value задачами или cosmetic reskins.
8. Reference-only content не копируется.
9. AI-generated code и tasks проходят те же gates, что человеческие.
10. В случае конфликта между количеством и качеством уменьшается batch, а не gate.

## 13. Что проектом не является

Go Practice Bank не должен превращаться в:

- детскую игровую платформу;
- библиотеку из тысяч непроверенных AI-задач;
- набор только алгоритмических интервью-задач;
- линейный текстовый курс на сотни страниц;
- список технологий из вакансии без глубины;
- cloud SaaS с обязательным аккаунтом;
- систему, где модель пишет решение вместо learner;
- универсальную платформу для всех языков до завершения Go v1.

## 14. Следующее решение

На текущем этапе правильная последовательность:

```text
две реальные задачи через CLI
        ↓
human feedback и Astra review
        ↓
один ограниченный learner-flow hardening PR
        ↓
8–12 real-task calibration cohort
        ↓
повторная human calibration
        ↓
source expansion и foundation batches
```

До завершения первого human smoke test не следует начинать массовое создание задач или UI. Сейчас главный источник истины — не количество реализованных команд, а фактический опыт первого learner.