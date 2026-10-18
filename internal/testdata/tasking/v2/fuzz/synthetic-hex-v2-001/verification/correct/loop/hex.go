package subject
import "encoding/hex"
func Encode(input []byte)[]byte{digits:=[]byte("0123456789abcdef");out:=make([]byte,len(input)*2);for i,v:=range input{out[i*2]=digits[v/16];out[i*2+1]=digits[v%16]};return out}
func Decode(wire []byte)([]byte,bool){out:=make([]byte,0,len(wire)/2);out,err:=hex.AppendDecode(out,wire);return out,err==nil}
