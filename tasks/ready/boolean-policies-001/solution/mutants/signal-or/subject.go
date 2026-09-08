package subject

func FastAllowed(blocked bool) bool                  { return !blocked }
func AnyActive(a, b, c bool) bool                    { return a || b || c }
func SignalAllowed(blocked, ready bool) bool         { return !blocked || ready }
func ReleaseAllowed(a, b, ready, override bool) bool { return !b && (override || (!a && ready)) }
