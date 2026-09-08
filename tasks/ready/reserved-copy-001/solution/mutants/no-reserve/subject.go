package subject

func ReservedCopy(in []int, extra int) []int { out := make([]int, len(in)); copy(out, in); return out }
