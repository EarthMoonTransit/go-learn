package subject

var doubled = record("double", base*2)
var base = record("base", 3)

func init()              { record("ready", 0) }
func Values() (int, int) { return base, doubled }
