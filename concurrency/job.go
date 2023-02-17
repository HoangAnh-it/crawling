package concurrency

type job struct {
	action func()
}

func (j *job) Start() {
	j.action()
}

func NewJob(action func()) job {
	return job{
		action,
	}
}

var (
	JobList = make(chan job)
)
