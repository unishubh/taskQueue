package main

import (
	"fmt"
	"time"
)

type Worker struct {
	ID int
	WorkQueue chan chan Task
	Work chan Task
	Quit chan bool
}

func NewWorker (id int, workerQueue chan chan Task) Worker {

	worker := Worker{
		ID:id,
		WorkQueue:workerQueue,
		Work:make(chan Task),
		Quit:make(chan bool),
	}
	return worker
}

func (w *Worker) Start()  {
	go func() {
		for {
			w.WorkQueue <- w.Work
			select {
					case work := <-w.Work :
					fmt.Println("Started doing work", work.name)
						time.Sleep(5*time.Second)
					fmt.Println("Stopped doing work", work.name)
					case <- w.Quit:
						fmt.Println("Work Done")
						return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}