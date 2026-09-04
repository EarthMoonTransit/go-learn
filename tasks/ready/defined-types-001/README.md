# Определённые числовые типы вместо aliases

Пакет `starter` использует aliases `Samples` и `Bytes`, поэтому компилятор не
различает количество отсчётов и размер закодированных данных.

Проведите рефакторинг:

```go
type Samples int
type Bytes int

func EncodedSize(count Samples, bytesPerSample int) Bytes
func Combine(left, right Samples) Samples
```

Сохраните поведение функций:

- `EncodedSize` умножает количество отсчётов на число байтов на отсчёт;
- `Combine` складывает два количества отсчётов;
- аргументы неотрицательны, результат помещается в `int`;
- `Samples` и `Bytes` должны остаться разными определёнными типами, а не
  aliases.

Примеры:

```text
EncodedSize(Samples(6), 2) → Bytes(12)
Combine(Samples(4), 3)     → Samples(7)
```

Работайте только в `starter/defined_types.go`. Public-проверка:

```bash
go test ./tests/public
```
