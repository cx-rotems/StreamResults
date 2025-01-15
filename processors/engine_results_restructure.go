package processors

import (
	"fmt"
	"github.com/cx-rotems/StreamResults/manager"
	"github.com/cx-rotems/StreamResults/types"
	"time"
)

type EngineResultsRestructure struct {
	resultChan     chan types.Result
	enrichmentChan chan types.Result
	jobManager     *manager.JobManager
}

func NewEngineResultsRestructure(resultChan, enrichmentChan chan types.Result, jm *manager.JobManager) *EngineResultsRestructure {
	return &EngineResultsRestructure{resultChan: resultChan, enrichmentChan: enrichmentChan, jobManager: jm}
}

func (er *EngineResultsRestructure) Start() {
	defer er.jobManager.WorkerDone()

	for result := range er.resultChan {
		result.CvssScores = fmt.Sprintf("%d", result.ResultID*10)
		fmt.Printf("EngineResultsRestructure: Restructuring result for result ID %d and job ID  %d\n", result.ResultID, result.JobID)
		time.Sleep(70 * time.Millisecond)
		er.enrichmentChan <- result
	}
	close(er.enrichmentChan)
}
