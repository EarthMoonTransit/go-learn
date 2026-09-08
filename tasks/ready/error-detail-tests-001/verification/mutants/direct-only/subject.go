package subject

type Detail struct {
	Code    int
	Message string
}

func (e *Detail) Error() string            { return e.Message }
func FindDetail(err error) (*Detail, bool) { d, ok := err.(*Detail); return d, ok }
