package concurrency

type Manager struct {
	MaxWorkers  int
	workerPool  chan chan job
	stopChannel chan bool
	workers     []*worker
}

func NewManager(maxWorkers int) *Manager {
	pool := make(chan chan job, maxWorkers)
	return &Manager{
		MaxWorkers:  maxWorkers,
		workerPool:  pool,
		stopChannel: make(chan bool),
	}
}

func (manager *Manager) Do(seeder func()) {
	for i := 0; i < manager.MaxWorkers; i++ {
		w := newWorker(i, manager.workerPool)
		manager.workerPool <- w.jobChannel
		manager.workers = append(manager.workers, w)
		go w.do()
	}

	go manager.watch()
	seeder()
}

/*
 * Watching if any jobs are added to jobList.
 */
func (manager *Manager) watch() {
	for {
		select {
		case job := <-JobList:
			jobChannel := <-manager.workerPool
			jobChannel <- job

		case <-manager.stopChannel:
			for _, w := range manager.workers {
				w.stop()
			}
			return
		}
	}
}

func (manager *Manager) Finish() {
	manager.stopChannel <- true
}
