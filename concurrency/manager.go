package concurrency

type Manager struct {
	MaxWorkers int
	WorkerPool chan chan job
}

func NewManager(maxWorkers int) *Manager {
	pool := make(chan chan job, maxWorkers)
	return &Manager{
		MaxWorkers: maxWorkers,
		WorkerPool: pool,
	}
}

func (manager *Manager) Do(seeder func()) {
	for i := 0; i < manager.MaxWorkers; i++ {
		w := newWorker(i, manager.WorkerPool)
		manager.WorkerPool <- w.jobChannel
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
			jobChannel := <-manager.WorkerPool
			jobChannel <- job
		}
	}
}
