# Foundation: карта подтем, глубина, покрытие и reconsideration

Приложение к [основному плану](foundation-next.md), редакция 2026-09-10. Новые источники и точные commits: [исследование банков](../research/exercise-banks-2026-09-10.md).

## Как читать карту

A–L — педагогические группы, не новые адреса вместо 01–09. A1...L4 — стабильные proposed topic IDs. Первая поставка связывает их с существующими разделами и task IDs без смены sequence или learner progress.

Prerequisites ниже относятся к полному объёму подтемы. Первое краткое знакомство может быть раньше, если реальная task metadata это разрешает. Существующие задачи в карточке — опорные примеры и поздние возвращения; перечисление не означает, что все их prerequisites уже выполнены на первой странице. Например, record-result, dictionary-add-tests и writer-greeting не перемещаются в первые lessons только потому, что используют функции/maps/interfaces. Перед каждым новым placement проверяется actual task prerequisite closure.

Роли: introduction — изолированная механика, practice — другой случай и меньше scaffold, transfer — самостоятельное сочетание. Activity implementation/test-writing/debugging независима от роли и algorithmic badge. Core — необходимая практика; дополнительная — альтернативы/повторения; углубление — необязательная часть основного прохождения. Углубление в банке всё ещё участвует в предусмотренном coverage active taxonomy, пока ADR не изменит policy.

Общий критерий достаточности каждого core-механизма: ученик объясняет наблюдаемый эффект, выполняет контракт без готового алгоритма, обнаруживает характерный defect тестом и применяет механизм в отличающемся контексте. Возврат нужен после другой работы/интервала. Это проектируемое evidence, не уже доказанный результат ученика. Ни одна перечисленная ready-задача сама по себе не подтверждает полную глубину темы.

## Теоретические опоры

Локальные English-страницы краткие и самостоятельные; ссылки раскрывают детали. Ниже сокращения для точных первоисточников:

- T: https://go.dev/tour/ ; последовательное первое знакомство, не исчерпывающая семантика.
- S: https://go.dev/ref/spec ; соответствующие разделы Declarations, Types, Expressions, Statements, Built-in functions.
- M: https://go.dev/ref/mem ; synchronization и happens-before после знакомства с primitives.
- P: https://pkg.go.dev/std ; выбирать конкретный пакет/symbol и сохранять verified Go version.
- F: https://go.dev/doc/tutorial/ ; getting-started, create-module, generics, fuzz.
- B: https://go.dev/blog/ ; slices-intro, slices, strings, maps, go1.13-errors, pipelines, testing-time.
- V: https://go.dev/doc/go1.27 ; version audit, не отдельный обязательный курс новинок.

Exercism/Ardan/Birsi/Golings/Koans в карточках — donor candidates, не автоматическое разрешение заимствования. Existing Learn Go with Tests adaptations не удалять, но новые donor bindings оттуда не создавать.

## A. Первые шаги

### A1. Рабочее место, запуск и модуль
Prerequisites: нет. Skills: toolchain.modules-workspaces/build-run/formatting-vetting/docs-environment.
Понимать package/module/repository/workspace; самостоятельно открыть learner workspace в GoLand, выбрать пакет, прочитать диагностику и найти symbol docs. Ошибки: чужая working directory, случайный go.work, изменение canonical tests, представление vet как доказательства корректности. Теория: F getting-started/create-module, go help/go doc.
Сейчас: preparation, отдельных assessed toolchain-задач нет. Donors: official selected examples/Ardan; originals только для безопасного offline module/import diagnostic, если donor требует сеть или неразрешённую FS. Виды: implementation небольшого package, debugging запуска, проверяемый diagnostic artifact. Достаточно: восстановить новый сломанный локальный запуск и объяснить причину, а не повторить команды. Core; workspace internals углубление.

### A2. Значения, типы, constants и zero values
Prerequisites: A1. Skills: types.*, conversions и numeric contracts. Понимать различие value/type, defined/alias, constant/runtime, zero/nil; самостоятельно выбрать представление и безопасное преобразование. Ошибки: widening после overflow, untyped representability, NaN equality, shadowing. Теория: T basics, S Types/Constants/Conversions.
Сейчас: literal-bindings-001, zero-results-001, wide-sum-001, named-block-001, constant-kinds-001. Donors: Exercism concept artifacts, Golings variables4, Birsi basics. Originals: лишь отсутствующие scope/overflow semantics после сравнения. Виды: implement/debug/test-writing. Достаточно: объяснить и проверить иной numeric boundary, не только литералы. Core; complex arithmetic и representation internals — углубление.

### A3. Ветвления, циклы, scope и инвариант
Prerequisites: A2. Skills: expressions.*, control-flow.*, scope как observable outcome. Понимать short-circuit, границы и прогресс цикла; самостоятельно сформулировать termination condition. Ошибки: off-by-one, мутируется range copy, затенение, неверное fallthrough, недостижимые ветви. Теория: T flowcontrol, S Statements.
Сейчас: access-logic-001, exclusive-loop-001, range-offset-001, kind-condition-001, decision-table-tests-001. Donors: Exercism armstrong-numbers, binary-search позднее, simple number classifiers; Golings scope. Originals: side effects/evaluation order лишь при distinct gap. Виды: implementation, debugging, тесты partitions. Достаточно: решение с новым инвариантом и контрпримером прежнему defect. Core; сложный поиск не обязательный gate.

### A4. Функции, параметры и результаты
Prerequisites: A2/A3. Skills: functions.signatures-results; variadic ввод после slices. Понимать копирование arguments, multiple/named results; самостоятельно проектировать ясный contract. Ошибки: tuple-модель, неявный return, переусложнённая сигнатура. Теория: T basics/functions, S Function types/Calls.
Сейчас: division-results-001, two-quantity-tests-001; variadic-tail-sums-001 и record-result-001 — поздние возвращения. Donors: Ardan functions, Birsi functions. Original нужен только для отсутствующего самостоятельного API, не ещё одного sum. Виды: implementation/debugging/tests. Достаточно: реализовать и вызвать новый multi-result API с явными границами. Core.

### A5. Первые тесты: наблюдаемый контракт
Prerequisites: A3/A4. Skills: testing.basics-diagnostics, начальное table-driven design. Понимать discovered tests и expected/actual; самостоятельно написать tests по внешнему contract. Ошибки: только happy path, assert самой реализации, молчаливое непопадание функции в test runner. Теория: P testing, официальный TableDrivenTests.
Сейчас: clamp-tests-001, total-counts-tests-001, decision-table-tests-001. Donors: SHAD tabletest после parser/time prerequisites, Birsi tests как изучаемые артефакты. Originals: narrow diagnostics/weak-suite только при gap. Виды: test-writing и debugging; не обязательно TDD-видеокурс. Достаточно: tests отличают reference от правдоподобных ошибок и дают useful failure. Core.

## B. Значения и коллекции

### B1. Массивы, индексирование и копирование
Prerequisites: A3/A4/A5. Skills: arrays.value-semantics/indexing-iteration. Понимать length как часть типа и копирование на assignment/call; самостоятельно менять элементы по индексу без range-copy ошибки. Теория: T arrays, S Array types/Index expressions. Ошибки: alias вместо copy, неверные bounds, предположение mutable range reference.
Сейчас: array-copy-001, array-copy-address-tests-001; seven-value-window-001 — смежная slice-практика, не доказательство array indexing. Donors: Ardan arrays, Golings arrays1, Exercism saddle-points/Pascal после matrix prerequisites. Originals: только отдельная indexing семантика без подходящего donor. Достаточно: тест показывает изменение оригинала/копии согласно contract. Core.

### B2. Слайс как представление: len/cap и границы
Prerequisites: B1. Skills: slices.header-model/len-cap-slicing/full-slice-expressions. Понимать copied descriptor и shared elements; самостоятельно вычислять chained view boundaries, отличать element sharing от будущего append reuse. Ошибки: cap=length, третий индекс задаёт length, full slice якобы копирует. Теория: T moretypes, S Slice expressions, B slices-intro.
Сейчас: slice-window-001, bounded-slice-view-001, guarded-item-001. Donors: Golings slices4, Ardan slices. Original — при новом API ownership contract, не косметической смене окна. Виды: implementation, test-writing, debugging. Достаточно: объяснить и проверить другую последовательность slice/mutate/append; nil public defect исправлен отдельно. Core.

### B3. Append, copy, overlap и владение
Prerequisites: B2. Skills: slices.append-growth/backing-array-aliasing/copy-overlap/preallocation/subslice-retention. Понимать reuse versus new storage, overlap, shallow copy и retention; самостоятельно выбрать view/owned copy. Ошибки: append без присваивания, сохранённые aliases, лишние нули make(len)+append, предположение о growth factor. Теория: S built-ins, B slices, P slices.
Сейчас: append-value-001, detached-append-001, overlapping-shift-tests-001, reserved-copy-001, detached-prefix-001; variadic-aliasing-001 позднее. Donors: Birsi collections, Ardan, Koans только после удаления exact-cap expectations. Originals: composite ownership gap. Достаточно: independent tests ловят mutation/overlap defects, компонент корректен при иной capacity. Core; allocation tuning после I5.

### B4. Nil/empty и применение коллекций
Prerequisites: B2/B3. Skills: slices.nil-empty и transfer arrays/slices/control-flow. Понимать представление пустого результата там, где API это различает; самостоятельно выбирать representation без случайного alias. Ошибки: len==0 как proof nilness, общая строка матрицы, недокументированная мутация. Теория: S Slice types/built-ins, P slices; JSON различия позже H3.
Сейчас: independent-scaled-copy-001, related filtering; отдельной nil-empty оценки нет. Donors: Birsi collections; Exercism binary-search/Pascal/sieve/sublist. Original nil-empty лишь после отрицательного donor review. Виды: implementation алгоритма, mutation-tests на representation, debug alias. Достаточно: новый контракт, а не повтор одного filter; simple algorithm может быть core, остальные — практика.

### B5. Maps: присутствие, ключи и наборы
Prerequisites: A3/A5/B1; dynamic keys после F2. Skills: maps.*. Понимать отсутствие versus stored zero, nil write, unspecified order и comparability; самостоятельно спроектировать ключ и order-independent test. Ошибки: потеря zero entries, строковая нормализация без контракта, map iteration как сортировка, incomparable dynamic key. Теория: S Map types/Comparison, B maps.
Сейчас: map-presence-001, map-take-001, quantity-presence-001; dictionary-add-tests-001/dictionary-delete-001 позднее. Donors: Exercism pangram, SHAD lrucache, Birsi collections. Original key-comparability только после review. Достаточно: тест различает absent/zero и корректность на другом comparable key. Core; cache component после E/F.

## C. Строки и Unicode

### C1. Строки, байты и форматирование
Prerequisites: A2/A3/B2. Skills: strings.* без преждевременного rune-transfer. Понимать immutable byte sequence и точные separators/format verbs; самостоятельно выбрать strings API и построить representation. Ошибки: Trim как точный prefix, Fields вместо Split, byte как отображаемый символ, случайное trimming. Теория: S Strings, P strings/fmt/strconv, B strings.
Сейчас: replace-byte-001, exact-prefix-001, join-values-001, typed-line-001, text-border-001. Donors: SHAD varfmt, Exercism isbn-verifier/crypto-square. Original: precision boundary без аналога. Достаточно: own tests на empty/repeated separators/exact whitespace и независимая текстовая функция. Core; длинная языковая грамматика дополнительная.

### C2. Unicode, UTF-8 и границы символов
Prerequisites: C1/B2. Skills: unicode.*, strings.byte-rune-conversions. Понимать bytes/code points/graphemes, offsets и malformed input; самостоятельно реализовать явную decoding policy. Ошибки: len как rune count, rune как grapheme, RuneError всегда означает плохой байт. Теория: B strings, P unicode/utf8, S range strings.
Сейчас: rune-offsets-001, decode-prefix-001, repair-utf8-001, utf8-validity-tests-001, rune-boundary-tests-001. Donors: SHAD utf8 только отдельные checked parts; Exercism micro-blog/reverse-string после license gate. Original: отсутствующий decoder edge case, не повтор готовых 11 lessons. Достаточно: tests различают malformed bytes и encoded replacement rune; новый boundary contract. Core; grapheme алгоритм не делать обязательным без отдельного scope.

### C3. Текстовый компонент и инвариант
Prerequisites: C1/C2/F3/H1. Skills: selected strings/unicode/errors/io transfer. Понимать границы parsing/validation/representation; самостоятельно собрать bounded text component. Ошибки: ложный round trip для lossy transform, partial invalid result, безразмерный input. Теория: P strings/io, source-specific grammar.
Сейчас: text-level-001, scan-lines-tests-001, formatted-assignment-001 — предшественники, не capstone. Donors: Exercism ocr-numbers/forth/tournament при допуске; SHAD line-reader. Original: только компонент, который демонстрирует отсутствующую комбинацию. Виды: multi-file implementation, tests, repair. Достаточно: новый format contract с validation и предсказуемым error policy. Один компонент core, остальные альтернативы.

## D. Функции глубже

### D1. Function values, closures и variadic semantics
Prerequisites: A4/B3. Skills: functions.function-values/closure-capture/variadic. Понимать capture lifetime и caller aliasing; самостоятельно создать callback API с independent state. Ошибки: snapshot assumption, общий state нескольких экземпляров, nil function, unintended variadic mutation. Теория: T closures, S Function values/Calls.
Сейчас: predicate-filter-001, period-predicate-001, category-closure-tests-001, stateful-counter-tests-001. Donors: Golings anonymous_functions3, Birsi functions, Exercism list-ops. Original: новый capture/ownership context, если donor повторяет counter. Достаточно: tests различают отдельные экземпляры и shared captures; пользователь проектирует callback самостоятельно. Core.

### D2. Defer, cleanup и panic boundaries
Prerequisites: A4/F3 для полного блока; простой defer можно объяснить раньше. Skills: functions.defer, errors.panic-recover-boundaries. Понимать argument evaluation, LIFO, function lifetime, narrow recover; самостоятельно обеспечить cleanup на всех допустимых путях. Ошибки: defer как block cleanup, args вычисляются на return, parent recover ловит child panic. Теория: S Defer/Panic/Recover, official defer-panic-recover.
Сейчас: deferred-release-001, capture-panic-001; cleanup-order-001 возвращение в testing. Donors: Ardan/Koans defer после review. Original: LIFO/argument evaluation, если текущая задача проверяет только release-once. Достаточно: объяснить и проверить изменённый order/lifetime без подавления ошибок. Core; произвольный panic management не самоцель.

## E. Структуры, pointers, методы и embedding

### E1. Модель значения: structs и pointers
Prerequisites: A2/B1/B3. Skills: structs.*, pointers.*. Понимать shallow value copy, copied pointer/shared pointee, nil и инвариант struct; самостоятельно выбрать value/pointer API. Ошибки: Go якобы pass-by-reference, pointer нужен всем shared values, scalar и referenced fields копируются одинаково. Теория: T structs/pointers, S Struct/Pointer types, FAQ pass-by-value.
Сейчас: initial-structs-001, struct-record-tests-001, struct-copy-tests-001, pointer-chain-001. Donors: Ardan, Birsi structs/pointers. Original: cross-field invariant при отсутствии donor. Достаточно: test отличает copied scalar и shared slice/map, API выдерживает zero/nil contract. Core.

### E2. Методы, method sets, values и expressions
Prerequisites: E1/D1. Skills: methods.*. Понимать receiver copying, membership T/*T и capture времени method value; самостоятельно построить method-based API. Ошибки: convenience addressable call равен interface satisfaction; method value и expression имеют одинаковую сигнатуру. Теория: T methods, S Method sets/Method values/Expressions.
Сейчас: receiver-reset-001, receiver-speakers-001, value-average-001; methods.values-expressions не имеет отдельной оценки. Donors: Ardan methods, Birsi methods, review Koans interfaces. Original: captured receiver snapshot/identity, если упражнения только вызывают метод. Достаточно: диагностировать новый mismatch и проверить наблюдение после изменения receiver. Core; редкие type-level детали углубление.

### E3. Embedding и promotion
Prerequisites: E2. Skills: embedding.field-promotion/method-promotion-ambiguity. Понимать composition вместо inheritance, selectors, pointer/value promotion; самостоятельно разрешать неоднозначность явно. Ошибки: first-field wins, embedded value и pointer одинаковы, promoted field считается непосредственно объявленным. Теория: S Selectors/Struct types, V для актуальных literals.
Сейчас: embedded-cache-tests-001; nested-stringers-001 не доказательство field promotion. Donors: Golings structs2 и Ardan embedding. Original: ambiguity/composition gap после просмотра конкретных templates. Достаточно: новый составной тип и tests исключают неверный selector/state sharing; trivial missing type name недостаточно. Core basics, сложные diamonds — дополнительная практика.

## F. Интерфейсы и ошибки

### F1. Минимальный поведенческий интерфейс
Prerequisites: E2/A5. Skills: interfaces.implicit-satisfaction/small-contracts/dynamic-type-value. Понимать consumer-side contract; самостоятельно написать потребителя с двумя независимыми реализациями/double. Ошибки: interface перечисляет все методы, concrete type switch вместо полиморфизма. Теория: T methods/interfaces, S Interface types, P io.
Сейчас: shape-contract-tests-001, greeter-dispatch-tests-001; writer-greeting-001 позднее. Donors: SHAD lrucache, Ardan interfaces, Birsi interfaces. Original: другой minimal consumer API только при gap. Достаточно: новый implementation подходит без изменения consumer и тесты проверяют contract, не внутренности. Core.

### F2. Dynamic type, assertions, typed nil, comparability
Prerequisites: F1/B5 для полного набора. Skills: interfaces.typed-nil/type-assertions/type-switches и dynamic values. Понимать type/value pair; самостоятельно проектировать nil и nonmatching policies. Ошибки: assertion является conversion, nil pointer внутри interface равен nil, сравнение любого dynamic value безопасно. Теория: S Assertions/Comparison, FAQ nil error.
Сейчас: typed-nil-result-001, exact-dynamic-number-001, dynamic-description-001, type-switch-default-tests-001. Donors: Ardan/Birsi/Koans после точного review. Original: dynamic incomparable key или иной nil boundary, не ещё один identical type switch. Достаточно: собственный test демонстрирует и устраняет typed-nil defect. Core.

### F3. Ошибки как recovery contract
Prerequisites: F1/F2 для всей глубины; plain error-return знакомство раньше A4. Skills: errors.*. Понимать sentinel/structured error/opaque boundary, wrapping и Is/As; самостоятельно выбрать что стабильно для caller. Ошибки: message equality, %v вместо %w, неправильный As target, ложное typed nil, изменение порядка joined tree. Теория: P errors/fmt, B go1.13-errors, S panic boundaries.
Сейчас: sentinel-lookup-001, ordered-error-propagation-001, error-context-wrap-001, error-chain-match-001, error-detail-tests-001. Donors: Birsi errors/Ardan; exact deep mechanisms уже есть. Original: missing multi-operation recovery boundary, не повтор As wrapper. Достаточно: consumer принимает решение через API и tests ловят реальный error-path defect. Core; custom Is/As intricacies дополнительная глубина.

### F4. Самостоятельный компонент interfaces/errors
Prerequisites: F3/H1/I1. Skills: выбранные interfaces/errors/testing transfer. Самостоятельно определить dependency boundary, classification/cleanup и tests. Ошибки: интерфейс ради интерфейса, loss of cause, tests зависят от concrete implementation. Теория: ссылки F1/F3/H1, без готового component implementation.
Сейчас: writer-greeting-001, writer-double-001, error-detail-tests-001 — только подготовка. Donors: SHAD lrucache/externalsort. Original: узкий package component при distinct gap. Виды: implementation + независимая test-writing работа в другом контексте. Достаточно: другую implementation зависимости можно подставить без изменения tests/consumer, error policy сохраняется. Один core component, альтернативы в банк.

## G. Пакеты, generics и итерация

### G1. Packages/modules/imports/initialization
Prerequisites: A1/E2/F1. Skills: packages.*, toolchain.modules-workspaces. Понимать visibility/internal boundary, dependency graph и init order; самостоятельно разделить package без import cycle. Ошибки: alias import решает цикл, init ломает цикл, всё repository может импортировать internal. Теория: F modules/workspaces, S Program initialization, go help packages.
Сейчас: package-inventory-001, initialization-values-001, external-window-tests-001. Donors: Ardan exporting/official examples. Original: offline import-cycle/internal fixtures, если donor не допускается в runner; не снимать go.work/external deps ограничения глобально. Достаточно: объяснить и исправить другой dependency graph. Core, build-system детали глубже отдельно.

### G2. Generic functions и constraints
Prerequisites: E2/F2. Skills: generics.type-parameters-inference/constraints-type-sets/approximation-unions. Понимать permitted operations/type set/defined types, самостоятельно выбрать минимальное ограничение. Ошибки: any разрешает всё, ~ это conversion, constraint interface всегда value type. Теория: F generics, S Type parameters/Interface types.
Сейчас: generic-find-tests-001, generic-reduce-001, named-numeric-constraint-001. Donors: SHAD genericsum, Birsi generics, Golings generics2. Original: distinct constraint boundary, не quota вариант Reduce. Достаточно: решение работает с неизвестным заранее defined type и отвергает недопустимые операции по contract. Core.

### G3. Generic types и version-sensitive методы
Prerequisites: G2/E2. Skills: generics.reusable-data-structures; новые Go-версионные outcomes требуют review. Понимать zero T и типовой API, самостоятельно написать container без лишних constraints. Ошибки: zero T всегда nil, generic methods ограничения старой версии переносятся на 1.27. Теория: S/V, P выбранного container API.
Сейчас: generic-stack-001. Donors: Birsi/Ardan checked generics; не объявлять ready coverage новых features. Original: generic alias/method contracts только после explicit taxonomy/version decision. Достаточно: новый container сохраняет invariant для pointer/value/zero types. Core types, optional advanced type-level cases.

### G4. Iteration и библиотечные helpers
Prerequisites: G2/B3/B5. Skills: control-flow.range + reviewed proposed iter/slices/maps outcomes. Понимать callback traversal versus iterator protocol, early-stop и order; самостоятельно выбрать подходящую абстракцию. Ошибки: source treeiter автоматически учит iter.Seq, stop callback игнорируется, map order assumed. Теория: S Range clauses, P iter/slices/maps.
Сейчас: range-offset-001 не подтверждает modern iterator protocol. Donors: SHAD treeiter — callback contract; official examples для iter после allowlist. Original: missing stop/lifetime behavior, если у donor нет этого механизма. Достаточно: tests прекращения обхода и другой consumer. Базовая iteration core, сложные iterator adapters углубление.

## H. I/O, encoding и time

### H1. Reader/Writer: partial result и error
Prerequisites: F3/D2. Skills: io.reader-writer-contracts/copy-limit-composition. Понимать независимость n и err; самостоятельно обработать short read/write и ограничение без чтения всего input. Ошибки: данные теряются при n>0+err, Read обязан заполнить buffer, caller-owned stream закрыт. Теория: P io, особенно Reader/Writer/Copy/LimitReader.
Сейчас: writer-greeting-001, bounded-copy-001, writer-double-001. Donors: SHAD externalsort line adapters. Original: exact ownership/partial-error intersection после donor review. Достаточно: custom chunked reader/writer выявляет неверную реализацию; full data policy публична. Core.

### H2. Buffering/scanning/files и cleanup
Prerequisites: H1/G1. Skills: io.buffered/scanning/files-cleanup. Понимать buffer owner/flush, token limits и file lifetime; самостоятельно собрать чтение записей с cleanup. Ошибки: Err после Scan игнорируется, final line теряется, defer в долгом loop, flush error потерян. Теория: P bufio/os/io/fs, D2.
Сейчас: buffered-flush-001, scan-lines-tests-001, filesystem-record-tests-001. Donors: SHAD externalsort. Original: отсутствующий flush/error ownership, не повтор newline case. Достаточно: tests на длинный token, EOF/partial/final record и cleanup всех owned handles. Core; внешняя сортировка интеграция после prerequisites.

### H3. JSON: форма и stream contract
Prerequisites: H1/F3/E1. Skills: encoding.json-shape-tags/streaming-strictness. Понимать absent/null/zero/unknown/trailing data; самостоятельно задать validation/error policy. Ошибки: успешный Decode означает отсутствие хвоста, unexported fields заполнены, partial result возвращён как valid. Теория: P encoding/json и версия API, V при сравнении v2.
Сейчас: json-field-shape-001, single-json-value-001. Donors: Birsi JSON после review, official selected examples. Original: существенный format boundary после проверки существующих задач. Достаточно: tests отвергают malformed/trailing/unknown согласно явному контракту. Core v1 с сохранением старых задач; v2 отличия отдельное чтение/практика, не тихая миграция.

### H4. Время как данные
Prerequisites: A2/E1/F3. Skills: time.duration-arithmetic/parsing-zones. Понимать duration versus civil instant/location, layouts и overflow; самостоятельно разобрать/сравнить входные времена без ambient now. Ошибки: strftime-layout, timezone abbreviation однозначна, 24h равно календарным суткам везде. Теория: P time, versioned docs.
Сейчас: duration-rounding-001, parse-timestamp-001, afternoon-window-001. Donors: SHAD tabletest ParseDuration after lineage; Exercism time concept. Original: calendar/monotonic boundary при explicit scope и отсутствии аналога. Достаточно: chosen contract проходит DST/offset/units cases без случайных системных часов. Core; timers позже L2.

### H5. Binary/text marshal contracts
Prerequisites: H1/A2/C1. Skills: encoding.binary-marshaling/text-marshaling. Понимать framing/endian/version и validation before mutation; самостоятельно написать симметричный codec. Ошибки: unchecked length allocation, truncation как validprefix, retained caller buffer, partial invalid receiver. Теория: P encoding/binary/encoding interfaces/strconv.
Сейчас: binary-record-tests-001, text-level-001. Donors: official allowlisted tests/examples; Exercism all-your-base не автоматически binary-format source. Original: bounded codec с reason, если донор требует другой protocol. Достаточно: valid roundtrip + invalid-input tests; не ложный roundtrip для lossy transform. Core small codec, сложные форматы углубление.

### H6. io.Pipe: связанный lifecycle
Prerequisites: H1/K2/L1. Skills: io.pipe-lifecycle и relevant concurrency transfer. Понимать synchronous backpressure и close/error authority обоих концов; самостоятельно исключить stranded reader/writer. Ошибки: write без reader завершится, cancellation сама закрывает pipe. Теория: P io.Pipe, M, L1.
Сейчас: bounded-copy-001 и cancel-worker-001 — подготовка, отдельной pipe-задачи нет. Donors: official example/Ardan if allowed; original missing cancellation/error intersection после review. Виды: implementation/debug/tests. Достаточно: normal/error/early-stop paths завершают обе стороны с выбранной ошибкой. Поздняя core практика или равнозначный компонент, не ранний I/O урок.

## I. Тестирование, fuzzing и experiments

### I1. Test design, helpers, tables и subtests
Prerequisites: A5/F3. Skills: testing.table-driven/helpers/subtests/basics-diagnostics. Понимать equivalence partitions, failure scope и имена; самостоятельно оформить tests читаемо и выбрать Fatal/Error. Ошибки: table ради table, индексы вместо diagnostics, Helper слишком поздно, assertion reproduces implementation. Теория: P testing, official TableDrivenTests.
Сейчас: diagnostic-helper-001, named-cases-001, decision-table-tests-001. Donors: SHAD tabletest, Birsi selected suites. Original: диагностика/независимый oracle после gap. Достаточно: weak suite не получает зачёт, полезная failure локализует нарушение. Core.

### I2. Изоляция, cleanup, doubles, examples/golden
Prerequisites: I1/H1/H4. Skills: testing.cleanup/seams-doubles/deterministic-boundaries/examples/golden. Понимать scope ресурсов и nondeterministic boundaries; самостоятельно проверить error/cleanup path без ambient clock/network. Ошибки: большой mock вместо small seam, defer и Cleanup одинаковы, golden автоматически обновляется, пример зависит от map order. Теория: P testing/fstest/io, B testing-time.
Сейчас: cleanup-order-001, writer-double-001, external-window-tests-001. Donors: SHAD selected testing, official examples; fileleak только после portability review. Original: explicit deterministic seam и controlled golden artifacts при отсутствии donor. Достаточно: тесты повторяются без изменения результатов, golden обновляется только явно. Core seams/examples, golden дополнительная практика.

### I3. Parallel tests и race diagnosis
Prerequisites: I1/J2/K1. Skills: testing.parallel/race-detector. Понимать parent-child pause/resume и isolation; самостоятельно прочитать конфликтующие accesses и сохранить regression. Ошибки: shared fixtures, t.Fatal из worker, старые loopvar assumptions, clean race run объявлен proof. Теория: P testing, official race detector docs, M.
Сейчас: join-all-tests-001/worker-completion-tests-001 связаны, но включённый -race не равен оценке этого skill. Donors: SHAD tparallel, Golings concurrent2 только если distinct from current counters. Original: missing diagnostic scenario. Достаточно: объяснён race report и исправлен защищаемый invariant без sleep. Core; performance of parallel tests отдельно.

### I4. Fuzzing: seeds, oracle, reproduction
Prerequisites: I1/H5. Skills: все три fuzzing skills. Понимать что именно утверждает property; самостоятельно написать target, сохранить counterexample, сделать deterministic regression. Ошибки: no crash достаточно всегда, тождественный roundtrip, ожидание случайной находки за N секунд. Теория: F fuzz, P testing.
Сейчас: assessed fuzzing-задач нет. Donors: официальные selected target examples; найденные банки не дают проверенной готовой замены. Original: invariant/oracle/reproduction gap после bounded source search. Профиль поддержать до контента. Достаточно: target различает seeded defects, reference проходит и воспроизведение не требует live fuzzing. Core; сложные corpus internals углубление.

### I5. Benchmarks и интерпретация эксперимента
Prerequisites: I1/G2. Skills: benchmarks.loops-timing/allocations/experimental-design. Понимать setup/timed work/result observation и representative workloads; самостоятельно сравнить семантически эквивалентные подходы. Ошибки: один запуск proof speedup, DCE, setup внутри измерения, faster but wrong. Теория: P testing benchmark API, V, official examples.
Сейчас: assessed benchmark-задач нет. Donors: SHAD allocs/utf8 как experiment candidates, Ardan selected examples. Original: trustworthy measurement task при отсутствии подходящего контрактного donor. Достаточно: корректное измерение и объяснение неопределённости, не fixed ns/op или universal allocations threshold. Core methodology; ручная оптимизация optional.

## J. Goroutines и shared state

### J1. Запуск, данные, результат и завершение
Prerequisites: D1/D2/I1. Skills: goroutines.launch-lifecycle/capture-arguments/result-error-collection/leak-prevention/panic-boundaries. Понимать что go-call не даёт join/ownership; самостоятельно зарегистрировать и собрать все завершения. Ошибки: return остановит children, capture всегда snapshot, parent recover, забытая goroutine. Теория: T concurrency, S Go statements, M.
Сейчас: worker-completion-tests-001, owned-values-001; capture/results/panic отдельной оценки нет. Donors: Birsi goroutines/channels, Ardan, loong. Original: missing all-error/early-failure boundary. Достаточно: final completion accounting и deterministic tests без wall-clock assumptions. Core launch/result, selective panic containment углубление.

### J2. Happens-before, Mutex и WaitGroup
Prerequisites: J1. Skills: sync.mutex-ownership/mutex-critical-sections/data-races-happens-before/waitgroup-lifecycle. Понимать защищаемый invariant и noncopyable lock; самостоятельно синхронизировать все read/write paths. Ошибки: locking fields отдельно сохраняет multi-field invariant, Add after launch safe, WaitGroup переносит errors. Теория: M, P sync.
Сейчас: mutex-counter-001, concurrent-counter-read-001, join-all-tests-001. Donors: SHAD keylock после review, Ardan; duplicate counter резерв. Original: multi-field invariant gap, если source только increments. Достаточно: аргумент happens-before + тест нарушенного invariant, не только green race detector. Core.

### J3. Once/RWMutex/atomics; Cond/Pool
Prerequisites: J2/I5 для всего блока; базовый Once можно раньше. Skills: sync.once-initialization/rwmutex-tradeoffs/atomics/cond/pool. Понимать failure/lifecycle/cost model и ограничения; самостоятельно выбрать простой корректный primitive. Ошибки: Once retries after panic, lock upgrade атомарен, atomics делают составное действие атомарным, Pool persistent cache, Cond signal запоминается. Теория: P sync/sync/atomic, M.
Сейчас: пять skills без assessed coverage. Donors: Ardan/SHAD, но SHAD reimplementation primitive — не introduction его обычного использования. Original: отсутствующие failure contracts. Достаточно: выбор обоснован invariant, tests и при необходимости measurements. Once/простая atomics-практика core; Cond/Pool/RWMutex optimization — углубление в банке.

## K. Каналы и pipeline

### K1. Operations, close, direction и ownership
Prerequisites: J1/J2. Skills: channels.operations-blocking/close-semantics/directional/ownership/range-drain. Понимать open/closed/nil/buffered state и safe close authority; самостоятельно завершить поток без потери данных. Ошибки: receiver closes, every channel must close, range stops on temporary empty, close discards buffer. Теория: T channels, S channels/range, P sync.
Сейчас: closed-receive-001, owned-values-001, worker-completion-tests-001. Donors: Birsi 11_channels, Golings concurrent3, Ardan. Original: ownership/drain contract при gap. Достаточно: tests отличают drained values от closed zero и все senders заканчивают. Core.

### K2. Select, buffers, backpressure и nil disabling
Prerequisites: K1. Skills: channels.select/buffering-backpressure/nil-disabling; cancellation introduction. Понимать множество ready outcomes, bounded in-flight work и отключение case; самостоятельно собрать multiplex loop. Ошибки: first source case priority, default waits, bigger buffer лечит leak, closed channel создаёт spin. Теория: S Select, M, P context как preview.
Сейчас: select-ready-case-001; отдельной оценки buffering/nil disabling нет. Donors: Birsi FirstValue, Ardan/loong. Original: controlled source close/no-spin boundary после review. Достаточно: принимаются все разрешённые schedules, progress и memory bounded по contract. Core.

### K3. Fan-out/fan-in и pipeline
Prerequisites: K2/L1/H1. Skills: channels.pipeline/fan-out-in/cancellation, goroutines.result-error-collection. Понимать topology/close authority/completeness; самостоятельно собрать bounded stages с partial-result/error policy. Ошибки: fan-in preserves source order, ранний consumer exit бросает blocked senders, close before producers done. Теория: B pipelines, M/P context.
Сейчас: primitives, отдельных assessed pipeline/fan задач нет. Donors: Birsi Merge, SHAD pubsub после review, Ardan patterns. Original: недостающая cancellation/error composition, не новый source label старого SumChannel. Достаточно: normal/empty/failure/early-stop завершают все стадии; concurrency limits публичны. Core один component, pubsub — альтернативная поздняя работа.

## L. Context и самостоятельные компоненты

### L1. Derivation, propagation и cancellation
Prerequisites: K1/K2/D2. Skills: context.derivation/cancellation/propagation-api. Понимать cooperative cancellation и child/parent relation; самостоятельно передать context и обеспечить cancel ownership/join. Ошибки: Background внутри chain, cancel автоматически убивает вычисление, child отменяет parent, cancel==worker done. Теория: P context, B pipelines/testing-time, M.
Сейчас: cancel-worker-001, propagate-context-001 — introduction. Donors: SHAD dupcall позднее, Ardan contexts, loong limit-service-time после адаптации. Original: отсутствующий blocking boundary. Достаточно: с новым API explicit/parent/early cancellation завершают owned work без утраты error policy. Core.

### L2. Deadlines, budgets и timers
Prerequisites: L1/H4/I3. Skills: context.deadlines/timeouts, time.timers-tickers. Понимать earliest deadline и относительный budget; самостоятельно тестировать timeout через controlled boundaries. Ошибки: child продлевает parent, Stop закрывает channel, ticker отдаёт все пропущенные ticks, full-duration retry даёт новый бюджет. Теория: P context/time и V.
Сейчас: нет отдельных assessed задач; synctest уже есть в инфраструктуре. Donors: Ardan/loong, SHAD ratelimit только после реального review. Original: exact budget/lifecycle policy если donor relies real sleep. Достаточно: тест на expiry/stop/reset/cancel не зависит от скорости machine и доказывает termination отдельно. Core.

### L3. Values, causes и API boundaries
Prerequisites: L1/F3. Skills: context.values-boundaries и selected context/errors outcomes. Понимать request-scoped metadata versus обычные arguments; самостоятельно сделать collision-safe accessor и cause/error contract. Ошибки: string global key, context как config bag, потеря cause, nil context. Теория: P context/errors.
Сейчас: cancel-worker-001 incidental values/cause preservation не оценивает самостоятельный values API. Donors: Ardan selected contexts, SHAD dupcall для совместной работы после review. Original: boundary gap при отсутствии allowed donor. Достаточно: независимые packages не конфликтуют и consumer не зависит от incidental error text. Core boundaries; advanced cause composition дополнительная практика.

### L4. Итоговые компоненты и возврат
Prerequisites: F4/G1/H2/H3/H6/I2/I4/K3/L2/L3 для всей итоговой подборки; каждый component использует только свой subset. Skills: выбранные core transfers, без автоматического покрытия всех prerequisites. Самостоятельно спроектировать несколько файлов по внешнему contract, написать tests и затем исправить новый defect. Ошибки: всё решено в starter, giant project без критериев, один компонент объявлен полным языком.
Теория: только ссылки и ранее изученные механизмы, без implementation recipe. Сейчас: 150 задач — подготовка. Donors: SHAD lrucache/externalsort/dupcall, Exercism forth/bowling/ocr-numbers при допуске. Originals: только не закрытая интеграция. Достаточно: выбранные independent components проверяют разные механизмы, есть delayed transfer и readable review. Несколько core результатов, остальные альтернативы/углубление; HTTP/DB не нужны.

## Матрица assessed coverage snapshot

Подсчёт по явным primary/secondary skills learner route; prerequisite и пример в чужих tests не counted occurrence. Это покрытие материалов, не mastery пользователя.

| Domain | Активных | Assessed | Без assessed |
|---|---:|---:|---:|
| toolchain | 4 | 0 | 4 |
| types | 7 | 7 | 0 |
| functions | 5 | 5 | 0 |
| expressions | 3 | 3 | 0 |
| control-flow | 4 | 4 | 0 |
| arrays | 2 | 1 | 1 |
| slices | 9 | 8 | 1 |
| strings | 5 | 5 | 0 |
| unicode | 4 | 4 | 0 |
| maps | 4 | 3 | 1 |
| structs | 3 | 3 | 0 |
| methods | 3 | 2 | 1 |
| pointers | 2 | 2 | 0 |
| embedding | 2 | 1 | 1 |
| interfaces | 7 | 7 | 0 |
| errors | 7 | 7 | 0 |
| generics | 4 | 4 | 0 |
| packages | 5 | 3 | 2 |
| io | 6 | 5 | 1 |
| encoding | 4 | 4 | 0 |
| time | 3 | 2 | 1 |
| testing | 11 | 5 | 6 |
| fuzzing | 3 | 0 | 3 |
| benchmarks | 3 | 0 | 3 |
| goroutines | 6 | 3 | 3 |
| channels | 11 | 4 | 7 |
| sync | 9 | 4 | 5 |
| context | 6 | 3 | 3 |
| Всего | 142 | 99 | 43 |

Точные 43 missing IDs:

```text
toolchain.modules-workspaces
toolchain.build-run
toolchain.formatting-vetting
toolchain.docs-environment
arrays.indexing-iteration
slices.nil-empty
maps.key-comparability
methods.values-expressions
embedding.field-promotion
packages.import-cycles
packages.internal-boundaries
io.pipe-lifecycle
time.timers-tickers
testing.table-driven
testing.deterministic-boundaries
testing.parallel
testing.race-detector
testing.examples
testing.golden
fuzzing.targets-seeds
fuzzing.invariants
fuzzing.reproduction
benchmarks.allocations
benchmarks.loops-timing
benchmarks.experimental-design
goroutines.capture-arguments
goroutines.result-error-collection
goroutines.panic-boundaries
channels.ownership
channels.range-drain
channels.buffering-backpressure
channels.nil-disabling
channels.cancellation
channels.fan-out-in
channels.pipeline
sync.rwmutex-tradeoffs
sync.once-initialization
sync.atomics
sync.cond
sync.pool
context.deadlines
context.timeouts
context.values-boundaries
```

Для каждого из 99 assessed skills следующий audit проверяет глубину, а не только ненулевой счётчик: какие точные outcomes statement/public/mutants действительно оценивают; introduction/practice/transfer; независимая разработка; test-writing/diagnosis; требуемые компоненты; что остаётся недоказанным. Такие labels, как arrays/indexing или testing.table-driven, нельзя дописать к manifests только потому, что конструкция случайно используется в reference/tests.

Полнота самой taxonomy требует versioned spec crosswalk. Проверить, где уже реально выражены scope/shadowing/evaluation order, bitwise/shifts, recursion, современный range/iter protocol, standard slices/maps helpers, generic aliases/methods, error trees и time monotonic/civil semantics. Для каждого — existing skill outcome, новое атомарное skill с ADR или явное optional/out-of-scope rationale. Не менять 142 автоматически и не включать весь runtime/compiler/unsafe/cgo/каждый stdlib package в обязательный фундамент.

## Углублённые критерии по трём опорным темам

### Слайсы: что должен доказать курс

Сначала ученик читает модель view и делает простую операцию bounds. Затем ему меняют capacity и проверяют, понимает ли он aliasing; далее он пишет test, который различает owned result и shared view. Самостоятельный component должен требовать выбрать ownership contract, а не назвать builtin. Поздняя задача с nested slice/retained prefix проверяет перенос понимания.

Проверять отдельно: nil versus empty, append result assignment, within-capacity mutation, reallocation independence, copy overlap/count, full-slice cap без copying, shallow versus deep copy, retention. Не все эти outcomes должны быть primary в одном упражнении. Runtime growth formula не проверяется как нормативный ответ. Allocation measurement — более поздний I5, не доказательство языка.

### Интерфейсы и ошибки: что должен доказать курс

Interface introduction не завершает тему. Нужны: method-set distinction; consumer с неизвестной заранее implementation; test double; nil interface versus typed nil; recovery через error identity/typed data; самостоятельный component, в котором boundary chosen осознанно. Дополнительные tests должны отличать поведение от implementation detail.

Existing error-detail-tests уже глубокий, поэтому его нельзя считать вводным без нужной теории и prerequisites. Custom As/join traversal могут быть углублением, в то время как самостоятельно спроектированный простой error API — core. Новый donor с тремя calls errors.Is не заменяет эту самостоятельность.

### Concurrency/context: что должен доказать курс

Для каждого owned worker требуется описание start/stop/completion, для каждого channel — send/close authority, для каждого blocking path — normal/error/cancel progression. Cancellation и completion не одно событие. Тесты строят условия, а не угадывают scheduler.

Проверять normal/empty/already-canceled/early-consumer-stop/failure-with-other-workers; explicit partial-result policy; read и write synchronization; no lost results/stranded goroutines. -race и synctest — инструменты свидетельства, не автоматическое доказательство отсутствия всех logical races. SHAD dupcall полезен после primitives именно разницей individual-caller versus shared-work lifetime.

## Пересмотр всех 46 прежних отказов Exercism

Это новые педагогические рекомендации. Исторические review records не переписаны этим planning PR. У каждой строки остаётся самостоятельный artifact/lineage gate: `exercism-problem-specifications` reference-only не снят. Не считать рекомендации release acceptance и не переносить canonical cases.

### 36 прежних reject-too-algorithmic

| Кандидат | Новое место | Обоснование / ограничение |
|---|---|---|
| affine-cipher | C1/A3 дополнительная практика | Арифметика и строковый contract; математические prerequisites объяснить, не обязательное введение криптографии |
| all-your-base | B4/H5 практика | Преобразование representation и validation; отдельно уточнить числовые bounds |
| alphametics | L4 углубление | Комбинаторный поиск не запрещён, но не core gate языка |
| armstrong-numbers | A3 практика | Небольшой цикл/границы; выбрать только при отличии от уже принятого числового контракта |
| binary-search | B2/B4 кандидат core | Индексный invariant, empty/singleton/not-found; порядок входа prerequisite |
| book-store | L4 углубление | Оптимизация комбинаций; высокая domain complexity |
| camicia | E1/B5/L4 optional capstone | State machine и повтор состояния; ограничить алгоритмические предпосылки |
| change | B4/L4 optional capstone | Минимальность/невозможность результата, не вводный loop |
| connect | B5/L4 optional capstone | Обход графа/visited; отдельно подготовить структуру |
| crypto-square | C1/B4 практика | Текстовые блоки и индексы; ASCII/Unicode contract явный |
| dominoes | B4/L4 optional capstone | Поиск и состояние, не превращать в простой arrays intro |
| go-counting | B5/L4 optional capstone | Связность и обход, независимое применение коллекций |
| killer-sudoku-helper | B4/L4 углубление | Combinatorial constraints, низкий приоритет относительно Go gaps |
| knapsack | B4/L4 optional capstone | DP/оптимизация после объяснения prerequisite, не внезапный hidden requirement |
| nth-prime | A3/B4 практика | Границы и выбор подхода; сопоставить с sieve по contract |
| palindrome-products | A3/C1 поздняя практика | Вложенный поиск, input bounds; не mandatory numerical puzzle |
| pascals-triangle | B1/B4 кандидат core | Многомерное хранилище/независимые строки/границы |
| perfect-numbers | A3 практика | Классификация и loop, не ещё один cosmetic classifier |
| pig-latin | C1 практика | Грамматика описана полностью; английские правила не скрытый prerequisite |
| pov | E1/L4 optional capstone | Переориентация дерева и связи; после дерева/identity |
| prism | E1/L4 углубление | Геометрическая симуляция и precision; поздний optional contract |
| pythagorean-triplet | A3 углубление | Числовой поиск допустим, но низкий приоритет |
| rail-fence-cipher | C1/B4 практика | Индексирование и обратимость, после simple text |
| rectangles | B4/C1 optional capstone | Сетка/несколько границ; ограниченный input contract |
| relative-distance | B5/L4 optional capstone | Графовая дистанция, явно объявить алгоритмические prerequisites |
| roman-numerals | A3/C1 практика | Табличная representation; сравнить с другими formatter contracts |
| saddle-points | B1/B4 кандидат core | Двумерный проход и два независимых условия |
| satellite | E1/L4 optional capstone | Восстановление дерева, не struct introduction |
| say | C1 дополнительная практика | Длинная грамматика/ветвления; не навязывать ради полного Go |
| sieve | B4 кандидат core | Слайс состояния и loop invariant |
| spiral-matrix | B1/B4 практика | Состояние/направление/bounds после простого matrix traversal |
| square-root | A2/A3 практика | Уточнить integer/float algorithm, tolerance и termination до authoring |
| sublist | B4 кандидат core | Отношения последовательностей/пустая подпоследовательность, иной contract чем поиск числа |
| two-bucket | B5/L4 углубление | Поиск в пространстве состояний, не обязательный языковой gate |
| word-search | B4/C1 optional capstone | Поиск по сетке и границы, не замена Unicode theory |
| zebra-puzzle | L4 углубление | Constraint puzzle допустима, но не нужна для основной готовности |

### 6 прежних reject-too-broad

| Кандидат | Новое место | Обоснование / ограничение |
|---|---|---|
| bowling | E1/F3/L4 optional capstone | State machine + validation; не ранняя атомарная задача |
| dot-dsl | E2/F1 дополнительный component | API composition, не обязательный отдельный курс DSL |
| forth | C3/B4/F3/L4 capstone option | Parser + stack + errors, bounded contract; сопоставить lineage SHAD forth |
| ocr-numbers | C3 capstone option | Text grid и malformed-input policy |
| piecing-it-together | E1/F3 углублённый component | Вывод параметров/противоречия; geometry assumptions публичны |
| poker | E1/B4/L4 optional capstone | Модель/сравнение/полные правила; не скрытая игра в grader |

### 4 прежних reject-duplicate

| Кандидат | Новое место | Почему прежнего основания недостаточно |
|---|---|---|
| bottle-song | C1 резерв практики | Одинаковый control-flow skill не доказывает эквивалентность грамматики; низкий приоритет относительно gaps |
| food-chain | C1 резерв практики | Накопительная representation может отличаться; объём fixtures отдельно оценить |
| isbn-verifier | C1/A5 кандидат core | Формальная grammar+checksum, другие ошибки, чем произвольная validation/Luhn |
| pangram | B5/C1 кандидат core | Полнота множества, не отсутствие повторов как у isogram |

Duplicate review должен назвать сравниваемую существующую задачу/источник и observable differences. Reconsideration «педагогически подходит» не снимает независимый provenance reject; восстановление источника через новую root license соседнего репозитория недопустимо.

## Контроль изменений карты

Сохранять historical route/task/progress IDs. При реализации генерировать coverage report из manifests/route, не из текста этой таблицы. Любой новый или переназначенный skill/outcome требует правил CURRICULUM_RULES и ADR при изменении контракта. Не использовать карту как готовый YAML schema: это проверяемая спецификация содержания, а не исполняемая миграция.
