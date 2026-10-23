package subject_test

import (
	"bytes"
	"encoding"
	subject "task.local/dual-frame-codec-001/subject"
	"testing"
)

var originalFrameError = subject.ErrFrame

type frameFormat struct {
	name      string
	marshal   func(subject.Frame) ([]byte, error)
	unmarshal func(*subject.Frame, []byte) error
}

var frameFormats = []frameFormat{
	{"binary", func(f subject.Frame) ([]byte, error) { var m encoding.BinaryMarshaler = f; return m.MarshalBinary() }, func(f *subject.Frame, b []byte) error {
		var u encoding.BinaryUnmarshaler = f
		return u.UnmarshalBinary(b)
	}},
	{"text", func(f subject.Frame) ([]byte, error) { var m encoding.TextMarshaler = f; return m.MarshalText() }, func(f *subject.Frame, b []byte) error { var u encoding.TextUnmarshaler = f; return u.UnmarshalText(b) }},
}

func equalFrame(a, b subject.Frame) bool { return a.Tag == b.Tag && bytes.Equal(a.Payload, b.Payload) }
func frameVector(t *testing.T, name string, tag byte, payload, binaryWire, textWire []byte) {
	t.Helper()
	for i, format := range frameFormats {
		{
			wire := binaryWire
			if i == 1 {
				wire = textWire
			}
			value := subject.Frame{Tag: tag, Payload: bytes.Clone(payload)}
			data, err := format.marshal(value)
			if err != nil || !bytes.Equal(data, wire) {
				t.Errorf("%s %s marshal=%x,%v; want %x,nil", name, format.name, data, err, wire)
			}
			if !bytes.Equal(value.Payload, payload) {
				t.Errorf("%s marshal changed receiver payload=%x; want %x", format.name, value.Payload, payload)
			}
			got := subject.Frame{Tag: 0xee, Payload: []byte{9, 8, 7, 6}}
			input := bytes.Clone(wire)
			before := bytes.Clone(input)
			err = format.unmarshal(&got, input)
			if err != nil || !equalFrame(got, subject.Frame{Tag: tag, Payload: payload}) {
				t.Errorf("%s %s decoded tag=%02x payload=%x error=%v; want tag=%02x payload=%x nil", name, format.name, got.Tag, got.Payload, err, tag, payload)
			}
			if !bytes.Equal(input, before) {
				t.Errorf("%s decode input=%x; want unchanged %x", format.name, input, before)
			}
		}
	}
}
func TestPublicExactFrameFormats(t *testing.T) {
	frameVector(t, "mixed-octets", 0x7e, []byte{0, 0xa5, 0xff}, []byte{1, 0x7e, 3, 0, 0xa5, 0xff}, []byte("1/7e/03/00a5ff"))
	frameVector(t, "empty-nil", 0, nil, []byte{1, 0, 0}, []byte("1/00/00/"))
	frameVector(t, "empty-allocated", 0xff, []byte{}, []byte{1, 0xff, 0}, []byte("1/ff/00/"))
	payload := make([]byte, 32)
	for i := range payload {
		payload[i] = byte(i)
	}
	frameVector(t, "maximum-payload", 0xff, payload, append([]byte{1, 0xff, 32}, payload...), []byte("1/ff/20/000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f"))
}
func rejectFrame(t *testing.T, format frameFormat, name string, wire []byte) {
	t.Helper()
	{
		backing := make([]byte, 32)
		for i := range backing {
			backing[i] = byte(100 + i)
		}
		backingBefore := bytes.Clone(backing)
		got := subject.Frame{Tag: 0x49, Payload: backing[:4]}
		input := bytes.Clone(wire)
		before := bytes.Clone(input)
		err := format.unmarshal(&got, input)
		if err != originalFrameError || originalFrameError == nil || subject.ErrFrame != originalFrameError {
			t.Errorf("%s reject %s error=%v; want original ErrFrame", format.name, name, err)
		}
		if got.Tag != 0x49 || !bytes.Equal(got.Payload, backingBefore[:4]) || !bytes.Equal(backing, backingBefore) {
			t.Errorf("%s rejection changed receiver: tag=%02x payload=%x; want49,%x", format.name, got.Tag, got.Payload, backingBefore[:4])
		}
		if !bytes.Equal(input, before) {
			t.Errorf("%s rejected input=%x; want unchanged %x", format.name, input, before)
		}
	}
}
func TestPublicBinaryRejectsInvalidFrames(t *testing.T) {
	cases := []struct {
		name string
		wire []byte
	}{
		{"nil", nil}, {"empty", []byte{}}, {"version-only", []byte{1}}, {"short-header", []byte{1, 0x7e}},
		{"wrong-version", []byte{2, 0x7e, 0}}, {"truncated-payload", []byte{1, 0x7e, 3, 0xa5, 0xff}},
		{"trailing-payload", []byte{1, 0x7e, 1, 0xa5, 0xff}}, {"empty-with-suffix", []byte{1, 0x7e, 0, 1}},
		{"oversize", append([]byte{1, 0x7e, 33}, make([]byte, 33)...)},
		{"largest-declared-length", append([]byte{1, 0x7e, 255}, make([]byte, 255)...)},
	}
	for _, c := range cases {
		rejectFrame(t, frameFormats[0], c.name, c.wire)
	}
}
func TestPublicTextRejectsInvalidFrames(t *testing.T) {
	cases := []struct{ name, wire string }{
		{"empty", ""}, {"short-header", "1/7e"}, {"wrong-version", "2/7e/00/"}, {"separator", "1-7e/00/"},
		{"uppercase-tag", "1/7E/00/"}, {"uppercase-payload", "1/7e/01/AF"}, {"short-tag", "1/7/00/"},
		{"invalid-length-digit", "1/7e/0z/"}, {"truncated-payload", "1/7e/01/a"}, {"trailing-payload", "1/7e/01/aaaa"},
		{"length-mismatch", "1/7e/02/aa"}, {"late-invalid-nibble", "1/7e/02/a5x0"}, {"empty-with-suffix", "1/7e/00/x"},
		{"leading-space", " 1/7e/00/"}, {"trailing-space", "1/7e/00/ "},
	}
	for _, c := range cases {
		rejectFrame(t, frameFormats[1], c.name, []byte(c.wire))
	}
	rejectFrame(t, frameFormats[1], "oversize", append([]byte("1/7e/21/"), bytes.Repeat([]byte("00"), 33)...))
}
func TestPublicOversizeMarshalIsAtomic(t *testing.T) {
	for _, format := range frameFormats {
		t.Run(format.name, func(t *testing.T) {
			payload := bytes.Repeat([]byte{0x9a}, 33)
			value := subject.Frame{Tag: 0xb2, Payload: payload}
			before := bytes.Clone(payload)
			data, err := format.marshal(value)
			if data != nil || err != originalFrameError || originalFrameError == nil {
				t.Errorf("%s oversize marshal=%x,%v; want nil,original ErrFrame", format.name, data, err)
			}
			if !bytes.Equal(payload, before) {
				t.Errorf("%s oversize marshal mutated payload=%x; want %x", format.name, payload, before)
			}
		})
	}
}
func TestPublicFrameStorageOwnership(t *testing.T) {
	for i, format := range frameFormats {
		t.Run(format.name, func(t *testing.T) {
			storage := make([]byte, 3, 32)
			copy(storage, []byte{0, 0xa5, 0xff})
			value := subject.Frame{Tag: 0x7e, Payload: storage}
			expected := []byte{1, 0x7e, 3, 0, 0xa5, 0xff}
			if i == 1 {
				expected = []byte("1/7e/03/00a5ff")
			}
			a, errA := format.marshal(value)
			b, errB := format.marshal(value)
			if errA != nil || errB != nil || !bytes.Equal(a, expected) || !bytes.Equal(b, expected) {
				t.Errorf("%s independent marshal outputs=%x/%x errors=%v/%v; want %x twice", format.name, a, b, errA, errB, expected)
			}
			if len(a) > 0 {
				a[len(a)-1] ^= 1
			}
			if !bytes.Equal(b, expected) {
				t.Errorf("%s marshal outputs share storage: second=%x; want %x", format.name, b, expected)
			}
			if !bytes.Equal(value.Payload, []byte{0, 0xa5, 0xff}) {
				t.Errorf("%s marshal changed receiver storage=%x; want00a5ff", format.name, value.Payload)
			}
			value.Payload[0] ^= 7
			if !bytes.Equal(b, expected) {
				t.Errorf("%s output aliases receiver after receiver mutation: output=%x; want %x", format.name, b, expected)
			}
			input := bytes.Clone(expected)
			before := bytes.Clone(input)
			var decoded subject.Frame
			if err := format.unmarshal(&decoded, input); err != nil {
				t.Errorf("%s ownership decode error=%v; want nil", format.name, err)
				return
			}
			if !bytes.Equal(input, before) {
				t.Errorf("%s ownership decode changed input=%x; want %x", format.name, input, before)
			}
			input[len(input)-1] ^= 1
			if decoded.Tag != 0x7e || !bytes.Equal(decoded.Payload, []byte{0, 0xa5, 0xff}) {
				t.Errorf("%s decoded payload retained input storage: tag=%02x payload=%x; want7e,00a5ff", format.name, decoded.Tag, decoded.Payload)
			}
			afterInputChange := bytes.Clone(input)
			if len(decoded.Payload) > 0 {
				decoded.Payload[0] ^= 3
			}
			if !bytes.Equal(input, afterInputChange) {
				t.Errorf("%s receiver mutation changed caller input=%x; want %x", format.name, input, afterInputChange)
			}
		})
	}
}
