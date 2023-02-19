package concurrency

type worker struct {
	id         int
	jobChannel chan job
	WorkerPool chan chan job
	Stop       chan bool
}

func newWorker(id int, workerPool chan chan job) *worker {
	return &worker{
		id:         id,
		WorkerPool: workerPool,
		jobChannel: make(chan job),
		Stop:       make(chan bool),
	}
}

func (w *worker) do() {
	for {
		select {
		case job := <-w.jobChannel:
			job.Start()
			w.WorkerPool <- w.jobChannel
		case <-w.Stop:
			return
		}
	}
}

func (w *worker) stop() {
	go func() {
		w.Stop <- true
	}()
}
