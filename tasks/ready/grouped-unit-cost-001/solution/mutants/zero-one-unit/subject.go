package subject

func BatchCost(count int) uint {
	if count == 0 {
		return 10000
	}
	return uint(count/10*95000 + count%10*10000)
}
