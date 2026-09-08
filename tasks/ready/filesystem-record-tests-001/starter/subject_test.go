package subject_test

import (
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

func TestStartingCase01(t *testing.T) {
	base := fstest.MapFS{"a.md": {Data: []byte("Title: Post 1\nDescription: Description 1\nTags: tdd, go\n---\nHello\nWorld")}, "b.md": {Data: []byte("Title: B\nDescription: D\nTags: rust, borrow-checker\n---\nB\nL\nM")}, "sub/file": {Data: []byte("ignored")}}
	fsys := &trackedFS{base: base}
	got, e := subject.ReadRecords(fsys)
	want := []subject.Post{{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: "Hello\nWorld"}, {Title: "B", Description: "D", Tags: []string{"rust", "borrow-checker"}, Body: "B\nL\nM"}}
	if e != nil || !reflect.DeepEqual(got, want) || !slices.Equal(fsys.opened, []string{"a.md", "b.md"}) || !slices.Equal(fsys.closed, fsys.opened) {
		t.Errorf("records got%+v,%v opens%v closes%v", got, e, fsys.opened, fsys.closed)
	}
}

// Add cases for the remaining requirements in TASK.md.
