package utils

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status int
	Err    error
}

// func CheckURLs(urls []string) []Result {
// 	results := make([]Result, len(urls))
// 	wg := sync.WaitGroup{}
// 	wg.Add(len(urls))
// 	var mu sync.Mutex
// 	if len(urls) == 0 {
// 		return results
// 	}
// 	c := http.Client{
// 		Timeout: 10 * time.Second,
// 	}
// 	for i, url := range urls {
// 		go func(i int, u string) {
// 			defer wg.Done()
// 			res, err := c.Get(u)
// 			if err != nil {
// 				mu.Lock()
// 				results[i] = Result{URL: u, Status: res.StatusCode, Err: err}
// 				mu.Unlock()
// 			} else {
// 				res.Body.Close()
// 				mu.Lock()
// 				results[i] = Result{URL: u, Status: res.StatusCode, Err: nil}
// 				mu.Unlock()
// 			}
// 		}(i, url)
// 	}
// 	wg.Wait()
// 	return results
// }

type ExtendedResults struct {
	Index  int
	Result Result
}

func CheckURLs(urls []string) []Result {
	results := make([]Result, len(urls))
	wg := sync.WaitGroup{}
	if len(urls) == 0 {
		return []Result{}
	}

	resultsChan := make(chan ExtendedResults, len(urls))
	c := http.Client{
		Timeout: 10 * time.Second,
	}

	for i, url := range urls {
		wg.Add(1)
		go func(i int, u string) {
			defer wg.Done()
			res, err := c.Get(u)
			if err != nil {
				resultsChan <- ExtendedResults{Index: i, Result: Result{URL: u, Status: 0, Err: err}}
			} else {
				res.Body.Close()
				resultsChan <- ExtendedResults{Index: i, Result: Result{URL: u, Status: res.StatusCode, Err: nil}}
			}
		}(i, url)
	}
	wg.Wait()
	close(resultsChan)

	go func() {
		defer wg.Done()
		for result := range resultsChan {
			results[result.Index] = result.Result
		}
	}()
	wg.Add(1)
	wg.Wait()
	return results
}
