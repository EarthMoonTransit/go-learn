package subject

func StartJob(job Job) Handle {
	job.Handle.Release()
	go job.Run()
	return job.Handle
}
