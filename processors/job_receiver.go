package processors

import (
	//"fmt"
	"github.com/cx-rotems/StreamResults/types"
	"time"
)

// JobReceiver simulates getting jobs and sending them to MinioExtractor
type JobReceiver struct {
	jobChan    chan types.Job
	minioChan  chan types.Job
}

func NewJobReceiver(jobChan, minioChan chan types.Job) *JobReceiver {
	return &JobReceiver{jobChan: jobChan, minioChan: minioChan}
}

func (jr *JobReceiver) Start() {
	defer close(jr.minioChan)

	for job := range jr.jobChan {
		//fmt.Printf("JobReceiver: Processing job ID %d\n", job.ID)
		time.Sleep(50 * time.Millisecond)
		jr.minioChan <- job
	}
}
