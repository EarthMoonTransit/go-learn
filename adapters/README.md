# Source adapters

Adapter превращает запись аудированного источника в нормализованный **candidate
manifest**. Exercism adapter ФАЗЫ 2A извлекает только metadata, paths и hashes и
не переписывает учебный контент.

Контракт будущего adapter:

- принимает source id и зафиксированный commit из `sources/registry.yaml`;
- проверяет license boundary и `usage_mode` до чтения контента;
- выдаёт только метаданные кандидата: source path, обнаруженные skills, тип
  артефакта, hashes и решение о допустимом преобразовании;
- сохраняет immutable provenance;
- для `reference-only` и `concepts-only` никогда не выводит исходный текст или
  код;
- для `blocked` завершает работу ошибкой.

Идентификатор adapter в registry — будущая точка выбора реализации. `manual-*`
означает ручной review-first процесс, а не разрешение на копирование.
