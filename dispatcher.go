package main

import "fmt"

var WorkerQueue chan chan Task
func StartDispatcher () {
	var numOfWorkers = 10
	 WorkerQueue = make(chan chan Task, numOfWorkers )
	for i:=0;i<numOfWorkers;i++ {
		fmt.Println("Starting Worker ",i+1)
		worker := NewWorker(i,WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Recieved Work Request")
				go func() {
					worker:= <-WorkerQueue

					worker <-work
				}()

			}
		}
	}()
}
