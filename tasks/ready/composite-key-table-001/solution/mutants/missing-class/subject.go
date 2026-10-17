package subject

type Key struct {
	Parts [2]int
	Class int
}
type table struct{ values map[Key]string }

func NewTable() *table                     { return &table{values: make(map[Key]string)} }
func Put(t *table, key *Key, value string) { copy := *key; copy.Class = 0; t.values[copy] = value }
func Lookup(t *table, key Key) (string, bool) {
	key.Class = 0
	value, ok := t.values[key]
	return value, ok
}
