package subject

func BatchCost(count int) uint { return uint(count/10*95000 + count%10*10000) }
