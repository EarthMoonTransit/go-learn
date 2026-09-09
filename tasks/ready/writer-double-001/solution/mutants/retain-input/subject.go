package subject

type RecordingWriter struct {
	limit   int
	failure error
	calls   int
	data    []byte
}

func NewWriter(limit int, failure error) *RecordingWriter {
	return &RecordingWriter{limit: limit, failure: failure}
}
func (w *RecordingWriter) Write(data []byte) (int, error) {
	w.calls++
	n := len(data)
	remaining := w.limit - len(w.data)
	if n > remaining {
		n = remaining
	}
	if len(w.data) == 0 {
		w.data = data[:n]
	} else {
		w.data = append(w.data, data[:n]...)
	}
	if n < len(data) {
		return n, w.failure
	}
	return n, nil
}
func (w *RecordingWriter) Snapshot() ([]byte, int) { return append([]byte(nil), w.data...), w.calls }
