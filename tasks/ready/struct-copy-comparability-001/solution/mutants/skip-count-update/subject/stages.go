package subject

func StructStages(count, a, b, pointed, newCount, newFirst, newPointed int) (source, copied RecordView, sameTarget bool) {

	cell := pointed
	original := Record{Count: count, Inline: [2]int{a, b}, Target: &cell}
	changed := original
	_ = newCount
	changed.Inline[0] = newFirst
	*changed.Target = newPointed
	source = viewRecord(original)
	copied = viewRecord(changed)
	sameTarget = original.Target == changed.Target
	return
}

func viewRecord(value Record) RecordView {
	return RecordView{Count: value.Count, Inline: value.Inline, Pointed: *value.Target}
}
func EqualRecord(left, right Record) bool          { return left == right }
func ComparableShapes() (record, sliceRecord bool) { return true, false }
