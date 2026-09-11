# Go Learn: дальнейший foundation-курс

Редакция 2, 2026-09-10. Это единая точка входа в обновлённый план реализации. Документ и приложения опубликованы как planning-only: учебный контент, код, схемы, registry и learner state этим PR не изменяются.

Обязательные приложения:

- [Карта 47 подтем, глубина и пересмотр Exercism](foundation-topic-map.md).
- [Дополнительные банки: проверенные версии, лицензии, ограничения и 24 кандидата](../research/exercise-banks-2026-09-10.md).

Основание: аудит repository snapshot `5f925a757f4bba63396d3b4b755776da75e68cec`; актуальность master повторно проверена 2026-09-10. В этой сессии выполнено статическое чтение, не запуск Go Learn, браузера, upstream code или полного validate. Доказательства прежних acceptance находятся в самом репозитории; они не являются новым прогоном автора плана. Private learner history не читалась.

## 0. Принятый вариант и изменения относительно первой редакции

Учебный цикл: короткая теория → упражнения → самостоятельное применение → разбор ошибок → отложенный возврат. GoLand остаётся редактором, UI и короткие learn-команды — навигацией и единым доступом к проверке.

Уточнение пользователя 2026-09-10 упрощает теорию: локальная страница с основным объяснением и небольшим примером, затем точные ссылки на официальную документацию для подробностей. Не строим редактор, MDX-платформу, импортёр книг, обязательные анимации или собственную копию всей документации Go. Сначала полезные страницы, затем улучшение глубины и покрытия.

Основной план:

1. Сохранить существующие 150 задач, решения, номера и историю. Добавить read-only теорию и подтемы поверх текущего маршрута, не ждать большого перестроения scheduler.
2. Расширить source-first отбор Exercism/Ardan источниками SHAD, Golings, Birsi и выборочно Go Koans. Не обещать число адаптаций или долю внешних задач до review.
3. Закрывать реальные пробелы и недостаток самостоятельности, а не производить сотни похожих функций ради 450. Приоритет: разрешённое готовое упражнение → содержательная адаптация → original с recorded gap reason.
4. Алгоритмические упражнения размещать внутри основных тем, сохранять activity implementation/test-writing/debugging и показывать дополнительный badge. Несложные алгоритмы допускаются в core, большие работы — после prerequisites.
5. Изменения порядка маршрута, evidence и повторений выполнять отдельно через ADR/миграцию, после работающей теории. Ни просмотр страницы, ни отсутствие пользовательских решений не являются mastery.

Backend/HTTP/PostgreSQL/production-service/Docker/CI как учебные темы по-прежнему вне текущего этапа. Технический CI самого репозитория сохраняется. Learn Go with Tests не используется для новых задач или как основной источник теории; уже принятые 16 адаптаций сохраняются.

Эти новые продуктовые решения заменяют прежнее исключение теории и прежний запрет алгоритмичности. Они не отменяют AGENTS.md, license restrictions, quality gates и обязательные ADR. Согласование ROADMAP/PROJECT_VISION — первая документационная партия реализации, а не молчаливое переосмысление старых отчётов.

## 1. Подтверждённое состояние и проблемы

Источники аудита: AGENTS.md; docs/ROADMAP.md, PROJECT_VISION.md, CURRICULUM_RULES.md, SKILL_EVIDENCE.md, QUALITY_GATES.md, DEBUGGING.md; curriculum/skills.yaml, route.yaml; sources/registry.yaml; Exercism candidates/reviews и seed-plan foundation-150; milestone-150 reports; internal/learner/{route,learning,views,ui,cli}.go, ui/app.js; несколько task packages с TASK/public/grader metadata.

| Показатель snapshot | Значение и единица |
|---|---|
| Учебный маршрут | 150 различных ready Task Contract v2 task IDs |
| Activity | 75 implementation, 40 test-writing, 35 debugging |
| Происхождение | 84 substantive adaptations, 61 originals, 5 concept seeds |
| Источники 84 адаптаций | 55 наших задач из 24 исходных Exercism concept exercises, 13 Ardan, 16 LGWT |
| Все lifecycle ready | 153; три legacy вне learner route не добавлять к 150 |
| Другие состояния | 8 drafts, 1 rejected |
| Модель навыков | go-fundamentals, 28 domains, 142 active assignable skills |
| Assessed coverage маршрута | 99 skills; ещё 2 preparation, не assessment; 43 без assessed встречи |
| Версия | go.mod и config/toolchain.yaml: Go 1.27, CI 1.27.x; три legacy task modules 1.24 |
| Human completion/mastery | Неизвестны; automation fixtures не подставлять вместо ученика |

Номера разделов 01–09 и размеры: 6/11/30/13/11/47/19/4/9. Разделы 06 и 07 перегружены разными механизмами; повторения по слайсам оказались внутри других глав. Это довод за тематический индекс, а не за немедленное массовое перемещение пакетов.

Новое ещё не реализовано в snapshot: теория/подтемы в UI, algorithmic badge, пересмотр 46 отказов, запрет новых LGWT bindings. ROADMAP продолжает исключать теорию. UI Markdown renderer не делает полноценные links/images, поэтому даже link-first теория требует небольшого безопасного расширения renderer.

Текущий route.go совмещает адресацию и последовательность: возрастающие номера и порядок используются при выборе урока. Для первой теории это не препятствие: attachment по stable topic/task ID не меняет sequence. Для последующей вставки prerequisites/alternatives отделить display number от pedagogical order; не перенумеровывать готовые уроки.

Текущий learning.go хранит completion evidence отдельно, но два независимых practice/transfer ID подтверждают skill без достаточной проверки интервала и различия контрактов. Due reviews основаны на числе других completions и прекращаются после подтверждения. Это первоначальная механика, не гарантия долговременного освоения. Отдельный последующий ADR должен добавить retention evidence, не переписывая историю задним числом.

Конкретный локальный дефект: bounded-slice-view-001 обещает nil input в обычном TASK contract; public/public_test.go не содержит nil-case, grader/grader_test.go содержит. Воспроизвести defect через public Service.Check, добавить понятную public regression и проверить compatibility существующего workspace. Не отменять из-за этого весь milestone.

Примеры уже достигнутой глубины: error-detail-tests-001 проверяет errors.As по wrapped/joined errors, custom As, pointer identity и typed nil; cancel-worker-001 использует ownership/cancellation semantics, synctest и race-required profile. Эти механизмы нельзя объявлять полностью новыми. Нужна последовательная подготовка и transfer, не повторный каталог одинаковых операций.

## 2. База теории и практика работы с источниками

### 2.1. Основная теория — короткий собственный текст + точные официальные ссылки

Начальное чтение: https://go.dev/tour/ . Нормативная опора: https://go.dev/ref/spec , https://pkg.go.dev/std , https://go.dev/ref/mem . Дополнительные механизмы: https://go.dev/doc/tutorial/ , https://go.dev/blog/slices-intro , https://go.dev/blog/slices , https://go.dev/blog/pipelines , https://go.dev/blog/go1.13-errors , https://go.dev/blog/testing-time . Актуальность проверять по https://go.dev/doc/go1.27 и локальной документации закреплённого toolchain.

В topic metadata сохранять официальный URL/anchor, коротко «что здесь прочитать», проверенную Go-версию и дату review. Ссылки на пакет лучше привязывать к версии toolchain или указывать, что внешняя страница плавающая. Не имитировать exact source commit для live website.

На ранних темах Tour даёт ориентир; спецификация — справочник точного правила, не обязательная книга от начала до конца. Memory Model вводить через конкретный happens-before argument после простых операций, а не давать новичку вместо объяснения goroutine.

Собственный текст нужен для связи между prerequisite и заданием, описания механики, типичных ошибок и границ контракта. Общая теория не должна пересказывать reference solution конкретной задачи. Пример меняет API/контекст и не выдаёт алгоритм upcoming transfer.

Внешние страницы могут быть недоступны offline; краткое локальное объяснение и задания работают без сети. UI не подгружает remote HTML, scripts, изображения или iframes и не использует автоматический scraping.

### 2.2. Разрешённый перенос — исключение, не условие готовности теории

Если существенно лучше адаптировать небольшой готовый фрагмент, reviewer проверяет конкретный source artifact и действующий registry scope. Сохранить commit/path/hash, license и attribution; не достаточно ссылки на GitHub. Для links-only страницы не требуется импортировать чужой текст или обновлять registry на весь сайт.

У go.dev текст обычно CC BY 4.0, code BSD с исключениями: https://go.dev/copyright . Root BSD лицензия website не автоматически лицензирует каждый опубликованный asset. Tour scope в registry ограничен allowlist. Иллюстрации не импортировать по умолчанию.

Exercism concept Markdown может служить точечным донором начальной теории после artifact review; не копировать автоматически формулировки вроде безусловной эквивалентности nil/empty. Go by Example — разрешённые аннотированные примеры, не банк с готовыми tests. Go 101 / inancgumus остаются в текущих concepts-only/reference-only modes; ссылаться на оригиналы, не обходить ограничения переводом или соседним fork.

Текущие pinned exercise sources сохранять:

- Exercism Go `ae1477abfeeb9fbeaf971caee220e91f9e4df790`, MIT; actual source count 177 = 29 concept + 148 practice.
- Ardan gotraining `4ae5a46bfc665909a7cb29e11c18775330c63401`, Apache-2.0; selected topics/go artifacts, не коммерческий курс/внешние материалы.
- Go by Example `7d705626375ba0263b616865a286e1587d6989c8`, CC BY 3.0; examples/comments, не automatic assets/toolchain import.
- golang website Tour `c4a9d59f9775d994f1700d18fa37414c3c85fa7b`; только текущий allowlist.
- golang/go `2c31595783c74564c1299a5401dc6e606f379cc7`; только selected first-party files по registry, BSD/PATENTS и nested lineage.

Новые четыре донора и их exact commits описаны в research appendix. Их наличие в плане ещё не меняет sources/registry.yaml. LGWT source record нельзя просто блокировать целиком: historical bindings должны продолжать валидироваться, новые selection/release bindings запрещаются отдельным совместимым правилом.

## 3. Иерархия курса и глубина

Иерархия: раздел → подтема → проверяемые outcomes/skills → короткое объяснение → выбранные упражнения → самостоятельный компонент → возврат в другом контексте.

Сохраняем старые номера 01–09. Карта A–L в приложении — 12 тематических групп и 47 подтем для педагогического проектирования, не обязательная новая нумерация UI. Первый UI может показать подтемы внутри старых разделов и сквозные ссылки на упражнения из других разделов. Исторический адрес 06.42 по-прежнему открывает error-detail-tests-001.

Рекомендуемые тематические группы: A первые шаги; B коллекции; C строки/Unicode; D функции/defer; E structs/pointers/methods/embedding; F interfaces/errors; G packages/generics/iteration; H I/O/encoding/time; I testing/fuzzing/benchmarks; J goroutines/shared state; K channels/pipelines; L context/integrated work. Граф зависит от prerequisites, а не от буквенного порядка: например, io.Pipe и parallel tests изучаются после нужных concurrency-механик.

Для core-подтемы должны быть различимые evidence:

| Уровень | Что проверяется | Что не засчитывается вместо него |
|---|---|---|
| Синтаксис | Читает и собирает конструкцию | Только просмотр теории |
| Семантика | Предсказывает наблюдаемое поведение и объясняет границу | Запоминание одного готового вывода |
| Самостоятельное применение | Пишет решение по контракту с уменьшенным scaffold | Вставляет недостающий символ |
| Тесты и диагностика | Находит реальные defects и пишет meaningful assertions | Coverage percentage или наличие test files |
| Компонент | Сочетает навыки и явно владеет ресурсами/ошибками | Автоматическое начисление всем prerequisites |
| Возврат | Решает другой допустимый контекст после интервала | Reset и повтор того же task ID как новое независимое доказательство |

Подробные примеры: слайсы — view/copy/aliasing и тесты владения, затем компонент; интерфейсы/ошибки — consumer contract, несколько реализаций, typed nil и ошибка как recovery API, затем интеграция; concurrency/context — start/stop/join/cancel у каждой операции, а не просто запуск goroutine. Appendix фиксирует критерии достаточности по подтемам.

Необязательные runtime/compiler детали отделить от гарантированного языка. Capacity growth factor, размещение stack/heap, nanoseconds и точное количество allocations не являются универсальными языковыми guarantees. Семантику loop variables, generic methods, iterators и timers сверять с Go 1.27, не со старой статьёй. Старые JSON v1 задания не переписывать молча на v2. Новые version-specific outcomes вводятся через taxonomy review/ADR, а не фиктивное объявление покрытия существующим ID.

## 4. Покрытие и план роста банка

142 — модель курса, не официальный сертификат полного языка. Текущие minima 409 primary/693 total отражают принятую policy; не снижать их только чтобы проще получить зелёный отчёт, но и не генерировать бессодержательные встречи. Вопрос пересмотра minima требует отдельного педагогического обоснования и ADR, если влияет на acceptance.

43 навыка без assessed coverage приведены в приложении. Сначала закрывать toolchain/package boundaries, nil/empty, comparability, method values/field promotion, time/pipe lifecycle, test design, fuzzing/benchmark methodology и развитые concurrency/context contracts. Затем укреплять навыки, представленные почти только introduction. Existing tests, использующие -race, не автоматически учат race diagnosis.

Ориентир полного банка остаётся около 450 задач, допустимые альтернативы/повторения не обязаны проходиться подряд. Длину core определять по покрытию и достаточности, не заранее обязательными 450 уроками. Распределение после первого source review записывать как allocation по skill/outcome/role/activity с reason, а не глобальный процент заимствований.

### Источник → задача: проверяемая матрица решения

| Область / skill | Теория | Существующая опора | Внешний отбор | Original только при подтверждённом дефиците |
|---|---|---|---|---|
| toolchain / package boundaries | module tutorial, go help/doc | preparation, package-inventory-001 | official selected tutorials, Ardan | offline modules/imports/build diagnosis при несовместимости donor с isolation |
| nil/empty, aliasing, индексирование | Tour + spec slice expressions | slice-window, bounded-slice-view, detached-append | Birsi collections, Golings slices/arrays, Exercism binary-search/sublist | explicit representation/ownership contract после review источников |
| methods/embedding | Tour methods + spec method sets/selectors | receiver-speakers, embedded-cache-tests | Golings structs2; Ardan templates | ambiguity/captured receiver, если доноры ограничены синтаксисом |
| текст/Unicode | strings/runes blog, unicode/utf8 docs | 11 Unicode lessons, replace-byte | SHAD utf8/varfmt, Exercism pangram/isbn | malformed decoding boundary только если не покрыт существующей задачей |
| интерфейсный компонент | Tour interfaces, container/list, errors docs | shape-contract-tests, error-detail-tests | SHAD lrucache; Birsi errors | новый consumer-side/error boundary при отсутствии distinct donor |
| I/O/lifecycle | io/bufio docs | bounded-copy, scan-lines-tests, filesystem-record-tests | SHAD externalsort | owned adapters/cleanup intersection, если source отсутствует |
| test-writing | testing docs, TableDrivenTests | diagnostic-helper, named-cases | SHAD tabletest/tparallel | independent oracle/mutation gaps, не «тесты потому что нужны тесты» |
| generics/iteration | tutorial, spec, iter/slices/maps docs | generic-find/reduce/stack | SHAD genericsum/treeiter; Birsi generics | modern iterator stop/ownership или Go 1.27 outcome после version audit |
| cancellation/concurrency | Memory Model, sync/context, pipelines | current 9 concurrency lessons | SHAD dupcall/keylock/pubsub; Birsi channels; loong | недостающие lifetime/error/partial-result intersections после review |
| fuzzing / benchmarks | official fuzzing tutorial, testing docs | assessed tasks отсутствуют | SHAD allocs/utf8 для experiments; official selected examples | deterministic oracle/reproduction и measurement design, если нет готовой задачи |

## 5. Теория в UI: минимальный контракт

### Содержание страницы

English по умолчанию. Краткая цель, prerequisites, 3–6 основных тезисов связным текстом, один небольшой пример (обычно 10–25 строк), распространённые ошибки, переход к 1–3 упражнениям и 1–3 точные official links с пояснением, что искать. Объём — столько, сколько нужно для первого действия; ориентир 300–700 слов, не формальный минимум и не hard schema quota. Нельзя раздувать лёгкий механизм до 700 слов ради шаблона.

Глубокое чтение сворачиваемо. Схема/ASCII-представление допустимы, когда объясняют ownership или lifecycle. Анимации, imported images и расширение CSP для remote assets не нужны в первой версии.

UI показывает обзор раздела, список подтем и кнопку «Theory» рядом с задачей; после public failure можно открыть релевантное объяснение. Generic theory не раскрывает task-specific answer и не расходует hint. Подсказки и разбор решения остаются отдельными действиями. Чтение не требует завершить страницу и не блокирует уже начатую работу.

### Предлагаемые файлы и компоненты

- `curriculum/topics.yaml`: versioned metadata подтем — stable id/title, existing section associations, prerequisite topic IDs, theory path, task ID references и verified reading links. Не дублировать primary/secondary task skills как независимую истину; получать их из manifests.
- `theory/en/<topic-id>.md`: собственные локальные страницы. Schema/loader допускают постепенное наполнение; отсутствующая страница отображается как planned, а не ломает весь текущий курс.
- Небольшой `schemas/topic.schema.json` и semantic validation после ADR. Это новый ограниченный контракт, не переписывание task schema.
- `internal/learner/theory.go`, tests: безопасная загрузка repository-relative regular files с лимитом размера и existing safeio/OpenRoot подходом; не отдавать произвольные FS paths.
- `internal/learner/views.go`, `ui.go`, `cli.go`, `ui/app.js`, `ui/index.html`, `ui/app.css`: read-only TheoryView и `learn theory <topic-id>`; представление тех же данных, без второго scheduler.
- Ссылки строить из проверенного metadata или узкого Markdown subset с URL allowlist. Только безопасные http/https чтения, никакого javascript/data URLs/raw HTML/innerHTML/eval. Loopback auth/origin controls сохранить. Внешние ссылки открывать без передачи session token и с безопасными rel/referrer настройками.

Просмотр TheoryView/CLI не меняет TaskProgress, Completions, CurrentTaskID, assistance, attempts и clocks. Не делать migration progress ради этой функции. Автоматическая пауза/новое time tracking не добавляются: явные pause/resume остаются прежними.

### Позднейшие route/evidence изменения

Отдельная партия должна выбрать единственный источник pedagogical order, не оставлять два конкурирующих маршрута. Минимум: current route сохраняет display addresses; versioned ordering/choice groups определяют последовательность, ready alternatives и reinforcement. Пока это не нужно, read-only topics используют старую sequence.

Progress migration при retention поддерживает append-only события, versioned interpretation policy, old completion preservation и conservative legacy evidence. Не помечать старые два успеха как новое spaced mastery задним числом. Повтор той же задачи может быть repetition, но не вторым distinct transfer. Параметры интервала — настройка курса с обоснованием, не научный факт от модели.

До записи миграций: ADR, dry-run, backup, atomic writes/recovery, repeated/interrupted runs, collision handling, byte-preservation learner solutions, UI/CLI concurrency fixtures. Read-only команды никогда не мигрируют данные. Сначала temporary learner homes, не реальный домашний каталог пользователя.

## 6. Пересмотр кандидатов и приоритет адаптации

Старые 36 algorithmic, 6 broad, 4 duplicate решения пересмотрены в приложении. До фактического обновления reviews они остаются историческими; их не стирать. Добавить versioned rationale: old decision → new pedagogical disposition → remaining license/lineage gate → candidate status.

Педагогический reconsideration не снимает `exercism-problem-specifications: reference-only`. Не копировать canonical statements/generated cases через Exercism track. Native concept artifacts и permissive source exercises использовать только в рамках scope и с chain-of-custody. Нельзя близкий пересказ запрещённого material назвать original.

Дублирование проверяется сравнением контрактов: множество всех букв (pangram) не то же, что отсутствие повторов (isogram); ISBN grammar/checksum не автоматически повтор generic validation/Luhn. Обратная крайность тоже запрещена: новый task ID или другой source repo не делают копию counter/shape/filter новой самостоятельной практикой.

Большая работа может оценивать один primary outcome и сочетать несколько изученных secondary навыков. Не требовать искусственно расколоть её на десятки accessor-задач. Prerequisites алгоритма/формата объявлять до работы; дополнительная badge не меняет тип submission. Готовый алгоритм в theory не должен превращать transfer в перепечатку.

## 7. Этапы и Definition of Done

| Этап | Продуктовый результат | DoD |
|---|---|---|
| M0. База решений | актуальная документация, baseline, новые источники в review | inventory присоединён к actual manifests; historical LGWT валиден; no new copied content; conflicts отмечены |
| M1. Theory MVP | читабельные локальные страницы и official links | одна вертикальная цепочка Theory → GoLand → Check → return to theory; старые 150 уроков доступны; state неизменен от чтения |
| M2. Первое содержательное расширение | небольшой batch и source-first allocation следующего | 3–6 полезных задач или меньшая доказанная партия; public/full/reference/mutants/review; no quota filler; IDs/history сохранены |
| M3. Полнота и глубина | отсутствующие skills, независимые компоненты, тестирование/fuzz/bench | каждый необходимый profile готов до content; coverage и semantic review, без fake encounters; route prerequisites не сломаны |
| M4. Retention и foundation acceptance | выбор alternatives, возврат после интервала, итоговый evidence report | migration fixtures зелёные; full -coverage; никакое human mastery не выводится из готовности банка |

После трёх content batches делать audit семантических повторов, типов работы, ошибочных вариантов и public feedback. Обратная связь от реальных completion используется, когда есть; её отсутствие не блокирует preparation и не подтверждает освоение.

Не требуется новый большой раунд продуктового выбора. Допустимы небольшие ADR для конкретного контракта и bounded source reviews, перечисленные здесь. Если donor не прошёл gate, recorded reject/defer и другой независимый approved candidate; не менять цель и не выполнять mass generation.

## 8. Первая исполнимая волна

Партии ниже — отдельные PR. Их обозначения P0...P7 не заменяют существующий batch ledger: при реализации назначить непротиворечивые номера, сохранив связку с этим планом.

### P0 — documentation/ADR, без content

Читать baseline и AGENTS; сверить master/dirty worktree, не перезаписывать чужие изменения. Обновить docs/ROADMAP.md, PROJECT_VISION.md, CURRICULUM_RULES.md, SKILL_EVIDENCE.md и LEARNER_CLI.md только в затронутых положениях. Зафиксировать small Theory metadata ADR, алгоритмичность как facet, historical LGWT freeze policy. Привязать milestone-150 как исторический результат, не пересчитать его под новую цель.

Результат: один authoritative active-plan pointer, baseline audit, bounded schema proposals без перемещения задач. Большой ADR по mastery или route reorder не является зависимостью Theory MVP.

### P1 — read-only Theory tooling

Зависит от P0. Реализовать loader/validation/views/CLI/UI в компонентах раздела 5 с synthetic fixtures, без реальных учебных страниц в этой tooling-партии. Принять старый route v2 и progress v3 без изменения байтов. Renderer links: valid https/anchors работают, raw HTML/javascript/path traversal/symlinks/oversize/unknown IDs отклоняются. Сеть не нужна для загрузки страниц. Empty/missing theory не ломает обычные 150 задач. UI/CLI открывают один и тот же topic.

Acceptance: targeted loader/renderer/no-side-effect tests, затем обязательная tooling acceptance. Зафиксировать результат, не объявлять теорию поставленной до P2.

### P2 — шесть коротких страниц theory-content

Зависит от P1. Страницы: A1 рабочее место; A5 первые tests; B2 slice view/bounds; B3 append/copy/ownership; F1 interface contracts; L1 cancellation/lifecycle. Последние две доступны в соответствующей части курса, не навязываются новичку на старте.

Каждая: собственный English текст, проверенный маленький пример, 1–3 official links, prerequisites и привязка к реальным task IDs. Для B2 не использовать CappedView как практически готовое решение bounded-slice-view. Для L1 не повторять целиком StartCancelable. Для A1 основной GoLand workflow и различие public/full; не ставить новый редактор.

Проверки: content/style/links metadata, компиляция и тесты собственных examples в disposable environment, no solution leakage, no hint/completion side effects, independent review. Broken remote link не должен ломать offline execution. Deferred links можно заменить точным official anchor после review, не выдумывать его.

### P3 — source review, без новых ready задач

Может идти параллельно P1/P2 после P0. Зафиксировать exact commits research appendix через существующий source-sync/allowlist процесс. Начать с 24 listed source units и 5 Koans files; обновить 46 старых Exercism решений. Полный inventory может быть крупнее десяти, но content release limit сохраняется.

Для первоочередного bounded artifact review выбрать SHAD tabletest/lrucache/externalsort, Golings structs2 и Birsi collections; дополнительно проверить первый algorithmic Exercism candidate. Это очередь проверки, не требование использовать всех пятерых. Record: exact files/hashes/license lineage/behavior/outcomes/actual prerequisites/dedup comparison/public defect ideas/adaptation burden/decision. При необходимости generic candidate schema вводится отдельной tooling-партией; не обходить существующие validators произвольным несовместимым YAML.

Результат: accepted-for-adaptation candidates и честный reject/defer, не ready counts. Писать originals только после зафиксированного отсутствия подходящего source behavior. Future content allocation должно ссылаться на этот результат, не на рекламное число задач.

### P4T + P4C — локальное исправление public feedback

Две связанные небольшие партии по существующему process: tooling regression через learner Service.Check и content amendment bounded-slice-view-001. Проверить nil-case с полезной диагностикой, корректный reference и mutant, который падает именно public path. Для изменившегося digest использовать existing content revision compatibility, не reset пользователя.

Затронутые места: tasks/ready/bounded-slice-view-001/{tests/public,review,provenance/lifecycle по актуальному contract}, internal/learner/public_feedback_integration_test.go и при необходимости existing compatibility mechanism. Не создавать новый task ID вместо исправления. Final gates после стабильного diff, один collection на неизменные inputs.

### P5 — первый новый content batch, максимум шесть задач

Зависит от P3 и поддержки нужного profile; theory references готовы либо короткая соответствующая страница поставлена отдельной партией. Не ждать mastery migration. Task IDs назначить после collision check; номера только добавить, не переназначать. Не вставлять задачу до prerequisites ради красивой темы; до route-order ADR использовать подходящий конец подготовленного раздела либо оставить verified package вне обязательной последовательности с честным статусом.

| Slot | Кандидат / задача | Activity и смысл | Gate до авторинга |
|---|---|---|---|
| W1 | SHAD tabletest → тестирование parsing contract | test-writing: классы valid/malformed/overflow/sign/units по выбранному source contract; mutation-check, не >95% coverage | проверить происхождение ParseDuration и prerequisite time/parser; не копировать ненаблюдаемые требования |
| W2 | Golings structs2 → поле через embedding | implementation или debugging: promoted versus explicit selector, independent state | прочитать actual file; если только trivial syntax без полезной семантики, отклонить, не раздувать artificial adaptation |
| W3 | Exercism binary-search | implementation/practice + algorithmic badge: boundaries и invariant | разрешённые artifact sources; generated cases не брать; duplicate policy и Go prerequisites |
| W4 | Exercism pangram | test-writing/practice + algorithmic badge: completeness множества, не uniqueness | explicit alphabet/normalization contract, permitted lineage, meaningful mutants |
| W5 | Nil/empty collection contract | test-writing: различить representation и длину там, где контракт это требует | сначала review Birsi collections/существующих задач; original разрешён только если donor не закрывает distinct outcome |
| W6 | SHAD lrucache | implementation/transfer: целостный bounded container, recency/eviction | source API/test review; maps/methods/interfaces готовы; public empty/capacity/update/eviction boundaries |

Это верхняя граница и приоритетная очередь, не квота шесть любой ценой. Не принятый W-slot не заменять близким пересказом reference-only. Возможен выпуск меньшей партии. W6 не дробить на десятки методов ради статистики. SHAD externalsort, treeiter, dupcall и Birsi channels — следующая волна, не скрытая дополнительная работа того же batch.

Каждый slot имеет TASK с наблюдаемым API, 4–7 requirements, starter/public/full/reference, realistic mutants или weak learner tests, mapping требования→public assertion→defect→diagnostic, independent review, attribution и exact lineage. Use task-promote, не lifecycle mv. Не считать теорию новой задачей банка.

### P6 — отчёт и следующая allocation

После P2/P3/P5 показать пользователю работающий theory loop и фактическое изменение counts. Отдельно: ready/new substantive adaptations/originals/concept seeds/source units/assessed skills/глубина по ролям/человеческие evidence неизвестны. Проверить отсутствие переименования старых IDs и числовых адресов. Сохранить prepared next block и план следующего ≤10 batch.

### P7 и далее — отдельные tooling зависимости

Только когда реально нужны: versioned route choice/order migration; retention evidence; fuzzing/benchmark profiles. Не блокировать ими шесть страниц теории. Для fuzzing: deterministic seeds/oracles и воспроизводимый counterexample; bounded randomized run не единственный grader. Для benchmarks: корректные timing boundaries и representative workloads, semantically equivalent alternatives, никаких общих fixed ns/op thresholds. Profile schema changes ADR до content.

## 9. Неизменные acceptance и рабочий процесс

Каждая партия — отдельная codex/* ветка и PR. Никакого push в master, merge только после независимого review и зелёного exact-head validate. Существующие protection/process gates и isolation не ослаблять. Этот planning-only PR не заявляет запуска этих команд.

Обязательный порядок для tooling и принимаемых content изменений согласно AGENTS/QUALITY_GATES:

```sh
go test ./...
go test -race ./...
go vet ./...
go test -tags=taskintegration ./...
go run ./cmd/task-check -root . -all -output .tmp/results/task-checks.yaml
go run ./cmd/validate -root . -phase1 -checks .tmp/results/task-checks.yaml
go run ./cmd/task-report -root . -check -checks .tmp/results/task-checks.yaml
```

Это существующий режим infrastructure/неполного банка из QUALITY_GATES, не разрешение ослабить quality конкретных задач. Требование release `-coverage` из AGENTS сохраняется; итог foundation обязательно проходит `-coverage`. Если в текущих документах есть неоднозначность применения полного coverage к промежуточному content batch, согласовать её документационным ADR/существующей принятой phase policy до выпуска, не самовольно обойти gate.

Collection создаётся один раз для неизменных inputs и повторно используется validator/report. Authored task code/tests никогда не запускаются непосредственно из canonical source checkout; repository test harness запускает disposable copies. При исправлении public feedback нужен именно public-only regression, а не один full failure. Timeout/compile error не доказывает требуемую диагностику.

Остановка допустима при реальном unresolved license/contract/safety blocker; сначала завершить независимые разрешённые партии. Отсутствие нового сообщения пользователя не блокирует последовательную работу в пределах принятого плана. Не выдавать partial inventory, planned profiles или fixture completions за delivered course/human mastery.
