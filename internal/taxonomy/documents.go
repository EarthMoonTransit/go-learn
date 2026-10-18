package taxonomy

import (
	"bytes"
	"fmt"
	"io"
	"sync"

	"go-practice-bank/internal/safeio"
	"gopkg.in/yaml.v3"
)

// Each role retains only its latest successful parse, keyed by exact content.
// Published documents (including nested maps and YAML nodes) are immutable;
// replacing a slot never changes values still held by another reader.
var immutableDocuments [2]struct {
	sync.Mutex
	loaded  bool
	content string
	value   document
}

// readImmutableDocument is only for private, read-only taxonomy consumers. It
// always reads the file through all path/size guards before consulting a slot.
// Validation results, roots, migration records and reviewed bytes are not kept.
// Callers needing mutable data must use readStrict instead.
func readImmutableDocument(root, relative string) ([]byte, document, error) {
	b, err := readDocumentBytes(root, relative)
	if err != nil {
		return nil, document{}, err
	}
	role := -1
	switch relative {
	case V3Path:
		role = 0
	case "curriculum/skills.yaml":
		role = 1
	}
	if role == -1 {
		var value document
		if err := decodeStrictContent(relative, b, &value); err != nil {
			return nil, document{}, err
		}
		return b, value, nil
	}
	slot := &immutableDocuments[role]
	slot.Lock()
	defer slot.Unlock()
	if slot.loaded && slot.content == string(b) {
		return b, slot.value, nil
	}
	var value document
	if err := decodeStrictContent(relative, b, &value); err != nil {
		return nil, document{}, err
	}
	// The string owns a copy; the returned fresh bytes may be modified by callers.
	slot.content, slot.value, slot.loaded = string(b), value, true
	return b, value, nil
}

// readStrict deliberately returns a new decode on every call. In particular,
// migration/report tests and other mutable consumers must not share cached maps.
func readStrict(root, relative string, dst any) ([]byte, error) {
	b, err := readDocumentBytes(root, relative)
	if err != nil {
		return nil, err
	}
	if err := decodeStrictContent(relative, b, dst); err != nil {
		return nil, err
	}
	return b, nil
}

func readDocumentBytes(root, relative string) ([]byte, error) {
	b, err := safeio.ReadFile(root, relative, 4<<20)
	if err != nil {
		return nil, fmt.Errorf("taxonomy: read %s: %w", relative, err)
	}
	return b, nil
}

func decodeStrictContent(relative string, b []byte, dst any) error {
	d := yaml.NewDecoder(bytes.NewReader(b))
	d.KnownFields(true)
	if err := d.Decode(dst); err != nil {
		return fmt.Errorf("taxonomy: decode %s: %w", relative, err)
	}
	var trailing any
	if err := d.Decode(&trailing); err != io.EOF {
		return fmt.Errorf("taxonomy: trailing data in %s", relative)
	}
	return nil
}
