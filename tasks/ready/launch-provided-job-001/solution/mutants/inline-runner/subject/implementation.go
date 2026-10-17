package subject

func StartJob(job Job) Handle {
	job.Run()
	return job.Handle
}
