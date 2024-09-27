package worker_pool_unbuffered

import "fmt"

type Task func() (result interface{}, err error)

type WorkerPool struct {
	taskQueue   chan Task   // unbuffered chan.
	resultChan  chan Result // unbuffered chan.
	workerCount int         // no of workers in the pool.
}

// Worker - wrapping the metadata for the worker.
type Worker struct {
	id         int
	taskQueue  <-chan Task   // Read/Receive channel
	resultChan chan<- Result // Write/Send channel.
}

// Result - final result struct
type Result struct {
	workerID int
	result   interface{}
	err      error
}

func (w *Worker) Start() {
	go func() {
		// Reading the task  from the taskQueue
		// as it is un-buffered - it will take data one by one.
		for task := range w.taskQueue {
			fmt.Println("Going to perform task with ID : ", w.id)
			// executing the task func.
			result, err := task()

			fmt.Println("Executed task with ID : ", w.id)

			// writing the result in wrapper object and returning.
			w.resultChan <- Result{workerID: w.id, result: result, err: err}
		}
	}()
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		taskQueue:   make(chan Task), //unbuffered channel
		resultChan:  make(chan Result),
		workerCount: workerCount,
	}
}

func (w *WorkerPool) Start() {
	for i := 0; i < w.workerCount; i++ {
		// every worker we are passing the same taskQueue and result chan to update on.
		worker := Worker{id: i, taskQueue: w.taskQueue, resultChan: w.resultChan}

		// starting the worker.
		// Note: Below is happening
		// 1. Every worker is trying to read from the same task queue.
		// Listening to read some task and then return the result to result chan.
		worker.Start()
	}
}

func (w *WorkerPool) Submit(task Task) {
	// this will be submitted one by one as there is no buffer.
	// Check for error condition ?
	w.taskQueue <- task
}

func (w *WorkerPool) GetResult() Result {
	// returning the read channel.
	return <-w.resultChan
}
