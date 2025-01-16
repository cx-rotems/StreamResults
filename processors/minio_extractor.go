package processors

import (
	//"fmt"
	"github.com/cx-rotems/StreamResults/types"
	"time"
)

type MinioExtractor struct {
	minioChan  chan types.Job
	resultChan chan types.Result

}

func NewMinioExtractor(minioChan chan types.Job, resultChan chan types.Result) *MinioExtractor {
	return &MinioExtractor{minioChan: minioChan, resultChan: resultChan}
}

func (me *MinioExtractor) Start() {
	defer close(me.resultChan)

	for job := range me.minioChan {
	//	fmt.Printf("MinioExtractor: Extracting data for job ID %d\n", job.ID)
		for i := 1; i <= 50; i++ {
			time.Sleep(100 * time.Millisecond) // simulate download from Minio
			me.resultChan <- types.Result{ResultID: i, JobID: job.ID}
		}
		
	}
}
