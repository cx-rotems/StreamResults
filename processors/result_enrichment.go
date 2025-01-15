package processors

import (
	//"fmt"
	"github.com/cx-rotems/StreamResults/manager"
	"github.com/cx-rotems/StreamResults/types"
	"time"
)

type ResultEnrichment struct {
	enrichmentChan chan types.Result
	loaderChan     chan types.Result
	jobManager     *manager.JobManager
}

func NewResultEnrichment(enrichmentChan, loaderChan chan types.Result, jm *manager.JobManager) *ResultEnrichment {
	return &ResultEnrichment{enrichmentChan: enrichmentChan, loaderChan: loaderChan, jobManager: jm}
}

func (re *ResultEnrichment) Start() {
	defer re.jobManager.WorkerDone()

	for result := range re.enrichmentChan {
	//	fmt.Printf("ResultEnrichment: Enriching result for result ID %d and job ID  %d\n", result.ResultID, result.JobID) // simulate result enrichment
		time.Sleep(60 * time.Millisecond)
		re.loaderChan <- result
	}
	close(re.loaderChan)
}
