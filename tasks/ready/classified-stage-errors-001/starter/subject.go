package subject

import "errors"

var ErrTemporary = errors.New("temporary")

type Detail struct {
	Code    int
	Message string
}

func (d *Detail) Error() string {
	if d == nil {
		return "nil detail"
	}
	return d.Message
}

type Stage struct {
	Name   string
	Expose bool
	Run    func() (int, error)
}
type Outcome struct {
	Value     int
	Err       error
	Temporary bool
	Detail    *Detail
	HasDetail bool
}
type visibleContext struct{}

func (e *visibleContext) Error() string { return "" }
func (e *visibleContext) Unwrap() error { return nil }

type opaqueContext struct{}

func (e *opaqueContext) Error() string { return "" }
func RunStages(stages []Stage) Outcome { return Outcome{} }
