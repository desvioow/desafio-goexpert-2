package stress_test

import (
	"context"
	"desafio-goexpert-2/pkg/random_stress_utils"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/leaanthony/spinner"
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
	start := time.Now()
	spinner := spinner.New()
	spinner.Start("Stress testing...")

	httpStatusCounter := HttpStatusCounter{counts: make(map[int]int)}
	workers := r.determineWorkers()
	jobs := make(chan int)
	httpClient := http.Client{}

	requestsWg := &sync.WaitGroup{}
	requestsWg.Add(r.Requests)
	workersWg := &sync.WaitGroup{}

	r.startWorkerPool(workers, &httpClient, jobs, requestsWg, workersWg, &httpStatusCounter)
	r.queueRequestJobs(jobs)
	r.waitForCompletion(requestsWg, workersWg)

	spinner.Success("Stress testing completed!")
	r.reportResults(start, time.Now(), httpStatusCounter.counts)
}

func (r *StressRunner) startWorkerPool(workers int, httpClient *http.Client, jobs <-chan int, requestsWg, workersWg *sync.WaitGroup, counter *HttpStatusCounter) {
	for w := 0; w < workers; w++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for range jobs {
				//status := doRequestWithClient(httpClient, r.Url)
				status := doDummyRequest()
				counter.Increment(status)
				requestsWg.Done()
			}
		}()
	}
}

func (r *StressRunner) determineWorkers() int {
	workers := r.Concurrency
	if workers > r.Requests {
		workers = r.Requests
	}
	if workers <= 0 {
		workers = 1
	}
	return workers
}

func (r *StressRunner) queueRequestJobs(jobs chan<- int) {
	for i := 0; i < r.Requests; i++ {
		jobs <- i
	}
	close(jobs)
}

func (r *StressRunner) waitForCompletion(requestsWg, workersWg *sync.WaitGroup) {
	requestsWg.Wait()
	workersWg.Wait()
}

func doDummyRequest() int {
	random_stress_utils.RandomSleep(10*time.Millisecond, 500*time.Millisecond)
	return random_stress_utils.RandomHttpStatus()
}

func doRequestWithClient(client *http.Client, url string) int {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0
	}

	response, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer response.Body.Close()
	return response.StatusCode
}

func (r *StressRunner) reportResults(start, end time.Time, counts map[int]int) {
	fmt.Printf("\n-----RESULTS-----\n")
	printDuration(start, end)
	printCounts(counts, r.Requests)
}

func printCounts(counts map[int]int, total int) {
	fmt.Printf("Total requests: %d\n", total)
	fmt.Printf("HTTP 200 - %d\n", counts[http.StatusOK])

	for k, v := range counts {
		if k != http.StatusOK {
			fmt.Printf("HTTP %d - %d\n", k, v)
		}
	}
}

func printDuration(start, end time.Time) {
	fmt.Printf("Duration: %s\n", end.Sub(start))
}
