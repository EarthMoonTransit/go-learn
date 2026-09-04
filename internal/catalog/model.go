package catalog

type Candidate struct {
	SchemaVersion          int             `json:"schema_version"`
	SourceID               string          `json:"source_id"`
	SourceCommit           string          `json:"source_commit"`
	SourcePath             string          `json:"source_path"`
	Slug                   string          `json:"slug"`
	ExerciseKind           string          `json:"exercise_kind"`
	UpstreamStatus         string          `json:"upstream_status"`
	UpstreamDifficulty     *int            `json:"upstream_difficulty"`
	UpstreamConcepts       []string        `json:"upstream_concepts"`
	UpstreamPrerequisites  []string        `json:"upstream_prerequisites"`
	UpstreamTopics         []string        `json:"upstream_topics"`
	StatementPaths         []string        `json:"statement_paths"`
	DocumentationPaths     []string        `json:"documentation_paths"`
	StarterFilePaths       []string        `json:"starter_file_paths"`
	PublicTestPaths        []string        `json:"public_test_paths"`
	GeneratedTestDataPaths []string        `json:"generated_test_data_paths"`
	ReferenceSolutionPaths []string        `json:"reference_solution_paths"`
	MetadataPaths          []string        `json:"metadata_paths"`
	AdditionalFilePaths    []string        `json:"additional_file_paths"`
	FileHashes             []FileHash      `json:"file_hashes"`
	License                string          `json:"license"`
	UsageMode              string          `json:"usage_mode"`
	ImportWarnings         []string        `json:"import_warnings"`
	ProvenanceRefs         []ProvenanceRef `json:"provenance_refs"`
}

type FileHash struct {
	Path   string   `json:"path"`
	SHA256 string   `json:"sha256"`
	Roles  []string `json:"roles"`
}

type ProvenanceRef struct {
	FilePath          string `json:"file_path"`
	SourceID          string `json:"source_id"`
	Repository        string `json:"repository"`
	Commit            string `json:"commit"`
	AbbreviatedCommit string `json:"abbreviated_commit"`
	License           string `json:"license"`
	Evidence          string `json:"evidence"`
}
