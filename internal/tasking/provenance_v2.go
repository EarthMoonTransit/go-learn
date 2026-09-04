package tasking

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type projectIdentity struct {
	SchemaVersion            int    `yaml:"schema_version" json:"schema_version"`
	ID                       string `yaml:"id" json:"id"`
	Repository               string `yaml:"repository" json:"repository"`
	License                  string `yaml:"license" json:"license"`
	ReleaseManifestDirectory string `yaml:"release_manifest_directory" json:"release_manifest_directory"`
}

type outboundLicensePolicy struct {
	SchemaVersion           int      `yaml:"schema_version" json:"schema_version"`
	OriginalContentLicense  string   `yaml:"original_content_license" json:"original_content_license"`
	ThirdPartyRule          string   `yaml:"third_party_rule" json:"third_party_rule"`
	CompatibleSourceLicense []string `yaml:"compatible_source_licenses" json:"compatible_source_licenses"`
	NoticesFile             string   `yaml:"notices_file" json:"notices_file"`
}

func loadProjectIdentity(root string) (projectIdentity, error) {
	var identity projectIdentity
	path := filepath.Join(root, "config", "project-identity.yaml")
	if err := decodeAndValidate(root, path, "project-identity.schema.json", &identity); err != nil {
		return projectIdentity{}, fmt.Errorf("project identity: %w", err)
	}
	return identity, nil
}

func loadOutboundLicensePolicy(root string) (outboundLicensePolicy, error) {
	var policy outboundLicensePolicy
	path := filepath.Join(root, "config", "outbound-license.yaml")
	if err := decodeAndValidate(root, path, "outbound-license-policy.schema.json", &policy); err != nil {
		return outboundLicensePolicy{}, fmt.Errorf("outbound license policy: %w", err)
	}
	return policy, nil
}

// TaskContentDigest binds authored learner-facing and verification artifacts,
// not lifecycle/review metadata or the provenance record containing the digest.
func TaskContentDigest(taskDirectory string) (string, error) {
	hash := sha256.New()
	var paths []string
	err := filepath.WalkDir(taskDirectory, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == taskDirectory || entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(taskDirectory, path)
		if err != nil {
			return err
		}
		relative = filepath.ToSlash(relative)
		if relative != "TASK.md" && !contentDigestPrefix(relative) {
			return nil
		}
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			return fmt.Errorf("content digest input must be a regular file: %s", relative)
		}
		paths = append(paths, relative)
		return nil
	})
	if err != nil {
		return "", err
	}
	sort.Strings(paths)
	for _, relative := range paths {
		content, err := os.ReadFile(filepath.Join(taskDirectory, filepath.FromSlash(relative)))
		if err != nil {
			return "", err
		}
		writeDigestField(hash, relative, content)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func contentDigestPrefix(path string) bool {
	for _, prefix := range []string{"starter/", "tests/", "solution/", "workspace/", "verification/", "hints/"} {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}
