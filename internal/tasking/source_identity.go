package tasking

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"io/fs"
	"sync"

	"go-practice-bank/internal/fuzznative"
	"go-practice-bank/internal/isolation"
)

// Bind the executable verifier as well as its finite admission kernel. Tests
// and fixtures never enter this embed or the execution source identity.
//
//go:embed check.go collection.go contract_fuzz.go contract_v2.go evidence.go
//go:embed execute_fuzz.go execute_v2.go fuzz_plan.go init.go learner.go
//go:embed learner_api.go learner_catalog.go learner_fuzz.go load.go model.go
//go:embed paths.go promote.go provenance_v2.go report.go runner.go
//go:embed runner_fuzz.go session.go source_identity.go statecheck.go
var verifierSources embed.FS

var verifierDigest = sync.OnceValue(func() string {
	h := sha256.New()
	writeDigestField(h, "native-kernel", []byte(fuzznative.SourceDigest()))
	writeDigestField(h, "isolation-policy", []byte(isolation.SourceDigest()))
	entries, err := fs.ReadDir(verifierSources, ".")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		data, err := verifierSources.ReadFile(entry.Name())
		if err != nil {
			panic(err)
		}
		writeDigestField(h, "verifier-source:"+entry.Name(), data)
	}
	return hex.EncodeToString(h.Sum(nil))
})

func VerifierSourceDigest() string { return verifierDigest() }
