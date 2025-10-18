package stress_test

import (
	"desafio-goexpert-2/pkg/random_sleep"
	"fmt"
	"sync"
	"time"
)

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
	for i := range r.Requests {
		wg.Add(1)
		go func() {
			// dummy request behavior for testing
			random_sleep.RandomSleep(100*time.Millisecond, 1000*time.Millisecond)
			fmt.Printf("%d - request done\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
}
