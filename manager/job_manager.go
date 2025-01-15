package manager

import "sync"

// JobManager coordinates and tracks the completion of ETL processes
type JobManager struct {
	wg       sync.WaitGroup
	doneChan chan struct{}
}

func NewJobManager() *JobManager {
	return &JobManager{
		doneChan: make(chan struct{}),
	}
}

func (jm *JobManager) AddWorker() {
	jm.wg.Add(1)
}

func (jm *JobManager) WorkerDone() {
	jm.wg.Done()
}

func (jm *JobManager) WaitForCompletion() {
	jm.wg.Wait()
	close(jm.doneChan)
}
