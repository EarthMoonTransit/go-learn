package subject

var saved Job

func StartJob(job Job) Handle {
	if saved.Run == nil {
		saved = job
	}
	go saved.Run()
	return saved.Handle
}
