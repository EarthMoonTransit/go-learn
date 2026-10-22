package subject

type State struct {
	Count  int
	Active bool
	Name   string
	Values []int
	Labels map[string]int
	Target *int
}

func ResetSelected(states []State, selected []bool) int {
	count := 0
	for i, yes := range selected {
		if yes {
			old := states[i].Target
			states[i] = State{Target: old}
			count++
		}
	}
	return count
}
