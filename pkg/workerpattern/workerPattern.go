package util

import (
	"fmt"
)

func StartDispatcher(nrWorkers int) chan WorkRequest {
	var WorkQueue = make(chan WorkRequest, 100)
	WorkerQueue := make(chan chan WorkRequest, nrWorkers)

	for i := 0; i < nrWorkers; i++ {
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()

	return WorkQueue
}

type WorkRequest struct {
	Work      func(...interface{}) string
	Resources []interface{}
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work
			select {
			case work := <-w.Work:
				fmt.Printf("Worker %d: %s\n", w.ID, work.Work(work.Resources...))
			case <-w.QuitChan:
				fmt.Printf("Worker %d: quitting...\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}
	return worker
}
