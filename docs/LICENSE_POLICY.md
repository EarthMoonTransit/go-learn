# License and source policy

Этот документ — инженерная политика provenance, не юридическая консультация.
При сомнении источник блокируется до отдельного review.

## Режимы использования

- `adapt`: лицензия и структура позволяют подготовить производную задачу при
  сохранении notices, attribution и иных условий. Каждый переносимый элемент
  фиксируется в `provenance.yaml`.
- `reference-only`: разрешено читать для фактчекинга и gap analysis. Запрещено
  копировать или близко перефразировать текст, код, тесты, fixtures и решения.
- `concepts-only`: используются только неохраняемые идеи языка и названия тем;
  реализация создаётся независимо без доступа author к исходной формулировке во
  время написания. Это консервативный режим для неоднозначной/неподходящей
  лицензии.
- `blocked`: adapter обязан отказать; источник не участвует даже в pipeline
  кандидатов до разрешения правового вопроса.

Usage mode — локальная политика проекта и может быть строже лицензии.

## Правила аудита

До регистрации источника проверяются canonical repository, exact 40-character
commit SHA, license file на этом commit, directory overrides, submodules и
смешанные лицензии. Тег или ветка не заменяют commit. SPDX используется только
когда условия однозначно ему соответствуют; иначе фиксируется `LicenseRef-*` и
источник обычно блокируется.

README badge, GitHub license detector и package metadata — подсказки, но не
заменяют текст лицензии. Отсутствие LICENSE означает «прав не предоставлено», а не
public domain.

## Производные материалы

Для `adapt` сохраняются copyright notice, license text/ссылка, source path, commit,
transformation type и описание существенных изменений. Совместимость лицензии
проверяется отдельно для публичного условия, starter, tests и solution. Нельзя
перелицензировать чужой материал шире исходных условий.

Copyleft, non-commercial, no-derivatives, custom/ambiguous terms и смешанные
лицензии по умолчанию переводятся в `reference-only`, `concepts-only` или
`blocked`; решение `adapt` для них требует отдельного ADR/legal review.

## Оригинальные задачи

`transformation_type: original` означает независимую авторскую работу на основе
taxonomy, спецификации языка и общих API-фактов. Она всё равно получает provenance
на внутреннюю source entry и commit первого review. Совпадение с известным
упражнением проверяется до публикации.

## Attribution output

Сборка публичного банка генерирует `THIRD_PARTY_NOTICES` из provenance, не из
текста условия. Любая ошибка или утрата provenance переводит задачу в
`quarantined` до восстановления цепочки.

## Authoring gate

Для нового task package `provenance.yaml` проходит отдельную schema и точно
совпадает с `task.yaml` и registry по source id, repository, path, commit,
license и usage mode. Promotion в `ready` требует reviewer evidence о лицензии,
clean-room для `reference-only`/`concepts-only` и отсутствии близкого копирования.
Автоматический validator проверяет согласованность metadata, но не объявляет
семантическую проверку копирования полностью автоматизируемой.
