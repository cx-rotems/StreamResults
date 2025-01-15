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
	defer rl.jobManager.WorkerDone()

	transaction := make([]types.Result, 0, transactionSize)

	for result := range rl.loaderChan {
		transaction = append(transaction, result)

		if len(transaction) == transactionSize {
			processTransaction(transaction)
			transaction = transaction[:0]
		}
	}

	if len(transaction) > 0 {
		processTransaction(transaction)
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
