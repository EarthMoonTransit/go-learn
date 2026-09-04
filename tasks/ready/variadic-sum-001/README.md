# Вариативный параметр и раскрытие slice

Реализуйте две функции в пакете `starter`:

```go
func Sum(values ...int) int
func SumSlice(values []int) int
```

`Sum` возвращает сумму всех переданных значений. Вызов без аргументов должен
возвращать `0`.

`SumSlice` возвращает сумму элементов `values`, вызывая `Sum` и раскрывая slice
как набор аргументов вариативного параметра. Входной slice изменять нельзя.

Примеры:

```text
Sum()                  → 0
Sum(1, 2, 3)           → 6
SumSlice([]int{4, -1}) → 3
```

Работайте только в `starter/variadic_sum.go`. Public-проверка:

```bash
go test ./tests/public
```
