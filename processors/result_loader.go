package processors

import (
	"fmt"
	"time"
	"github.com/cx-rotems/StreamResults/manager"
	"github.com/cx-rotems/StreamResults/types"
)

const transactionSize = 4

type ResultLoader struct {
	loaderChan chan types.Result
	jobManager *manager.JobManager
}

func NewResultLoader(loaderChan chan types.Result, jm *manager.JobManager) *ResultLoader {
	return &ResultLoader{loaderChan: loaderChan, jobManager: jm}
}

func (rl *ResultLoader) Start() {
    transactions := make(map[int][]types.Result)

    for result := range rl.loaderChan {
        jobID := result.JobID
        transactions[jobID] = append(transactions[jobID], result)

        if len(transactions[jobID]) == transactionSize {
            processTransaction(transactions[jobID])
            transactions[jobID] = transactions[jobID][:0]
        }
    }

    for jobID, transaction := range transactions {
        if len(transaction) > 0 {
            processTransaction(transaction)
        }
        rl.jobManager.JobCompleted(jobID)
    }
}

var transactionCounter int

func processTransaction(transaction []types.Result) {
	transactionCounter++
	fmt.Printf("\nResultLoader: Saving transaction #%d (%d results)\n", transactionCounter, len(transaction))
	fmt.Println("Results in this transaction:")
	for i, result := range transaction {
		fmt.Printf("  [%d] Result ID: %d, Job ID: %d\n",
			i+1, result.ResultID, result.JobID)
	}
	time.Sleep(30 * time.Millisecond)
}
