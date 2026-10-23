package subject

func Pace(input <-chan int, release <-chan struct{}) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		var values []int
		for value := range input {
			values = append(values, value)
		}
		for i := len(values) - 1; i >= 0; i-- {
			<-release
			output <- values[i]
		}
	}()
	return output
}
