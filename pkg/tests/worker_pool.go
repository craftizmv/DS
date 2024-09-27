package tests

import (
	"errors"
	"sync"
)

// WorkerPool errors, do not change!
var (
	ErrBadParams  = errors.New("bad params")
	ErrBadTask    = errors.New("bad task")
	ErrNotStarted = errors.New("not started")
)

// WorkerPool represents a pool of goroutines.
type WorkerPool struct {
	workerBuffer chan Worker // any number of tasks can be added to this list
	resultChan   chan error
	workerList   []Worker
	mu           sync.Mutex
	poolSize     int
}

type Worker struct {
	Task   Task
	TaskID int
}

// NewWorkerPool creates a new pool with a given size.
func NewWorkerPool(size int) (*WorkerPool, error) {
	if size <= 0 {
		return nil, ErrBadParams
	}
	workerPool := &WorkerPool{
		workerBuffer: make(chan Worker, size), // buffered queue
		workerList:   make([]Worker, 0, size),
		poolSize:     size,
	}
	return workerPool, nil
}

// Task to be computed by the WorkerPool.
type Task func() error

func (wp *WorkerPool) removeTask(taskPos int) error {

	if wp.workerList == nil {
		return ErrNotStarted
	}

	if taskPos < 0 || taskPos >= len(wp.workerList) {
		return ErrBadParams
	}

	wp.mu.Lock()
	defer wp.mu.Unlock()

	for _, w := range wp.workerList {
		if w.TaskID == taskPos {
			wp.workerList = append(wp.workerList[:taskPos], wp.workerList[taskPos+1:]...)
		}
	}

	return nil
}

// AddTask will add a task to the worker pool queue.
func (wp *WorkerPool) AddTask(task Task) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	if task == nil {
		return ErrBadTask
	}

	if wp.workerList == nil {
		return ErrNotStarted
	}

	wp.workerList = append(wp.workerList, Worker{
		Task:   task,
		TaskID: len(wp.workerList) + 1,
	})

	return nil
}

// Results returns channel of non-nil errors.
func (wp *WorkerPool) Results() <-chan error {
	// caller can range over this channel to get the results.
	return wp.resultChan // not implemented yet
}

// Run will start workers(goroutines) for tasks computation.
func (wp *WorkerPool) Run() {
	// not implemented yet
	for worker := range wp.workerBuffer {
		// run as go routine.
		go func(worker Worker) {
			// Run task
			err := worker.Task()
			if err != nil {
				wp.resultChan <- err
			}

			// task is done remove from the queue, in async way.
			//wp.removeTask(worker.TaskID)

		}(worker)
	}
}
