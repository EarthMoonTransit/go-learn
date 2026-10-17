package subject_test

import (
	"task.local/launch-provided-job-001/subject"
	"testing"
	"testing/synctest"
)

func startHeld(t *testing.T, job subject.Job) subject.Handle {
	t.Helper()
	returned := make(chan subject.Handle, 1)
	go func() { returned <- subject.StartJob(job) }()
	synctest.Wait()
	select {
	case handle := <-returned:
		return handle
	default:
		t.Error("StartJob blocked with held work, want return before Release")
		job.Handle.Release()
		synctest.Wait()
		select {
		case handle := <-returned:
			return handle
		default:
			t.Fatal("StartJob did not return after Release, want supplied handle")
		}
	}
	return subject.Handle{}
}
func checkJobs(t *testing.T, count int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		jobs := make([]subject.Job, count)
		for i := range jobs {
			jobs[i] = subject.NewJob()
		}
		defer func() {
			for _, job := range jobs {
				job.Handle.Release()
			}
			synctest.Wait()
		}()
		for i, job := range jobs {
			handle := startHeld(t, job)
			if got := handle == job.Handle; !got {
				t.Errorf("job %d returned supplied handle = %t, want true", i, got)
			}
		}
		synctest.Wait()
		for i, job := range jobs {
			if got := job.Handle.Calls(); got != 1 {
				t.Errorf("held job %d calls = %d, want 1", i, got)
			}
			select {
			case <-job.Handle.Done():
				t.Errorf("held job %d completed before Release, want pending", i)
			default:
			}
		}
		for i, job := range jobs {
			job.Handle.Release()
			synctest.Wait()
			select {
			case _, open := <-job.Handle.Done():
				if open {
					t.Errorf("released job %d completion sent a value, want close", i)
				}
			default:
				t.Errorf("released job %d completion = pending, want closed", i)
			}
			if got := job.Handle.Calls(); got != 1 {
				t.Errorf("released job %d calls = %d, want 1", i, got)
			}
			for j := i + 1; j < len(jobs); j++ {
				select {
				case <-jobs[j].Handle.Done():
					t.Errorf("held independent job %d completed with job %d, want pending", j, i)
				default:
				}
			}
		}
	})
}
func TestTwoHeldJobs(t *testing.T) { checkJobs(t, 2) }
