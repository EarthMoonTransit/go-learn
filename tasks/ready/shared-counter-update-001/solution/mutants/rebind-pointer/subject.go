package subject

func AddCounter(counter *int, delta int) { n := *counter + delta; counter = &n }
