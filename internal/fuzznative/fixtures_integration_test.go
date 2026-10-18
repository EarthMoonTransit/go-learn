//go:build taskintegration

package fuzznative

import "strings"
import "fmt"

// Own synthetic codecs and tests from the reviewed private mechanism proof.
// They are source strings; only Engine may run them in disposable workspaces.
const encodeBody = `return hex.AppendEncode(nil,input)`
const decodeBody = `out,err:=hex.AppendDecode(nil,encoded);return out,err==nil`
const beforeDecode = `encoded:=subject.Encode(input);if len(encoded)!=len(input)*2{t.Fatal("length mismatch")};if !bytes.Equal(encoded,hex.AppendEncode(nil,input)){t.Error("canonical wire mismatch")};`
const strong = beforeDecode + `decoded,ok:=subject.Decode(encoded);if !ok||!bytes.Equal(decoded,input){t.Error("decoded payload/status mismatch")}`

var seeds = []string{"[]byte(nil)", "[]byte{}", "[]byte{0}", "[]byte{128}", "[]byte{2,1}", "[]byte{7,7,7,7,7}", "[]byte{0x0a}"}
var correct = []string{"correct-copy", "correct-loop", "correct-nil-empty"}
var defects = []string{"encode-drop-zero", "encode-drop-high", "encode-reverse", "encode-truncate", "decode-drop-zero", "decode-drop-high", "decode-reverse", "decode-truncate", "decode-false-status", "paired-uppercase", "paired-xor"}

func source(body string, remove int) map[string]string {
	adds := ""
	for i, s := range seeds {
		if i != remove {
			adds += "f.Add(" + s + ");"
		}
	}
	return map[string]string{"fuzz_test.go": `package subject_test
import("testing";"bytes";"encoding/hex";"task.local/integration-hex-001/subject")
func FuzzHex(f *testing.F){` + adds + `f.Fuzz(func(t *testing.T,input []byte){` + body + `})}
func unusedEquality(a,b []byte)bool{return bytes.Equal(a,b)}
func unusedOracle(seed []byte)[]byte{return hex.AppendEncode(nil,seed)}
`}
}
func maximum(assertions int) map[string]string {
	body := `encoded:=subject.Encode(input);decoded,ok:=subject.Decode(encoded);_ = ok;` + strings.Repeat(`subject.Decode(encoded);`, 6) + `if !bytes.Equal(decoded,input){for i:=range input{if i<` + fmt.Sprint(assertions) + `{t.Error("mismatch")}}}`
	return map[string]string{"fuzz_test.go": `package subject_test
import("testing";"bytes";"task.local/integration-hex-001/subject")
func FuzzHex(f *testing.F){seed:=[]byte{4095:0};cases:=[][]byte{` + strings.Repeat(`seed,`, 64) + `};for _,seed:=range cases{f.Add(seed)};f.Fuzz(func(t *testing.T,input []byte){` + body + `})}
`}
}
func mutatePayload(kind, input string) string {
	switch kind {
	case "drop-zero":
		return `out:=make([]byte,0,len(` + input + `));for _,v:=range ` + input + `{if v!=0{out=append(out,v)}};`
	case "drop-high":
		return `out:=make([]byte,0,len(` + input + `));for _,v:=range ` + input + `{if v<128{out=append(out,v)}};`
	case "reverse":
		return `out:=make([]byte,len(` + input + `));for i:=range ` + input + `{out[len(` + input + `)-1-i]=` + input + `[i]};`
	case "truncate":
		return `n:=len(` + input + `);if n>4{n=4};out:=make([]byte,n);copy(out,` + input + `);`
	default:
		panic("unknown payload mutation")
	}
}
func subject(variant string) string {
	enc, dec, extra := encodeBody, decodeBody, ""
	switch {
	case variant == "correct-copy":
	case variant == "correct-loop":
		enc = `digits:=[]byte("0123456789abcdef");out:=make([]byte,len(input)*2);for i,v:=range input{out[2*i]=digits[v/16];out[2*i+1]=digits[v%16]};return out`
		dec = `out:=make([]byte,0,len(encoded)/2);out,err:=hex.AppendDecode(out,encoded);return out,err==nil`
	case variant == "correct-nil-empty":
		enc = `digits:=[]byte("0123456789abcdef");var out []byte;for _,v:=range input{out=append(out,digits[v/16],digits[v%16])};return out`
		dec = `if len(encoded)==0{return nil,true};out:=make([]byte,hex.DecodedLen(len(encoded)));n,err:=hex.Decode(out,encoded);return out[:n],err==nil`
	case strings.HasPrefix(variant, "encode-"):
		enc = mutatePayload(strings.TrimPrefix(variant, "encode-"), "input") + `return hex.AppendEncode(nil,out)`
	case variant == "decode-false-status":
		dec = `out,_:=hex.AppendDecode(nil,encoded);return out,false`
	case variant == "decode-xor-boundary":
		dec = `out,err:=hex.AppendDecode(nil,encoded);for i:=range out{out[i]^=0x55};return out,err==nil`
	case strings.HasPrefix(variant, "decode-"):
		dec = `payload,err:=hex.AppendDecode(nil,encoded);if err!=nil{return nil,false};` + mutatePayload(strings.TrimPrefix(variant, "decode-"), "payload") + `return out,true`
	case variant == "paired-uppercase":
		enc = `out:=hex.AppendEncode(nil,input);for i,v:=range out{if v>='a'&&v<='f'{out[i]=v-32}};return out`
		dec = `for _,v:=range encoded{if v>='a'&&v<='f'{return nil,false}};out,err:=hex.AppendDecode(nil,encoded);return out,err==nil`
	case variant == "paired-xor":
		enc = `payload:=make([]byte,len(input));for i,v:=range input{payload[i]=v^0x55};return hex.AppendEncode(nil,payload)`
		dec = `out,err:=hex.AppendDecode(nil,encoded);for i:=range out{out[i]^=0x55};return out,err==nil`
	case strings.HasPrefix(variant, "raw-"):
		fault := ""
		switch variant {
		case "raw-panic":
			fault = `panic("instructor panic")`
		case "raw-goexit":
			extra = `import "runtime"`
			fault = `runtime.Goexit()`
		case "raw-exit1":
			extra = `import "os"`
			fault = `os.Exit(1)`
		case "raw-timeout":
			extra = `import "time"`
			fault = `time.Sleep(time.Hour)`
		case "raw-race":
			extra = `import "sync"`
			fault = `start:=make(chan struct{});var group sync.WaitGroup;var shared int;group.Add(2);go func(){defer group.Done();<-start;shared=1}();go func(){defer group.Done();<-start;shared=2}();close(start);group.Wait();_ = shared`
		case "raw-io-after-error":
			extra = `import "os"`
			fault = `p:=os.Getenv("GO_LEARN_CODEC_TRACE");os.Remove(p);os.Mkdir(p,0700)`
		case "raw-output-over":
			fault = `return make([]byte,8193)`
		default:
			panic(variant)
		}
		enc = `if len(input)==1&&input[0]==128{` + fault + `};if len(input)==1&&input[0]==0{return nil};` + encodeBody
	default:
		panic(variant)
	}
	return `package subject
import "encoding/hex"
` + extra + `
func Encode(input []byte)[]byte{` + enc + `}
func Decode(encoded []byte)([]byte,bool){` + dec + `}
`
}
