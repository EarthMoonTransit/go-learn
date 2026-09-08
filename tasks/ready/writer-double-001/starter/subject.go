package subject

type RecordingWriter struct {
	limit   int
	failure error
	calls   int
	data    []byte
}

func NewWriter(limit int, failure error) *RecordingWriter { return nil }
func (w *RecordingWriter) Write(data []byte) (int, error) { return 0, nil }
func (w *RecordingWriter) Snapshot() ([]byte, int)        { return nil, 0 }
