package processors

import (
	"fmt"
	"github.com/cx-rotems/StreamResults/manager"
	"github.com/cx-rotems/StreamResults/types"
	"time"
)

type MinioExtractor struct {
	minioChan  chan types.Job
	resultChan chan types.Result
	jobManager *manager.JobManager
}

func NewMinioExtractor(minioChan chan types.Job, resultChan chan types.Result, jm *manager.JobManager) *MinioExtractor {
	return &MinioExtractor{minioChan: minioChan, resultChan: resultChan, jobManager: jm}
}

func (me *MinioExtractor) Start() {
	defer me.jobManager.WorkerDone()

	for job := range me.minioChan {
		fmt.Printf("MinioExtractor: Extracting data for job ID %d\n", job.ID)
		for i := 1; i < 10; i++ {
			me.resultChan <- types.Result{ResultID: i, JobID: job.ID}
		}
		time.Sleep(100 * time.Millisecond)
		
	}
	close(me.resultChan)
}
