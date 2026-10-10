package learner

// ADR 0014: this exact public-test correction leaves every exported workspace
// file and editable flag unchanged. Keep the exception directional and bounded;
// unknown revisions still require explicit reconciliation.
const (
	boundedBeforeNilCheck      = "0e8e493addf0119352a834748e66ea26da49ef61b44db5a4e39f2bbf1212f963"
	boundedWithNilCheck        = "960ba2896fb8d14d5911f0e3e0564aea05f5b22b129c5cb5b5c54c56df1ddbf8"
	bindingsBeforePublicChecks = "f427ca2b6b99603eab0f689a70e9955066c974f5b91d1521b795e5a14292947b"
	bindingsWithPublicChecks   = "f96c2cdcdd423a608cc7c02214e595bc02fdd3aa00ce9fc03478df2ef2afe3d5"
)

func compatibleContentRevision(taskID, stored, current string) bool {
	if stored == current {
		return true
	}
	if taskID == "bounded-slice-view-001" {
		return stored == boundedBeforeNilCheck && current == boundedWithNilCheck
	}
	return taskID == "literal-bindings-001" &&
		stored == bindingsBeforePublicChecks && current == bindingsWithPublicChecks
}
