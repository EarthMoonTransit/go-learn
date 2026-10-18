# Attribution

Concept source: The Go Authors, `encoding/hex` tests, registry `golang-go`.
Repository: https://go.googlesource.com/go
Commit: `2c31595783c74564c1299a5401dc6e606f379cc7`.
Artifact: `src/encoding/hex/hex_test.go`, SHA-256
`b59de7b29f9b3d0bd75d661666d57fed17d87574b79f5a299a92f9c0f88123bb`.

The retained concept is canonical lowercase hexadecimal encoding with
value-preserving decoding. This source contains ordinary string/error tests,
not a native fuzz exercise. Its vectors, statements, code and tests were not
copied or closely paraphrased. The local byte/bool API, seed choices, assertions,
correct implementations and defect models are independently authored.
Classification: adapted provenance, concept-seed / concept-only;
no substantive adaptation or upstream native-fuzz exercise is claimed.
Admission: `catalog/intake/2026-10-18/p20-hex-source-review.json`.

The source is BSD-3-Clause. Exact first-party rights lineage at the same commit:

- [LICENSE](https://go.googlesource.com/go/+/2c31595783c74564c1299a5401dc6e606f379cc7/LICENSE),
  SHA-256 `911f8f5782931320f5b8d1160a76365b83aea6447ee6c04fa6d5591467db9dad`.
- [PATENTS](https://go.googlesource.com/go/+/2c31595783c74564c1299a5401dc6e606f379cc7/PATENTS),
  SHA-256 `96f408bfae65bf137fc2525d3ecb030271c50c1e90799f87abf8846d8dd505cc`.

Locally authored task content follows the project's Apache-2.0 license.
Its content digest and eventual release binding are recorded in provenance.yaml;
no release commit is invented here. No reference-only source contributed code,
tests, text or solutions.
