package subject

func FastAllowed(blocked bool) bool                  { return false }
func AnyActive(a, b, c bool) bool                    { return false }
func SignalAllowed(blocked, ready bool) bool         { return false }
func ReleaseAllowed(a, b, ready, override bool) bool { return false }
