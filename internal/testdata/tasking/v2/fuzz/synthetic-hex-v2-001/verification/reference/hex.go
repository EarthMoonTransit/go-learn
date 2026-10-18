package subject
import "encoding/hex"
func Encode(input []byte)[]byte{return hex.AppendEncode(nil,input)}
func Decode(wire []byte)([]byte,bool){out,err:=hex.AppendDecode(nil,wire);return out,err==nil}
