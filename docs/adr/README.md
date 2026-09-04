# Architecture decision records

ADR фиксируют изменения стабильных контрактов taxonomy, schema и pipeline.
`accepted` решения обязательны; `proposed` нельзя применять автоматически до
отдельного принятия и миграционного batch.

- `0001` — taxonomy v2 и expressions/control-flow layer (`accepted`)
- `0002` — декомпозиция широких skills (`superseded by 0003`)
- `0003` — taxonomy v3, atomic skill contracts и v2→v3 migration (`accepted`)
- `0004` — task lifecycle и первоначальная isolation model (`accepted`)
- `0005` — checker-owned test assembly, minimal env и promotion hardening (`accepted`)
