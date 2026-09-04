package zerostate

type BufferState struct {
	Count   int
	Enabled bool
	Label   string
	Data    []byte
}

func InitialState() BufferState {
	return BufferState{}
}

func IsZero(state BufferState) bool {
	return state.Count == 0 && !state.Enabled && state.Label == "" && len(state.Data) == 0
}
