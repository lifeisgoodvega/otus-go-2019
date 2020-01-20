package parallel

import (
	"errors"
	"sync"
)

func worker(taskChannel <-chan func() error, completionChannel chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task := <-taskChannel
		if task == nil {
			break
		}
		completionChannel <- task()
	}
}

// Run func array in parallel. N - workers number, M - maximum error count.
func Run(task []func() error, N int, M int) error {
	taskChannel := make(chan func() error)
	completionChannel := make(chan error)
	errorNum := 0
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		go worker(taskChannel, completionChannel, &wg)
		wg.Add(1)
	}

	pendingTasks := 0
	tIndex := 0
	for tIndex < len(task) {
		select {
		case err := <-completionChannel:
			pendingTasks--
			if err != nil {
				errorNum++
				if errorNum == M {
					tIndex = len(task) // Stop pushing new tasks
				}
			}

		case taskChannel <- task[tIndex]:
			pendingTasks++
			tIndex++
		}
	}

	for pendingTasks > 0 {
		err := <-completionChannel
		if err != nil {
			errorNum++
		}
		pendingTasks--
	}

	if errorNum >= M {
		return errors.New("Limit of errors was exceeded")
	}

	close(taskChannel)
	close(completionChannel)

	wg.Wait()

	return nil
}
