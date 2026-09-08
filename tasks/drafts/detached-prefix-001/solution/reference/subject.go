package subject

func DetachedPrefix(in []int, n int) []int { out := make([]int, n); copy(out, in[:n]); return out }
