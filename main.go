package main

import (
	"os"
	"os/signal"
	"syscall"
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
		go process.Start()
	}
	

	go func() {
		for i := 1; i <= 3; i++ {
			jobManager.AddJob(i)
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
