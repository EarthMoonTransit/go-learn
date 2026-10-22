package subject

type Record struct {
	Count  int
	Inline [2]int
	Target *int
}
type RecordView struct {
	Count   int
	Inline  [2]int
	Pointed int
}
type SliceRecord struct {
	Base  Record
	Extra []int
}
