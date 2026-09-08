package subject_test

import (
	"strings"
	"task.local/single-json-value-001/subject"
	"testing"
)

type jsonReader struct {
	*strings.Reader
	closes int
}

func (r *jsonReader) Close() error { r.closes++; return nil }

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		text string
		want subject.Record
	}{{"{}", subject.Record{}}, {" \n{\"name\":\" λ \",\"count\":-2}\t", subject.Record{Name: " λ ", Count: -2}}, {"{\"name\":\"a\",\"name\":\"b\"}", subject.Record{Name: "b"}}} {
		v, e := subject.DecodeOne(strings.NewReader(c.text))
		if e != nil || v != c.want {
			t.Errorf("decode%q got%+v,%v want%+v", c.text, v, e, c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	sentinel := originalSentinel
	if sentinel == nil {
		t.Fatal("nil JSON sentinel")
	}
	for _, text := range []string{"", "null", "[]", "{\"extra\":1}", "{\"name\":2}", "{\"count\":1.5}", "{} {}", "{} null", "{} trailing", "{\"name\":\"x\""} {
		v, e := subject.DecodeOne(strings.NewReader(text))
		if v != (subject.Record{}) || e != sentinel || subject.ErrJSON != sentinel {
			t.Errorf("reject%q got%+v,%v wantzero,original", text, v, e)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	r := &jsonReader{Reader: strings.NewReader("{}")}
	_, e := subject.DecodeOne(r)
	if e != nil || r.closes != 0 {
		t.Errorf("success closes%d err%v", r.closes, e)
	}
	r = &jsonReader{Reader: strings.NewReader("{} x")}
	_, _ = subject.DecodeOne(r)
	if r.closes != 0 {
		t.Errorf("failure closes%d", r.closes)
	}
}

func TestDiscardPartialObject(t *testing.T) {
	for _, text := range []string{`{"name":"x"} {}`, `{"name":"x","extra":1}`} {
		v, e := subject.DecodeOne(strings.NewReader(text))
		if v != (subject.Record{}) || e != subject.ErrJSON {
			t.Errorf("partial object got%+v,%v wantzero,ErrJSON", v, e)
		}
	}
}

var originalSentinel = subject.ErrJSON
