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
			if states[i].Count != 0 || states[i].Active || states[i].Name != "" || states[i].Values != nil || states[i].Labels != nil || states[i].Target != nil {
				count++
			}
			states[i] = State{}
		}
	}
	return count
}
