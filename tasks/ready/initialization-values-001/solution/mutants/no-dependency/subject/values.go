package subject

var doubled = record("double", 0)
var base = record("base", 3)

func init()              { record("ready", 0) }
func Values() (int, int) { return base, doubled }
