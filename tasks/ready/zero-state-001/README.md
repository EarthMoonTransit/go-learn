# Проверка zero value составного состояния

В `starter/zero_state.go` объявлен тип:

```go
type BufferState struct {
    Count   int
    Enabled bool
    Label   string
    Data    []byte
}
```

Исправьте `IsZero`, чтобы она возвращала `true` только для точного zero value
`BufferState`. Для zero value выполняются все условия:

- `Count == 0`;
- `Enabled == false`;
- `Label == ""`;
- `Data == nil`.

Выделенный, но пустой slice (`[]byte{}` или `make([]byte, 0)`) не считается
нулевым состоянием, хотя его длина равна нулю.

`InitialState` менять не требуется. Работайте только в `starter/zero_state.go`.
Public-проверка:

```bash
go test ./tests/public
```
