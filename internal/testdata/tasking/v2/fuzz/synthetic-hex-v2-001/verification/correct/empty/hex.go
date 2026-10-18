package subject
import "encoding/hex"
func Encode(input []byte)[]byte{return hex.AppendEncode(nil,input)}
func Decode(wire []byte)([]byte,bool){if len(wire)==0{return nil,true};out,err:=hex.AppendDecode(nil,wire);return out,err==nil}
