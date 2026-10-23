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
type visibleContext struct {
	text  string
	cause error
}

func (e *visibleContext) Error() string { return e.text }
func (e *visibleContext) Unwrap() error { return e.cause }

type opaqueContext struct{ text string }

func (e *opaqueContext) Error() string { return e.text }
func contextual(name string, cause error, expose bool) error {
	text := name + ": " + cause.Error()
	if expose {
		return &visibleContext{text: text, cause: cause}
	}
	return &opaqueContext{text: text}
}
func RunStages(stages []Stage) Outcome {
	var out Outcome
	for _, stage := range stages {
		value, cause := stage.Run()
		out.Value = value
		if cause != nil {
			out.Err = contextual(stage.Name, cause, stage.Expose)
			out.Temporary = errors.Is(out.Err, ErrTemporary)
			out.HasDetail = errors.As(out.Err, &out.Detail)
			if out.Detail != nil {
				copy := *out.Detail
				out.Detail = &copy
			}
			return out // failure
		}
	}
	return out // success
}
