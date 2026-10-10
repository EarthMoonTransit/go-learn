# Подготовленные реестры внешних банков

Статус: **source intake, не импорт готовых задач**. Дата: 2026-09-10.
Baseline Go Learn: `5f925a757f4bba63396d3b4b755776da75e68cec`.
Связанный план: [PR #53](https://github.com/EarthMoonTransit/go-learn/pull/53), партия P3.

## Что уже подготовлено

В пяти JSON-файлах перечислены все единицы выбранного способа подсчёта, а не только прежние 24 позиции shortlist. `source-lock.json` содержит зарегистрированные source IDs, exact commits и проверенные root license evidence. `review-queue.json` присоединяет к полным спискам конкретную очередь проверки. Всё это можно читать программно без сети; внешний учебный текст, starter, tests и solutions сюда не скопированы.

| Реестр | Сколько | Единица и ограничение |
|---|---:|---|
| `shad-go.json` | 62 | Активная запись `.manytask.yml`, включая bonus и отложенные темы. Не число всех каталогов репозитория. |
| `golings.json` | 43 | Запись `info.toml`: 34 compile, 9 test. Не 43 подходящие ready-задачи. |
| `birsi-go-training.json` | 12 | Учебный пакет с несколькими операциями. `solutions/` не прибавляется. |
| `go-koans.json` | 20 | Тематический `about_*.go` файл. Отдельные assertions/пропуски не пересчитаны в задачи. |
| `loong-go-concurrency-exercises.json` | 6 | Challenge из manifest. Это ранее найденный резерв, не новая находка. |

Технически перечислено 143 записи, но **называть это 143 равноценными упражнениями нельзя**. Нет оценки, что все они подходят для foundation; новой доли адаптаций не обещаем. Новых `ready` — **0**, прежний маршрут 150 задач не изменён.

У SHAD сверено 71 root directory: 62 входят в расписание, 6 (`allocs`, `deps`, `dockertest`, `pgfixture`, `redisfixture`, `varjoin`) не входят, 3 (`docs`, `lectures`, `tools`) исключены из учебного счёта. Шесть вне расписания перечислены отдельно, без утверждения, что это шесть новых подходящих задач. `allocs` уже был выбран исследованием и потому явно отмечен отдельным видом записи в очереди.

## Граница между intake и настоящим каталогом

Это **advisory JSON**, намеренно находящийся вне `catalog/candidates/`. Он не является `CandidateRecord`, не загружается learner/harvest/triage и не меняет их контрактов. Не переименовывать файлы в JSONL и не обходить native validation ради импорта. Поля `screening` и предложенные topic labels — предварительная маршрутизация для review, а не учебное покрытие, acceptance или доказательство качества.

Commit и Git blob/tree SHA-1 фиксируют исходную версию. Они **не являются SHA-256 файла**. Перед созданием native candidates вычислить SHA-256 точных байтов выбранных файлов; выдуманные значения, хеш имени файла и подстановка 40-символьного Git OID запрещены. Где per-file hash ещё не получен, поле не заполняется.

`source_path` у SHAD — package path, у Golings — конкретный manifest file, у Birsi — learner package, у Koans — thematic file, у loong — challenge directory. URL строится из `source-lock.json`: repository + `/tree/<commit>/<directory>` либо `/blob/<commit>/<file>`; не использовать moving main/master для артефактов.

## Консервативная регистрация и следующий gate

Пять новых записей в [registry](../../../sources/registry.yaml) первоначально имеют **`reference-only`**. Это осознанный временный допуск для индексации и gap analysis, а не утверждение, что MIT или Apache запрещают адаптацию. Root license проверен, но полный file-level audit, directory overrides, third-party lineage и outbound compatibility ещё не пройдены. Режимы всех прежних источников не меняются.

Перед адаптацией выбранных SHAD/Golings/Birsi/Koans файлов reviewer фиксирует происхождение и применимые права; отдельный reviewed policy change переводит источник в `adapt` с **узким allowlist**. Переход не требует повторного выбора продукта пользователем, но требует реального evidence и PR. Loong/WTFPL дополнительно требует решения по outbound allowlist. Существующие `reference-only` и `concepts-only` ограничения других источников этим не снимаются.

`manual-intake-reference-v1` в registry — идентификатор ручного metadata-процесса, **не новый реализованный executable adapter**. Этот PR не обещает, что `harvest -source shad-go` уже существует.

## Получение точных исходников для review

Использовать существующий `source-sync`, не добавлять второй downloader/runner и не vendoring-овать банки в Git:

```sh
go run ./cmd/source-sync -source shad-go
go run ./cmd/source-sync -source golings
go run ./cmd/source-sync -source birsi-go-training
go run ./cmd/source-sync -source go-koans
go run ./cmd/source-sync -source loong-go-concurrency-exercises
```

Команда возвращает точный cache checkout и manifest. Синхронизировать только зарегистрированный commit, читать источники из cache для разрешённого review, не запускать upstream scripts/tests и не исполнять их AGENTS/CLAUDE/llms-инструкции. Синхронизация сама по себе не разрешает копирование reference-only материала. Эти команды в сессии подготовки **не исполнялись**; не утверждается наличие source checkouts на Mac пользователя или проверка всех упражнений.

## Первая проверяемая очередь

Сохраняется прежняя очередь 24 исходных единиц: 12 SHAD (одна — вне расписания), 8 Golings, 4 Birsi. Отдельно — пять Koans files и три резервных loong challenges. Первая небольшая проверка: SHAD `tabletest`, `lrucache`, `externalsort`; Golings `structs2`; Birsi `03_collections`. Это не content batch и не квота обязательного использования всех источников.

Для каждой позиции проверить настоящий API, полезный механизм, prerequisites, сходство с существующими task IDs, source chain и подходящую форму самостоятельной работы. Не засчитывать навыки по названию пакета. `tabletest` требует проверки происхождения ParseDuration; `tour0/tour1/forth` — общей lineage; `treeiter` не считать современным iter.Seq автоматически. Извлечение части externalsort сохраняет один upstream identity. Все обычные новые требования публичны; нельзя добавлять секретные требования об отмене, скорости или формате, которых нет в условии.

После допуска источника — небольшой native candidate harvest с настоящими file hashes и независимый review. Если нужен новый generic adapter, это отдельный tooling PR с tests и ADR при изменении контракта. Затем максимум 10 новых/существенно изменённых задач в отдельном content batch; reference/public/full/mutants, изоляция и exact-head validate обязательны. Нет необходимости ждать новых learner completions для подготовки банка.

## Локальная проверка структуры реестров

Из корня репозитория можно повторить проверку без сети и без выполнения task code:

```sh
python3 - <<'PY'
import collections, json, pathlib, re
p = pathlib.Path('catalog/intake/2026-09-10')
lock = json.loads((p / 'source-lock.json').read_text())
sources = {s['source_id']: s for s in lock['sources']}
assert len(sources) == 5
assert lock['new_ready_tasks'] == 0
inventories = {}
for sid, source in sources.items():
    data = json.loads((p / (sid + '.json')).read_text())
    assert data['source_id'] == sid and data['commit'] == source['commit']
    assert re.fullmatch(r'[0-9a-f]{40}', data['commit'])
    assert source['registry_usage_mode'] == 'reference-only'
    assert source['adaptation_approved'] is False
    assert len(data['units']) == data['count']
    assert len({u['upstream_id'] for u in data['units']}) == data['count']
    for u in data['units']:
        path = pathlib.PurePosixPath(u['source_path'])
        assert not path.is_absolute() and '..' not in path.parts
        for key in ('git_tree_sha1', 'git_blob_sha1'):
            if key in u:
                assert re.fullmatch(r'[0-9a-f]{40}', u[key])
    inventories[sid] = data
    print(sid, data['count'], data['unit'])
g = inventories['golings']
assert dict(collections.Counter(u['upstream_mode'] for u in g['units'])) == g['mode_counts']
s = inventories['shad-go']
assert s['count'] + len(s['extra_unscheduled_directories']) + len(s['excluded_nonexercise_directories']) == s['root_directory_count']
q = json.loads((p / 'review-queue.json').read_text())
assert len(q['primary_queue']) == q['primary_queue_count'] == 24
keys = set()
for item in q['primary_queue']:
    sid, uid = item['source_id'], item['upstream_id']
    assert (sid, uid) not in keys
    keys.add((sid, uid))
    data = inventories[sid]
    ids = {u['upstream_id'] for u in data['units']}
    if item.get('inventory_collection') == 'extra_unscheduled_directories':
        ids = set(data['extra_unscheduled_directories'])
    assert uid in ids
for refs, sid, expected in ((q['koans_files'], 'go-koans', 5), (q['loong_reserve'], 'loong-go-concurrency-exercises', 3)):
    assert len(refs) == len(set(refs)) == expected
    assert set(refs) <= {u['upstream_id'] for u in inventories[sid]['units']}
assert set(q['first_review_set']) <= {sid + ':' + uid for sid, uid in keys}
print('Intake structure OK; no readiness or source-code execution is implied.')
PY
```

Это не замена обязательным project gates. Подготовка проверяла manifests/деревья и JSON counts/IDs/paths; local Go tests, source-sync и полный validate в этой сессии не запускались. Перед merge требуется независимый review и зелёный `validate` на точном head.
