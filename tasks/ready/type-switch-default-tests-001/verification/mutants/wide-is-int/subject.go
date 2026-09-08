package subject

type NamedInt int

func TypeLabel(value any) string {
	switch value.(type) {
	case nil:
		return "nil"
	case int, int64:
		return "int"
	case string:
		return "string"
	default:
		return "other"
	}
}
