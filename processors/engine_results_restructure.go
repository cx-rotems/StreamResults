package processors

import (
	"fmt"
	"github.com/cx-rotems/StreamResults/types"
	"time"
)

type EngineResultsRestructure struct {
	resultChan     chan types.Result
	enrichmentChan chan types.Result
}

func NewEngineResultsRestructure(resultChan, enrichmentChan chan types.Result) *EngineResultsRestructure {
	return &EngineResultsRestructure{resultChan: resultChan, enrichmentChan: enrichmentChan}
}

func (er *EngineResultsRestructure) Start() {
	defer close(er.enrichmentChan)

	for result := range er.resultChan {
		result.CvssScores = fmt.Sprintf("%d", result.ResultID*10)
		time.Sleep(70 * time.Millisecond) // simulate restructure
		//fmt.Printf("EngineResultsRestructure: Restructuring result for result ID %d and job ID  %d\n", result.ResultID, result.JobID)
		er.enrichmentChan <- result
	}
}
