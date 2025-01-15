package main

import (
	"fmt"
	"time"
	"github.com/cx-rotems/StreamResults/manager"
	"github.com/cx-rotems/StreamResults/processors"
	"github.com/cx-rotems/StreamResults/types"
)

func main() {
	// Create channels with buffer size to prevent potential deadlocks
	jobChan := make(chan types.Job, 5)
	minioChan := make(chan types.Job, 5)
	resultChan := make(chan types.Result, 5)
	enrichmentChan := make(chan types.Result, 5)
	loaderChan := make(chan types.Result, 5)

	jobManager := manager.NewJobManager()

	processes := []processors.ETLProcess{
		processors.NewJobReceiver(jobChan, minioChan, jobManager),
		processors.NewMinioExtractor(minioChan, resultChan, jobManager),
		processors.NewEngineResultsRestructure(resultChan, enrichmentChan, jobManager),
		processors.NewResultEnrichment(enrichmentChan, loaderChan, jobManager),
		processors.NewResultLoader(loaderChan, jobManager),
	}

	// Start all processes
	for _, process := range processes {
		jobManager.AddWorker()
		go process.Start()
	}
	
	startTime := time.Now()

	// Send jobs in a separate goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			jobChan <- types.Job{ID: i}
		}
		close(jobChan)
	}()

	// Wait for completion
	jobManager.WaitForCompletion()
	fmt.Println("All jobs completed")

	// Calculate and print the total time taken
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time taken: %v\n", elapsedTime)
}
