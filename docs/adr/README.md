# Architecture decision records

ADR фиксируют изменения стабильных контрактов taxonomy, schema и pipeline.
`accepted` решения обязательны; `proposed` нельзя применять автоматически до
отдельного принятия и миграционного batch.

- `0001` — taxonomy v2 и expressions/control-flow layer (`accepted`)
- `0002` — декомпозиция широких skills (`superseded by 0003`)
- `0003` — taxonomy v3, atomic skill contracts и v2→v3 migration (`accepted`)
- `0004` — task lifecycle и первоначальная isolation model (`accepted`)
- `0005` — checker-owned test assembly, minimal env и promotion hardening (`accepted`)
- `0006` — single-pass task verification и reusable check collection (`accepted`)
- `0007` — Task Contract v2 и declarative verification profiles (`accepted`)
- `0008` — multi-source Provenance v2 и original content identity (`accepted`)
- `0009` — learner work sessions across verification attempts (`accepted`)
- `0010` — sequential personal learning route (`accepted`)
- `0011` — editor API declarations for test submissions (`accepted`)
- `0012` — IDE metadata outside learner submissions (`accepted`)
- `0013` — local UI, shared operation coordination and workspace generations (`accepted`)
