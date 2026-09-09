package subject

import (
	"bufio"
	"errors"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description string
	Tags               []string
	Body               string
}

var ErrFormat = errors.New("invalid record format")

func parsePost(src io.Reader) (Post, error) {
	scan := bufio.NewScanner(src)
	var lines []string
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	if err := scan.Err(); err != nil {
		return Post{}, err
	}
	if len(lines) < 4 || !strings.HasPrefix(lines[0], "Title: ") || !strings.HasPrefix(lines[1], "Description: ") || !strings.HasPrefix(lines[2], "Tags: ") || lines[3] != "---" {
		return Post{}, ErrFormat
	}
	tags := strings.TrimPrefix(lines[2], "Tags: ")
	var values []string
	if tags != "" {
		values = strings.Split(tags, ", ")
	}
	return Post{Title: strings.TrimPrefix(lines[0], "Title: "), Description: strings.TrimPrefix(lines[1], "Description: "), Tags: values, Body: strings.Join(lines[4:], "\n")}, nil
}
func ReadRecords(fileSystem fs.FS) ([]Post, error) {
	entries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		file, err := fileSystem.Open(entry.Name())
		if err != nil {
			return nil, err
		}
		post, readErr := parsePost(file)
		closeErr := file.Close()
		if readErr != nil {
			return posts, readErr
		}
		if closeErr != nil {
			return nil, closeErr
		}
		posts = append(posts, post)
	}
	return posts, nil
}
