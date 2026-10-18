package subject_test
import("testing";"bytes";"encoding/hex";"task.local/synthetic-hex-v2-001/subject")
func FuzzHex(f *testing.F){f.Add([]byte(nil));f.Add([]byte{});f.Add([]byte{0});f.Add([]byte{128});f.Add([]byte{2,1});f.Add([]byte{7,7,7,7,7});f.Add([]byte{0x0a});f.Fuzz(func(t *testing.T,input []byte){wire:=subject.Encode(input);decoded,ok:=subject.Decode(wire);if !ok||!bytes.Equal(decoded,input){t.Error("decoded value/status mismatch")}})}
func unusedOracle(input []byte)[]byte{return hex.AppendEncode(nil,input)}
