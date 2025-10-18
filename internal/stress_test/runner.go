package stress_test

import (
	"desafio-goexpert-2/pkg/random_stress_utils"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type HttpStatusCounter struct {
	counts map[int]int
	mutex  sync.Mutex
}

func (c *HttpStatusCounter) Increment(status int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.counts[status]++
}

type StressRunner struct {
	Url         string
	Requests    int
	Concurrency int
}

func NewStressRunner(url string, requests int, concurrency int) *StressRunner {
	return &StressRunner{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}

func (r *StressRunner) Run() {

	wg := &sync.WaitGroup{}
	httpStatusCounter := &HttpStatusCounter{
		counts: make(map[int]int),
		mutex:  sync.Mutex{},
	}
	start := time.Now()
	for range r.Requests {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//httpStatusCounter.Increment(doRequest(r.Url))
			httpStatusCounter.Increment(doDummyRequest())

		}()
	}
	wg.Wait()
	end := time.Now()

	fmt.Printf("\n-----RESULTS-----\n")
	printDuration(start, end)
	printCounts(httpStatusCounter.counts)
}

func doRequest(url string) int {

	response, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer response.Body.Close()
	return response.StatusCode
}

func doDummyRequest() int {
	random_stress_utils.RandomSleep(100*time.Millisecond, 1000*time.Millisecond)
	return random_stress_utils.RandomHttpStatus()
}

func printCounts(counts map[int]int) {

	for k, v := range counts {
		fmt.Printf("HTTP %d - %d\n", k, v)
	}
}

func printDuration(start, end time.Time) {
	fmt.Printf("Duration: %s\n", end.Sub(start))
}
