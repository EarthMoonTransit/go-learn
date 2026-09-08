package subject

const targetMinutes = 40

func RemainingMinutes(elapsed int) int {
	if elapsed == 0 {
		return 0
	}
	return targetMinutes - elapsed
}
