package subject

type Result struct{ Code int }

func (r *Result) CodeValue() int { return r.Code }

type Coder interface{ CodeValue() int }

func ReadCode(value Coder) (int, bool) {
	if value == nil {
		return 0, false
	}
	if r, ok := value.(*Result); ok && r == nil {
		return 0, false
	}
	return value.CodeValue(), true
}
