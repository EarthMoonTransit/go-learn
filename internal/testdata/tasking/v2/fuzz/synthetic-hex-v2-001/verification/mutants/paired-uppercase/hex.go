package subject
import "encoding/hex"
func Encode(input []byte)[]byte{out:=hex.AppendEncode(nil,input);for i,v:=range out{if v>='a'&&v<='f'{out[i]=v-32}};return out}
func Decode(wire []byte)([]byte,bool){out,err:=hex.AppendDecode(nil,wire);return out,err==nil}
