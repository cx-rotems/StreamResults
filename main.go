package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/cx-rotems/StreamResults/processors"
	"github.com/cx-rotems/StreamResults/types"
)

const bufferSize = 1000

func main() {
	// Create channels with buffer size to prevent potential deadlocks
	jobChan := make(chan types.Job, bufferSize)
	minioChan := make(chan types.Job, bufferSize)
	resultChan := make(chan types.Result, bufferSize)
	enrichmentChan := make(chan types.Result, bufferSize)
	loaderChan := make(chan types.Result, bufferSize)

	jobCompleted := func(jobID int) {
        fmt.Printf("Job %d completed\n", jobID)
    }


	processes := []processors.ETLProcess{
		processors.NewJobReceiver(jobChan, minioChan),
		processors.NewMinioExtractor(minioChan, resultChan),
		processors.NewEngineResultsRestructure(resultChan, enrichmentChan),
		processors.NewResultEnrichment(enrichmentChan, loaderChan),
		processors.NewResultLoader(loaderChan, jobCompleted),
	}

	// Start all processes
	for _, process := range processes {
		go process.Start()
	}
	

	go func() {
		for i := 1; i <= 3; i++ {
			jobChan <- types.Job{ID: i}
		}
		close(jobChan)
	}()

	// Create a channel to handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for either a signal or keep running
	select {
	case sig := <-sigChan:
		println("Received signal:", sig)
		// Add any cleanup code here if needed
	}
}
