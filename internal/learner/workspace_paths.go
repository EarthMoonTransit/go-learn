package learner

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (s *Service) numberedWorkspaceRelative(id string) (string, error) {
	if !learnerTaskIDPattern.MatchString(id) {
		return "", errors.New("invalid task ID")
	}
	route, err := loadRoute(s.root)
	if err != nil {
		return "", err
	}
	if route != nil && route.SchemaVersion == 2 {
		for _, b := range route.Blocks {
			for _, step := range b.Steps {
				if step.TaskID == id {
					slug := id[:strings.LastIndex(id, "-")]
					return fmt.Sprintf("workspaces/%02d-%s/%02d-%s", b.Number, b.ID, step.Number, slug), nil
				}
			}
		}
	}
	return "workspaces/" + id, nil
}
func (s *Service) checkWorkspaceRelative(id, relative string) error {
	if !learnerTaskIDPattern.MatchString(id) || !safeWorkspaceRelative(relative) {
		return errors.New("invalid learner workspace path")
	}
	if relative == "workspaces/"+id {
		return nil
	}
	expected, err := s.numberedWorkspaceRelative(id)
	if err != nil {
		return err
	}
	if relative != expected {
		return errors.New("workspace path does not match its route assignment")
	}
	return nil
}
func safeWorkspaceRelative(relative string) bool {
	if !filepath.IsLocal(relative) || filepath.ToSlash(filepath.Clean(relative)) != relative {
		return false
	}
	parts := strings.Split(relative, "/")
	if len(parts) < 2 || len(parts) > 3 || parts[0] != "workspaces" {
		return false
	}
	for _, part := range parts[1:] {
		if !routeBlockIDPattern.MatchString(part) {
			return false
		}
	}
	return true
}
func pathAncestors(relative string) []string {
	var result []string
	current := ""
	for _, part := range strings.Split(filepath.ToSlash(relative), "/") {
		current = filepath.Join(current, part)
		result = append(result, current)
	}
	return result
}

// Traverse real directories through pinned handles. Every component is checked;
// a symlinked parent is rejected even if its target stays inside learner home.
func openPrivateDirectory(home, relative string) (*os.Root, error) {
	if !filepath.IsLocal(relative) {
		return nil, errors.New("invalid private directory")
	}
	info, err := os.Lstat(home)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return nil, errors.New("learner home must be a real directory")
	}
	root, err := os.OpenRoot(home)
	if err != nil {
		return nil, err
	}
	opened, err := root.Stat(".")
	if err != nil || !os.SameFile(info, opened) {
		root.Close()
		return nil, errors.New("learner home changed")
	}
	for _, part := range strings.Split(filepath.ToSlash(relative), "/") {
		if part == "." {
			continue
		}
		before, err := root.Lstat(part)
		if err != nil {
			root.Close()
			return nil, err
		}
		if !before.IsDir() || before.Mode()&os.ModeSymlink != 0 {
			root.Close()
			return nil, errors.New("private directory must be a real directory")
		}
		next, err := root.OpenRoot(part)
		if err != nil {
			root.Close()
			return nil, err
		}
		after, err := next.Stat(".")
		root.Close()
		if err != nil || !os.SameFile(before, after) {
			next.Close()
			return nil, errors.New("private directory changed")
		}
		root = next
	}
	return root, nil
}
func ensurePrivateDirectory(home, relative string) error {
	if !filepath.IsLocal(relative) {
		return errors.New("invalid private directory")
	}
	root, err := openPrivateDirectory(home, ".")
	if err != nil {
		return err
	}
	for _, part := range strings.Split(filepath.ToSlash(relative), "/") {
		if part == "." {
			continue
		}
		if err := root.Mkdir(part, 0700); err != nil && !errors.Is(err, os.ErrExist) {
			root.Close()
			return err
		}
		info, err := root.Lstat(part)
		if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
			root.Close()
			return errors.New("private directory must be a real directory")
		}
		next, err := root.OpenRoot(part)
		if err != nil {
			root.Close()
			return err
		}
		after, err := next.Stat(".")
		root.Close()
		if err != nil || !os.SameFile(info, after) {
			next.Close()
			return errors.New("private directory changed")
		}
		root = next
	}
	return root.Close()
}
