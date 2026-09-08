package subject

type NamedInt int

func TypeLabel(value any) string {
	switch v := value.(type) {
	case *int:
		if v == nil {
			return "nil"
		}
		return "other"
	case []int:
		if v == nil {
			return "nil"
		}
		return "other"
	case nil:
		return "nil"
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "other"
	}
}
