package learner

// ADRs 0014, 0026 and 0027: exact, directional compatibility pairs preserve existing
// learner work. Unknown revisions still require explicit reconciliation.
const (
	boundedBeforeNilCheck       = "0e8e493addf0119352a834748e66ea26da49ef61b44db5a4e39f2bbf1212f963"
	boundedWithNilCheck         = "960ba2896fb8d14d5911f0e3e0564aea05f5b22b129c5cb5b5c54c56df1ddbf8"
	bindingsBeforePublicChecks  = "f427ca2b6b99603eab0f689a70e9955066c974f5b91d1521b795e5a14292947b"
	bindingsWithPublicChecks    = "f96c2cdcdd423a608cc7c02214e595bc02fdd3aa00ce9fc03478df2ef2afe3d5"
	clampBeforeTestingBridge    = "ed280dce8f21e6154adeb90a368972e45589af8788ddd6eedf030ce4a65b630b"
	clampWithTestingBridge      = "8f1012bd993cac539813d6658a24ef7d12d5b887b52a3165835172eedc3e08c8"
	prependBeforeVariadicBridge = "334ffccf8a0b227463a7ef227e02709e0278961f074d32bef639cbcc4aa050f6"
	prependWithVariadicBridge   = "bc2c7c7028ae1e42ea2bf5144995a5f9698e095b81f86fa1e6f8192e5bff98b8"
)

func compatibleContentRevision(taskID, stored, current string) bool {
	if stored == current {
		return true
	}
	if taskID == "bounded-slice-view-001" {
		return stored == boundedBeforeNilCheck && current == boundedWithNilCheck
	}
	if taskID == "clamp-tests-001" {
		return stored == clampBeforeTestingBridge && current == clampWithTestingBridge
	}
	if taskID == "prepend-values-tests-001" {
		return stored == prependBeforeVariadicBridge && current == prependWithVariadicBridge
	}
	return taskID == "literal-bindings-001" &&
		stored == bindingsBeforePublicChecks && current == bindingsWithPublicChecks
}
