package subject_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	subject "task.local/bounded-record-read-001/subject"
)

func TestGraderEveryWireSplit(t *testing.T) {
	values := []uint16{0, 1, 65535, 256, 32768, 127, 513, 60001}
	var wire []byte
	var want []subject.Record
	for _, value := range values {
		wire = append(wire, 1, byte(value/256), byte(value%256))
		want = append(want, subject.Record{Value: value})
	}
	for split := 1; split < len(wire); split++ {
		t.Run(fmt.Sprintf("split-%d", split), func(t *testing.T) {
			source := publicSource(wire, io.EOF, true)
			source.splitAt = split
			got, err := subject.ReadRecords(source, len(want))
			checkPublicResult(t, got, err, want, nil)
			checkPublicSource(t, source, len(wire))
		})
	}
}

func TestGraderInvalidPositions(t *testing.T) {
	for position := 0; position < 8; position++ {
		for _, version := range []byte{0, 2, 255} {
			t.Run(fmt.Sprintf("position-%d-version-%d", position, version), func(t *testing.T) {
				var wire []byte
				var want []subject.Record
				for i := 0; i < position; i++ {
					wire = append(wire, 1, 0, byte(i+11))
					want = append(want, subject.Record{Value: uint16(i + 11)})
				}
				wire = append(wire, version, 23, 42, 1, 9, 9)
				source := publicSource(wire, io.EOF, false)
				source.maxChunk = 2
				got, err := subject.ReadRecords(source, position+1)
				checkPublicResult(t, got, err, want, subject.ErrVersion)
				checkPublicSource(t, source, 3*(position+1))
			})
		}
	}
}

func TestGraderEveryAcceptedBoundary(t *testing.T) {
	for count := 1; count <= 8; count++ {
		t.Run(fmt.Sprintf("count-%d", count), func(t *testing.T) {
			cause := errors.New(fmt.Sprintf("stopped at %d", count))
			var wire []byte
			var want []subject.Record
			for i := 0; i < count; i++ {
				wire = append(wire, 1, byte(i+1), byte(255-i))
				want = append(want, subject.Record{Value: uint16(i+1)*256 + uint16(255-i)})
			}
			source := publicSource(wire, cause, true)
			source.maxChunk = 1
			source.emptyReads = 2
			got, err := subject.ReadRecords(source, count)
			checkPublicResult(t, got, err, want, cause)
			checkPublicSource(t, source, len(wire))
		})
	}
}
