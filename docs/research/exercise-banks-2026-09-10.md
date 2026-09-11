# Дополнительные банки упражнений Go: отбор для foundation

Дата исследования: 2026-09-10. Статус: исследование и очередь review, не разрешение на выпуск задач.

Основной исполнимый план: [foundation-next.md](../plans/foundation-next.md). Педагогическая карта: [foundation-topic-map.md](../plans/foundation-topic-map.md).

## 1. Вывод и границы проверки

Рекомендуем расширить существующие Exercism / Ardan Labs четырьмя источниками: прежде всего `slon/shad-go` для самостоятельных компонентов, затем `mauricioabreu/golings` для выборочной начальной практики, `birsi/go-training` для компактных пакетов с тестами и reference, `cdarwin/go-koans` для отдельных упражнений на семантику. Последний имеет низший приоритет из-за старых предположений и формата заполнения пропусков.

Поиск включал готовые курсы с тестами, Rustlings-подобные тренажёры, concurrency workshops, алгоритмические банки и небольшие проекты. Ни число звёзд, ни название «100 exercises», ни наличие LICENSE не считаются подтверждением качества или происхождения каждого файла.

В этой сессии читались публичные страницы, деревья, manifests, лицензии и перечисленные ниже отдельные файлы. Код источников не запускался. Полный независимый педагогический, лицензионный и adversarial review каждого кандидата ещё не выполнен. Ничего из перечисленного не является новой ready-задачей Go Learn.

Единицы считать отдельно:

- SHAD: исходный именованный package; многочастный `externalsort` пока одна исходная единица.
- Golings: запись `[[exercises]]` в `info.toml`, не файл тестов и не hint.
- Birsi: учебный каталог; зеркальный каталог решения не второе упражнение.
- Go Koans: тематический файл содержит несколько проверок; число файлов не равно числу упражнений.
- Gophercises: номер мини-проекта; ветки p1...p8 являются стадиями одного проекта.

Конкретная очередь ниже содержит 12 SHAD packages + 8 Golings entries + 4 Birsi packages = 24 исходные единицы на review. Отдельно назначены 5 тематических файлов Go Koans. Это не обещание 29 готовых задач и не прогноз доли заимствований. Разделение многочастного упражнения допускается только по самостоятельным контрактам с сохранением общего upstream identity.

## 2. Основные источники

### SHAD: slon/shad-go — главный новый донор компонентов

- Репозиторий: https://github.com/slon/shad-go
- Проверенный main: `6718601587ec77cfaa632bbd29a8ca0716fe81ec`.
- Пин: https://github.com/slon/shad-go/tree/6718601587ec77cfaa632bbd29a8ca0716fe81ec
- Лицензия: MIT, Copyright (c) 2023 Fedor Korotkiy. Доказательство: https://github.com/slon/shad-go/blob/6718601587ec77cfaa632bbd29a8ca0716fe81ec/LICENSE ; Git blob `4009ee95a6a858019852623a7b1c7d37bf7ab899`.
- `go.mod`: Go 1.24.0; есть многочисленные внешние зависимости. Не переносить модуль целиком. Проверять адаптацию на закреплённом Go Learn Go 1.27 и только разрешённом stdlib-профиле.
- Аудитория: знающий основы программирования ученик; многие задания intermediate/advanced, не последовательное введение с нуля.
- Формат: русские README, starter Go, tests, examples; программа в `docs/syllabus.md`. Публичная полнота reference solutions не подтверждена. Свои reference/mutants всё равно обязательны.
- Лицензия root применима к собственному коду и associated documentation, но не отменяет стороннее происхождение конкретных скопированных файлов. Иллюстрации, внешние slides, testdata и заимствованные фрагменты отдельно проверять; автоматически не импортировать.
- Объём: в `.manytask.yml` 62 активных записей расписания, включая bonus и backend/low-level; закомментированные записи не учитывать. Это не 62 подходящие foundation-задачи. Например, `allocs/` существует в дереве, но не входит в это расписание.
- Изучены: root README, syllabus, `.manytask.yml`, LICENSE, go.mod, дерево; README `lrucache`, `dupcall`, `tabletest`, `externalsort`, `treeiter`, `utf8`; первые 180 строк `dupcall/dupcall_test.go`.

Содержательные находки:

`lrucache` — ограниченный контейнер с recency-политикой и интерфейсом. Хорошая самостоятельная работа по maps, методам и интерфейсам без HTTP/DB. Проверять eviction, обновление существующего ключа, отсутствие/нулевое значение, ёмкость и повторное использование. Не называть его concurrency-задачей без соответствующего исходного контракта.

`dupcall` — дедупликация одновременно выполняемых callback с независимой отменой ожидающих клиентов. Отмена одного клиента не отменяет общую работу, пока остаются другие. Это полноценный поздний capstone по lifecycle/context, не очередной counter. В прочитанных тестах есть `time.Sleep`, `time.After`, testify, goleak и testtool: их нельзя переносить как готовый детерминированный grader. Перестроить проверку на управляемые события и существующий synctest/race harness.

`externalsort` — Reader/Writer строк, merge, затем файловая композиция. Полезный источник для I/O и алгоритмического компонента. Контракт допускает очень длинную строку и последнюю строку без newline. Зафиксировать EOF/partial-result semantics публично, устранить неоднозначность имён в upstream README. Последняя стадия допускает перезапись входных файлов: в Go Learn разрешать это только в явно owned disposable fixtures, не в произвольных путях пользователя. Адаптация только merge не должна заявлять оценку всей внешней сортировки.

`tabletest` — написание table-driven tests для ParseDuration, исходный критерий >95% coverage. Нам нужен отбор граничных классов и обнаружение реалистичных ошибок, а не процент как единственный зачёт. До копирования реализации отдельно проверить её происхождение: название функции само по себе не доказывает независимость от stdlib. Если файл заимствован, сохранять дополнительную source chain.

`treeiter` — обход бинарного дерева с callback, сигнатуру ученик выводит из наблюдаемого контракта. Не выдавать это автоматически за современный `iter.Seq`: другой протокол требует отдельной адаптации и нового outcome.

`utf8` — несколько строковых подзадач с benchmark-сравнением и allocation-ограничениями. Не выпускать весь пакет как раннюю Unicode introduction. Сначала выбрать конкретный функциональный контракт; performance вынести в отдельную работу после benchmark tooling, не оценивать скорость относительно авторского ноутбука.

Предлагаемый registry ID: `shad-go`. Допуск после bounded artifact review; allowlist только выбранных packages. MIT notice сохранить. Existing `tour0`, `tour1` и `forth` требуют проверки общей lineage с Tour/Exercism; другой репозиторий не делает одно упражнение независимым новым источником.

### Golings: mauricioabreu/golings — выборочная начальная практика

- Репозиторий: https://github.com/mauricioabreu/golings
- Проверенный main: `65fb8ce827c716e54650e88a3ccc7c52308ab59b`.
- Inventory: https://github.com/mauricioabreu/golings/blob/65fb8ce827c716e54650e88a3ccc7c52308ab59b/info.toml
- License: Apache-2.0, https://github.com/mauricioabreu/golings/blob/65fb8ce827c716e54650e88a3ccc7c52308ab59b/LICENSE ; blob `261eeb9e9f8b2b4b0d119366dda99c6fd7d35c64`.
- Пересчёт inventory: 43 записи = 34 `compile` + 9 `test`. Это проверка metadata, не прогон всех решений.
- Аудитория: новичок, короткие механики declarations, control flow, arrays/slices/maps, closures, generics, первые concurrency-ошибки.
- Формат: Go-файлы, comments, hints в TOML; часть заданий в `_test.go`. Отдельного полноценного учебника и большого массива independent tests не установлено.
- Версия: совместимость с Go Learn 1.27 проверить для каждого выбранного файла; новый runner не устанавливать и root toolchain не переносить.
- Изучены root страницы, LICENSE, дерево, весь `info.toml`, файл `exercises/anonymous_functions/anonymous_functions3/main.go`.

Проверенный пример anonymous_functions3 действительно сохраняет индекс в замыкании. Но это scaffold с пропуском и panic в main, а не наш public/full контракт. В Go Learn нужны компилирующийся starter, наблюдаемые переходы состояния, независимые проверки и разумный terminal-state contract. Сначала сравнить со `stateful-counter-tests-001` и `period-predicate-001`: заменённые слова не создают новую практику.

Apache license/notices и пометки изменений сохранить; TUI, watcher, install scripts, branding не импортировать. `compile-only` не равен готовому debugging-профилю Go Learn: по текущим gates starter обязан компилироваться. Не менять ради массового импорта глобальное правило проверки. Малополезные пропуски оставить вне банка либо использовать как собственные демонстрации без начисления покрытия.

Предлагаемый registry ID: `golings`. Рассматривать как источник дополнительных вариантов и нескольких точечных пробелов, не основу ещё 300 задач.

### Birsi: go-training — небольшой банк с решениями и тестами

- Репозиторий: https://github.com/birsi/go-training
- Проверенный main: `3ae230e734644d637118f2a6818166922d5197c5`.
- License MIT, Copyright (c) 2026 Michael Birsak: https://github.com/birsi/go-training/blob/3ae230e734644d637118f2a6818166922d5197c5/LICENSE ; blob `b289838ebf911f31c4ed5bd535eb7978990e4a55`.
- `go.mod` проверен: Go 1.22, без require. Адаптации и examples всё равно проверить под Go Learn 1.27.
- README перечисляет 12 учебных каталогов в трёх стадиях; reference находится в зеркальном `solutions/`. Это 12 учебных пакетов по заявленной структуре, не 24, и не независимый пересчёт каждого TODO.
- Темы: базовые значения, функции, коллекции/pointers; structs/methods/interfaces/errors; generics/goroutines/channels/JSON.
- Формат: English comments с задачами, tests, reference, официальное дополнительное чтение. README/код можно адаптировать в пределах MIT с сохранением notices; screenshots/branding исключить. Инструкции чужому AI-наставнику не исполнять и не переносить в наш workflow.
- Изучены README, LICENSE, go.mod, каталог `exercises/stage3/11_channels`, полностью `channels_test.go`. Остальные выбранные каталоги пока curriculum-level review.

Плюсы прочитанного `channels_test.go`: реальные signature assertions, сравнение результата Merge без предположения о порядке, FirstValue проверяется для обоих расположений ready-input, внешних зависимостей в этом файле нет.

Недостатки: успешный range может зависнуть, empty/closed/nil/early-stop контракты раскрыты не полностью; отсутствуют наши adversarial mutants и public/full separation. Не добавлять произвольные требования о cancellation в скрытый grader. В исходнике четырём функциям соответствует один пакет: разделение допустимо лишь при разных outcomes, а не автоматически 4 новых task IDs.

Предлагаемый registry ID: `birsi-go-training`. Не принимать slogan «tests pass => solved» за нашу модель освоения.

### Go Koans: cdarwin/go-koans — точечный резерв, не современный каркас курса

- Репозиторий: https://github.com/cdarwin/go-koans
- Проверенный master: `7f0b28e4f551f2120d55eec3a03b4f472ec38044`.
- License MIT, Copyright (c) 2019 Jason Thigpen: https://github.com/cdarwin/go-koans/blob/7f0b28e4f551f2120d55eec3a03b4f472ec38044/LICENSE ; blob `b4b35234acf98b0fbcec38fea30eaed9fb0a67c4`.
- На просмотренной странице дерева 20 `about_*.go` тематических файлов; каждый содержит несколько assertions. Не считать каждый assertion самостоятельным упражнением.
- Аудитория: знакомство с семантикой через наблюдения; не полноценная независимая разработка компонентов.
- Формат: Go/assert/пропуски; старый custom harness. Go 1.27 compatibility не подтверждена исполнением.
- Изучены root страница/harness, лицензия и pinned `about_slices.go`.

В `about_slices.go` проверяются конкретные значения capacity после append и предлагается угадывать алгоритм роста. Такие expectations запрещено переносить как гарантии языка. Допустимая новая работа должна различать reuse/reallocation при заданной capacity, а не угадывать runtime formula. Сюжетный и неуместный декоративный текст удалить. Любое сохранённое поведение и код честно обозначать адаптацией.

Выбранные для последующего чтения 5 файлов: `about_slices.go`, `about_interfaces.go`, `about_defer.go`, `about_channels.go`, `about_allocation.go`. Кроме slices, bounded file-level review ещё нужен. Возможный registry ID `go-koans`; не переносить assertions framework, `os.Exit`-диагностику или весь курс.

## 3. Расширенный поиск: почему другие находки не стали основными донорами

| Источник | Что найдено | Решение и предел проверки |
|---|---|---|
| https://github.com/stonecharioteer/goforgo | README заявляет 234 упражнения, 57 категорий, Go 1.24+, exercises/solutions | Main `77c68285f14eb3ae6ecc4253b7a662fc04b27476` закреплён. README ссылается на MIT LICENSE, но запрос LICENSE дал 404, полное root-tree его не содержит. Это недостаточность license evidence, а не окончательное юридическое заключение. Не копировать до уточнения; 234 — заявление автора, не наш пригодный inventory. Есть много out-of-scope тем. |
| https://gophercises.com/ | 20 нумерованных мини-проектов, хорошая самостоятельная работа | Рассмотрен deck: https://github.com/gophercises/deck/tree/299593da734404625664bf216ca8739226fbbf63 . В этой ветке README и students index, LICENSE отсутствует; p1...p8 — стадии. Не обобщать проверку на лицензии всех 20 репозиториев. Ссылки на оригинал допустимы; импорт выбранного материала пока не разрешён. |
| https://github.com/cblte/100-golang-exercises | Заявлены 100; на просмотренной странице 20 practice/reference каталогов; README сообщает перенос найденных Python-упражнений | Не считать 100 подтверждёнными задачами. Происхождение Python-условий не установлено; license root сам по себе не устраняет этот риск. Не выбран для импорта; полный lineage audit не проведён. |
| https://github.com/bregman-arie/go-exercises | 98 смешанных вопросов и exercises по описанию README | Вопросы/готовые ответы не считать 98 executable tasks. Источник просмотрен как тематический список; file-level license/quality review не выполнен. Не включать в обещанный объём адаптации. |
| https://github.com/blueprismo/100-golang-exercices | Последовательный список, включая HTTP/DB/context и testing | Внешний поиск и обзор программы, не полный corpus review. Разрешение переноса не подтверждено. Не выбран вместо проверенных банков. |
| https://github.com/plutov/practice-go | Именованные задачи, tests/benchmarks; много алгоритмов | Алгоритмичность не отказ. До разрешения импорта нужно отдельное подтверждение license/lineage и ограничений тестов; оно в этом исследовании не завершено. |
| https://github.com/danyjacob45/go-concurrency-exercises | Тематическая практика channels, WaitGroup, Mutex и других примитивов | Найден и просмотрен overview; точный corpus/license review ещё не выполнен. Резерв, не основной обязательный ресурс. |
| https://github.com/l7mp/learning-go | Последовательные главы, CC-BY-NC-SA и связь с inancgumus | Не обходить `reference-only` и NC/SA ограничения через производный репозиторий. Только внешнее чтение/проверка тем по действующей policy. |
| https://github.com/bbc/eng-golang-workshop | Workshop, submitted code Apache; упражнения связаны с книжным материалом, использованным с отдельного разрешения | Разрешение автору workshop не является нашим разрешением. Не переносить книжные условия по лицензии соседнего кода. |

Ранее выбранный `loong/go-concurrency-exercises` остаётся резервом, не новой находкой этого раунда: `39c1ab9b02c0cc21bd66ea1e31335c4bf4dd6708`, README WTFPL v2, каталог 6 challenges, Go 1.19 по exercises.json. Producer-consumer, race-in-cache и limit-service-time полезны после удаления wall-clock требований; остальное не расширяет текущий foundation scope. Источники: https://github.com/loong/go-concurrency-exercises/blob/39c1ab9b02c0cc21bd66ea1e31335c4bf4dd6708/README.md и https://github.com/loong/go-concurrency-exercises/blob/39c1ab9b02c0cc21bd66ea1e31335c4bf4dd6708/exercises.json .

## 4. Конкретная очередь из 24 исходных единиц

`R` = прочитан указанный отдельный контракт/код; `M` = выбран по manifest/дереву/программе, содержательная проверка ещё впереди. Ни R, ни M не означают acceptance. Имена A1...L4 — педагогические подтемы, не новые task IDs или перенумерация.

| № | Источник и путь на закреплённом commit | Review сейчас | Подтема, предполагаемая роль и требуемая адаптация |
|---:|---|---|---|
| 1 | shad-go `utf8/` | R README; части не разобраны | C2: выбрать одну самостоятельную text-операцию; снять performance gate для basic-версии. Сравнить с уже имеющимися 11 Unicode-уроками. |
| 2 | shad-go `varfmt/` | M | C1: проверить форматную грамматику/escaping; не делать ещё один literal formatter. |
| 3 | shad-go `lrucache/` | R README | B5/E2/F1: самостоятельный bounded container, recency/eviction tests, implementation или test-writing. |
| 4 | shad-go `externalsort/` | R README | H1/H2: line contract → merge → поздняя FS-композиция; total source count остаётся один. |
| 5 | shad-go `tabletest/` | R README | A5/I1/H4: data partitions и mutation-tests вместо quota coverage; source chain ParseDuration проверить до копирования. |
| 6 | shad-go `tparallel/` | M | I3: изоляция fixtures и parent/child scheduling; проверить, что предполагаемая механика действительно есть. |
| 7 | shad-go `genericsum/` | M | G2: named types/constraints; не брать косметический дубль generic-reduce. |
| 8 | shad-go `treeiter/` | R README | G2/G4: callback traversal; API/signature и порядок определить по examples/tests; не приписывать iter.Seq без новой адаптации. |
| 9 | shad-go `dupcall/` | R README + tests excerpt | L1/L2/L4: независимая отмена клиентов и shared callback lifetime; поздний capstone. |
| 10 | shad-go `keylock/` | M | J2/K2/L1: возможный concurrency-компонент; реальный contract/blocking policy ещё проверить. |
| 11 | shad-go `pubsub/` | M | K3/L4: возможная композиция каналов; локальная библиотека, не внешний broker. |
| 12 | shad-go `allocs/` | M дерево | I5/B3: allocation experiment только после профиля; точный semantic/performance contract ещё проверить. |
| 13 | golings `exercises/variables/variables4/main.go` | M info.toml | A2/A3: scope/shadowing; executable behavior, не исправление одного символа ради task count. |
| 14 | golings `exercises/arrays/arrays1/main.go` | M info.toml | B1: индексы/границы; проверить отличие от array-copy. |
| 15 | golings `exercises/slices/slices4/main_test.go` | M info.toml | B2: bounds, варианты empty/singleton; дополнение, не обязательный повтор старого API. |
| 16 | golings `exercises/structs/structs2/main.go` | M info.toml | E3: field promotion; compiled starter и осмысленная mutation-проверка. |
| 17 | golings `exercises/anonymous_functions/anonymous_functions3/main.go` | R код | D1: stateful closure; сначала сравнить с stateful-counter, иначе оставить remedial. |
| 18 | golings `exercises/generics/generics2/main.go` | M info.toml | G2: constraint operations; не засчитывать if solution только дописывает известную сигнатуру. |
| 19 | golings `exercises/concurrent/concurrent2/main_test.go` | M info.toml | J2/I3: remediation race counter; высокий риск дубля двух текущих counters. |
| 20 | golings `exercises/concurrent/concurrent3/main_test.go` | M info.toml | K1: send-after-close lifecycle; проверить ownership, а не подавление panic. |
| 21 | birsi `exercises/stage1/03_collections/` | M README | B3/B4/B5: review nil/empty/aliasing, не заявлять их покрытие по одному названию. |
| 22 | birsi `exercises/stage2/08_errors/` | M README | F3: возможные альтернативы, но existing Is/As уже глубокие; проверить новый контракт. |
| 23 | birsi `exercises/stage3/09_generics/` | M README | G2: constraints/zero/comparable; не размножать текущий stack/find/reduce. |
| 24 | birsi `exercises/stage3/11_channels/` | R каталог + tests | K1/K2/K3: нормальное закрытие и completeness fan-in; empty/closed/ownership и ошибки проверять явно. |

Нулевой результат отбора отдельного источника допустим. Не добавлять заведомо слабые варианты ради того, чтобы каждый исследованный банк был представлен в ready.

## 5. Общие правила допуска

1. Зафиксировать repository/commit/relative paths/license и SHA-256 реально использованных файлов, а не только root README. Upstream reference остаётся авторским материалом, не learner payload.
2. Зарегистрировать источник до адаптации; для нового source ID отдельно решить usage_mode и allowlist. Лицензии сохранять в исходной редакции; другие вложенные источники и restrictions не терять.
3. Сопоставить source behavior с текущими task IDs, не только skills. Отличия проверять по API, инварианту, требуемой самостоятельности, failure modes, representation, типу работы и роли в маршруте.
4. Проверить точную Go-версию, compile/runtime guarantees, loop capture, timers, generics, standard-library semantics. Не переносить runtime-specific cap/allocations как вечные гарантии. Нормативная опора: https://go.dev/ref/spec и https://go.dev/doc/go1.27 .
5. Сделать компилирующийся starter, открытые проверки всех обычных требований, приватное расширение входов, reference и реалистичные mutants. Для test-writing различать assertions ученика и скрытые mutants; новый тип `algorithmic` не создавать — это facet существующего activity.
6. Concurrency проверять через existing disposable task checker, channels/synctest/race; случайные sleep/скорость CI не превращать в оценку ученика. Fuzzing не должен зависеть от счастливого обнаружения ошибки за случайное время, benchmarks — от абсолютного ns/op.
7. Каждая выпускаемая партия ≤10 задач, tooling отдельно, independent review и exact-head validate. Реальное происхождение сохранить и при значительном rewrite.

Приоритет разработки originals: только после recorded review ближайших доноров либо при отсутствии подходящего проверяемого source behavior. Сильные вероятные дефициты остаются в toolchain/import boundaries, устойчивых fuzzing-oracles, benchmark methodology и в собственной адаптации повторений; наличие нового банка не делает эти пробелы автоматически закрытыми.
