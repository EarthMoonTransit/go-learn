package subject

type Result struct{ Code int }

func (r *Result) CodeValue() int { return r.Code }

type Coder interface{ CodeValue() int }

func ReadCode(value Coder) (int, bool) { return 0, false }
