package subject

func StartJob(job Job) Handle {
	go job.Run()
	go job.Run()
	return job.Handle
}
