package concurrency

type worker struct {
	id         int
	jobChannel chan job
	WorkerPool chan chan job
}

func newWorker(id int, workerPool chan chan job) *worker {
	return &worker{
		id:         id,
		WorkerPool: workerPool,
		jobChannel: make(chan job),
	}
}

func (w *worker) do() {
	for {
		select {
		case job := <-w.jobChannel:
			job.Start()
			w.WorkerPool <- w.jobChannel
		}
	}
}
