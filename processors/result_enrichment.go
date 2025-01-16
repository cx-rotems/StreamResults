package processors

import (
	//"fmt"
	"github.com/cx-rotems/StreamResults/types"
	"time"
)

type ResultEnrichment struct {
	enrichmentChan chan types.Result
	loaderChan     chan types.Result
}

func NewResultEnrichment(enrichmentChan, loaderChan chan types.Result) *ResultEnrichment {
	return &ResultEnrichment{enrichmentChan: enrichmentChan, loaderChan: loaderChan}
}

func (re *ResultEnrichment) Start() {
	defer close(re.loaderChan)

	for result := range re.enrichmentChan {
	//	fmt.Printf("ResultEnrichment: Enriching result for result ID %d and job ID  %d\n", result.ResultID, result.JobID) // simulate result enrichment
		time.Sleep(60 * time.Millisecond)
		re.loaderChan <- result
	}
}
