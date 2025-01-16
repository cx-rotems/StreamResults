package processors

import (
	"fmt"
	"time"
	"github.com/cx-rotems/StreamResults/types"
)

const transactionSize = 4

type ResultLoader struct {
	loaderChan chan types.Result
	callback   func(int)
}

func NewResultLoader(loaderChan chan types.Result, callback func(int)) *ResultLoader {
	return &ResultLoader{loaderChan: loaderChan, callback: callback}
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
        rl.callback(jobID)
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
