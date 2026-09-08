package subject_test

import (
	"errors"
	"io/fs"
	"reflect"
	"slices"
	"task.local/filesystem-record-tests-001/subject"
	"testing"
	"testing/fstest"
)

type trackedFS struct {
	base                               fstest.MapFS
	failName                           string
	dirErr, openErr, readErr, closeErr error
	opened, closed                     []string
}

func (f *trackedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	if f.dirErr != nil {
		return nil, f.dirErr
	}
	return fs.ReadDir(f.base, name)
}
func (f *trackedFS) Open(name string) (fs.File, error) {
	if name == f.failName && f.openErr != nil {
		return nil, f.openErr
	}
	file, e := f.base.Open(name)
	if e != nil {
		return nil, e
	}
	f.opened = append(f.opened, name)
	return &trackedFile{File: file, owner: f, name: name}, nil
}

type trackedFile struct {
	fs.File
	owner  *trackedFS
	name   string
	closed bool
}

func (f *trackedFile) Read(p []byte) (int, error) {
	if f.closed {
		return 0, fs.ErrClosed
	}
	n, e := f.File.Read(p)
	if f.name == f.owner.failName && f.owner.readErr != nil {
		return n, f.owner.readErr
	}
	return n, e
}
func (f *trackedFile) Close() error {
	f.closed = true
	f.owner.closed = append(f.owner.closed, f.name)
	e := f.File.Close()
	if f.name == f.owner.failName && f.owner.closeErr != nil {
		return f.owner.closeErr
	}
	return e
}

func TestContractCase01(t *testing.T) {
	base := fstest.MapFS{"a.md": {Data: []byte("Title: Post 1\nDescription: Description 1\nTags: tdd, go\n---\nHello\nWorld")}, "b.md": {Data: []byte("Title: B\nDescription: D\nTags: rust, borrow-checker\n---\nB\nL\nM")}, "sub/file": {Data: []byte("ignored")}}
	fsys := &trackedFS{base: base}
	got, e := subject.ReadRecords(fsys)
	want := []subject.Post{{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: "Hello\nWorld"}, {Title: "B", Description: "D", Tags: []string{"rust", "borrow-checker"}, Body: "B\nL\nM"}}
	if e != nil || !reflect.DeepEqual(got, want) || !slices.Equal(fsys.opened, []string{"a.md", "b.md"}) || !slices.Equal(fsys.closed, fsys.opened) {
		t.Errorf("records got%+v,%v opens%v closes%v", got, e, fsys.opened, fsys.closed)
	}
}

func TestContractCase02(t *testing.T) {
	fsys := &trackedFS{base: fstest.MapFS{"x": {Data: []byte("Title:  x \nDescription: \nTags: \n---\n a \n\nλ\n")}}}
	got, e := subject.ReadRecords(fsys)
	if e != nil || len(got) != 1 || got[0].Title != " x " || got[0].Description != "" || len(got[0].Tags) != 0 || got[0].Body != " a \n\nλ" || len(fsys.closed) != 1 {
		t.Errorf("exact fields got%+v,%v closes%v", got, e, fsys.closed)
	}
	got, e = subject.ReadRecords(&trackedFS{base: fstest.MapFS{}})
	if e != nil || len(got) != 0 {
		t.Errorf("empty directory got%v,%v", got, e)
	}
}

func TestContractCase03(t *testing.T) {
	cause := errors.New("filesystem failed")
	valid := []byte("Title: T\nDescription: D\nTags: go\n---\nbody")
	for _, kind := range []string{"directory", "open", "read", "close"} {
		fsys := &trackedFS{base: fstest.MapFS{"a": {Data: valid}, "b": {Data: valid}}, failName: "b"}
		switch kind {
		case "directory":
			fsys.dirErr = cause
		case "open":
			fsys.openErr = cause
		case "read":
			fsys.readErr = cause
		case "close":
			fsys.closeErr = cause
		}
		got, e := subject.ReadRecords(fsys)
		if got != nil || e != cause || !slices.Equal(fsys.closed, fsys.opened) {
			t.Errorf("%s got%v,%v opens%v closes%v", kind, got, e, fsys.opened, fsys.closed)
		}
	}
}

func TestContractCase04(t *testing.T) {
	sentinel := subject.ErrFormat
	if sentinel == nil {
		t.Fatal("nil format sentinel")
	}
	for _, body := range []string{"Title: T\nDescription: D\nTags: go\nwrong", "Title:T\nDescription: D\nTags: go\n---", "Title: T\nDescription:D\nTags: go\n---", "Title: T\nDescription: D\nTags:go\n---", "Title: T"} {
		fsys := &trackedFS{base: fstest.MapFS{"x": {Data: []byte(body)}}, failName: "x", closeErr: errors.New("close failed")}
		got, e := subject.ReadRecords(fsys)
		if got != nil || e != sentinel || subject.ErrFormat != sentinel || !slices.Equal(fsys.closed, []string{"x"}) {
			t.Errorf("malformed got%v,%v closes%v", got, e, fsys.closed)
		}
	}
	readErr := errors.New("read first")
	fsys := &trackedFS{base: fstest.MapFS{"x": {Data: []byte("partial")}}, failName: "x", readErr: readErr, closeErr: errors.New("close second")}
	got, e := subject.ReadRecords(fsys)
	if got != nil || e != readErr || !slices.Equal(fsys.closed, []string{"x"}) {
		t.Errorf("error priority got%v,%v closes%v", got, e, fsys.closed)
	}
}
