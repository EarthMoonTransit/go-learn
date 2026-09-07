package learner

// ADR 0014: this exact public-test correction leaves every exported workspace
// file and editable flag unchanged. Keep the exception directional and bounded;
// unknown revisions still require explicit reconciliation.
const (
	bindingsBeforePublicChecks = "f427ca2b6b99603eab0f689a70e9955066c974f5b91d1521b795e5a14292947b"
	bindingsWithPublicChecks   = "f96c2cdcdd423a608cc7c02214e595bc02fdd3aa00ce9fc03478df2ef2afe3d5"
)

func compatibleContentRevision(taskID, stored, current string) bool {
	if stored == current {
		return true
	}
	return taskID == "literal-bindings-001" &&
		stored == bindingsBeforePublicChecks && current == bindingsWithPublicChecks
}
