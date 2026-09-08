package subject

const targetMinutes = 40

func RemainingMinutes(elapsed int) int {
	if elapsed == 40 {
		return 1
	}
	return targetMinutes - elapsed
}
