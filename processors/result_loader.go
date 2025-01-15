package processors

import (
	"fmt"
	"time"

	"github.com/cx-rotems/StreamResults/manager"
	"github.com/cx-rotems/StreamResults/types"
)

type ResultLoader struct {
	loaderChan chan types.Result
	jobManager *manager.JobManager
}

func NewResultLoader(loaderChan chan types.Result, jm *manager.JobManager) *ResultLoader {
	return &ResultLoader{loaderChan: loaderChan, jobManager: jm}
}

func (rl *ResultLoader) Start() {
	defer rl.jobManager.WorkerDone()

	batch := make([]types.Result, 0, 4) // Pre-allocate slice with capacity of 4

	for result := range rl.loaderChan {
		batch = append(batch, result)

		// Process batch when it reaches size 4 or when the channel is closed
		if len(batch) == 4 {
			processBatch(batch)
			batch = batch[:0] // Clear the batch while keeping capacity
		}
	}

	// Process any remaining results
	if len(batch) > 0 {
		processBatch(batch)
	}
}

var batchCounter int

func processBatch(batch []types.Result) {
	batchCounter++
	fmt.Printf("\nResultLoader: Saving batch #%d (%d results)\n", batchCounter, len(batch))
	fmt.Println("Results in this batch:")
	for i, result := range batch {
		fmt.Printf("  [%d] Result ID: %d, Job ID: %d\n",
			i+1, result.ResultID, result.JobID)
	}
	time.Sleep(30 * time.Millisecond)
}
