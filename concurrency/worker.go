package concurrency

type worker struct {
	id          int
	jobChannel  chan job
	workerPool  chan chan job
	stopChannel chan bool
}

func newWorker(id int, workerPool chan chan job) *worker {
	return &worker{
		id:          id,
		workerPool:  workerPool,
		jobChannel:  make(chan job),
		stopChannel: make(chan bool),
	}
}

func (w *worker) do() {
	for {
		select {
		case job := <-w.jobChannel:
			job.Start()
			w.workerPool <- w.jobChannel
		case <-w.stopChannel:
			return
		}
	}
}

func (w *worker) stop() {
	go func() {
		w.stopChannel <- true
	}()
}
