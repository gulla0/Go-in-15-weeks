package utils

import (
	"fmt"
	"sync"
)

type Job struct {
	ID      int
	Payload int
}

type Out struct {
	ID    int
	Value int
	Err   error
}

// func RunWorkers(jobs []Job, workerCount int) []Out {
// 	if len(jobs) == 0 {
// 		return []Out{}
// 	}

// 	jobsChan := make(chan Job)
// 	mu := sync.Mutex{}
// 	o := make([]Out, 0, len(jobs))
// 	w := sync.WaitGroup{}
// 	w.Add(workerCount + 1)

// 	go func() {
// 		defer w.Done()
// 		for _, job := range jobs {
// 			jobsChan <- job
// 		}
// 		close(jobsChan)
// 	}()

// 	if workerCount > len(jobs) {
// 		workerCount = len(jobs)
// 	}

// 	for range workerCount {
// 		go func() {
// 			defer w.Done()
// 			for job := range jobsChan {
// 				if job.Payload < 0 {
// 					mu.Lock()
// 					o = append(o, Out{ID: job.ID, Value: 0, Err: fmt.Errorf("payload is negative")})
// 					mu.Unlock()
// 				} else {
// 					mu.Lock()
// 					o = append(o, Out{ID: job.ID, Value: job.Payload * 2, Err: nil})
// 					mu.Unlock()
// 				}
// 			}
// 		}()
// 	}
// 	w.Wait()
// 	return o
// }

func RunWorkers(jobs []Job, workerCount int) []Out {
	if len(jobs) == 0 {
		return []Out{}
	}

	jobsChan := make(chan Job)
	resultsChan := make(chan Out)
	o := make([]Out, 0, len(jobs))
	w := sync.WaitGroup{}
	if workerCount > len(jobs) {
		workerCount = len(jobs)
	}
	w.Add(workerCount + 1)

	go func() {
		defer w.Done()
		for _, job := range jobs {
			jobsChan <- job
		}
		close(jobsChan)
	}()

	for range workerCount {
		go func() {
			defer w.Done()
			for job := range jobsChan {
				if job.Payload < 0 {
					resultsChan <- Out{ID: job.ID, Value: 0, Err: fmt.Errorf("payload is negative")}
				} else {
					resultsChan <- Out{ID: job.ID, Value: job.Payload * 2, Err: nil}
				}
			}
		}()
	}

	go func() {
		w.Wait()
		close(resultsChan)
	}()

	for range len(jobs) {
		result := <-resultsChan
		o = append(o, result)
	}
	return o
}
