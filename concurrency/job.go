package concurrency

type job struct {
	action func()
}

func (j *job) Start() {
	j.action()
	WG.Done()
}

func NewJob(action func()) job {
	WG.Add(1)
	return job{
		action,
	}
}
