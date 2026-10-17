package subject

func StartJob(job Job) Handle {
	go job.Run()
	return job.Handle
}
