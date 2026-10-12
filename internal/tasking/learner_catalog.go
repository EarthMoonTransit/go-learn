package tasking

import (
	"sync"

	"go-practice-bank/internal/sources"
	"go-practice-bank/internal/triage"
)

// A loader belongs to one catalog read, never to a Service or global cache.
// Shared policy is decoded once; every task still receives all export checks.
// A subsequent read observes edits to policy, schemas and task files.
type learnerPackageLoader struct {
	root       string
	contracts  func() (map[string]triage.SkillContract, error)
	provenance func() provenanceInputs
}

type provenanceInputs struct {
	registry sources.Registry
	identity projectIdentity
	license  outboundLicensePolicy
	problems []string
}

func newLearnerPackageLoader(root string) *learnerPackageLoader {
	contracts := sync.OnceValues(func() (map[string]triage.SkillContract, error) {
		_, contracts, err := triage.LoadSkillContracts(root)
		return contracts, err
	})
	// Keep loading lazy so legacy-only/empty catalogs retain their behavior.
	provenance := sync.OnceValue(func() provenanceInputs {
		registry, err := sources.LoadRegistry(root)
		if err != nil {
			return provenanceInputs{problems: []string{"cannot load source registry: " + err.Error()}}
		}
		identity, err := loadProjectIdentity(root)
		if err != nil {
			return provenanceInputs{problems: []string{err.Error()}}
		}
		license, err := loadOutboundLicensePolicy(root)
		if err != nil {
			return provenanceInputs{problems: []string{err.Error()}}
		}
		return provenanceInputs{registry: registry, identity: identity, license: license}
	})
	return &learnerPackageLoader{root: root, contracts: contracts, provenance: provenance}
}

func (l *learnerPackageLoader) checkProvenance(data packageData) []string {
	if data.ProvenanceVersion != 2 {
		return checkProvenance(l.root, data)
	}
	inputs := l.provenance()
	if len(inputs.problems) != 0 {
		return inputs.problems
	}
	return checkProvenanceV2WithInputs(data, inputs.registry, inputs.identity, inputs.license)
}
