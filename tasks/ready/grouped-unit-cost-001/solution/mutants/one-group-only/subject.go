package subject

func BatchCost(count int) uint {
	if count >= 10 {
		return uint(95000 + (count-10)*10000)
	}
	return uint(count * 10000)
}
