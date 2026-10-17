package subject

type Key struct {
	Parts [2]int
	Class int
}
type table struct{ values map[Key]string }

var shared = make(map[Key]string)

func NewTable() *table                        { return &table{values: shared} }
func Put(t *table, key *Key, value string)    { t.values[*key] = value }
func Lookup(t *table, key Key) (string, bool) { value, ok := t.values[key]; return value, ok }
