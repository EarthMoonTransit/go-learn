package subject

const (
	code0, mask0 = iota + 10, uint8(1 << iota)
	code2, mask2
	code3, mask3
)
const (
	restart0 = iota
	restart1
)

func Layout() ([3]int, [3]uint8, [2]int) {
	return [3]int{code0, code2, code3}, [3]uint8{mask0, mask2, mask3}, [2]int{restart0, restart1}
}
