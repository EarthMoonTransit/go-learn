package isolation

import (
	"crypto/sha256"
	"embed"
	"encoding/binary"
	"encoding/hex"
	"io/fs"
	"sync"
)

// Only production sources enter the runtime policy identity or executable.
// The companion test requires every new production file to be listed here.
//
//go:embed isolation.go legacy.go runtime.go source_identity.go
var policySources embed.FS

var policyDigest = sync.OnceValue(func() string {
	h := sha256.New()
	entries, err := fs.ReadDir(policySources, ".")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		data, err := policySources.ReadFile(entry.Name())
		if err != nil {
			panic(err)
		}
		_ = binary.Write(h, binary.BigEndian, uint64(len(entry.Name())))
		_, _ = h.Write([]byte(entry.Name()))
		_ = binary.Write(h, binary.BigEndian, uint64(len(data)))
		_, _ = h.Write(data)
	}
	return hex.EncodeToString(h.Sum(nil))
})

func SourceDigest() string { return policyDigest() }
