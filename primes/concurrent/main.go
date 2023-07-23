package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type PrimeResult struct {
	n     int
	prime bool
}

type PrimeTime struct {
	PrimeResult
	elapsed time.Duration
}

var Sample = []PrimeResult{
	{2, true},
	{142702110479723, true},
	{299593572317531, true},
	{3333333333333301, true},
	{3333333333333333, false},
	{4444444444444423, true},
	{4444444444444444, false},
	{4444444488888889, false},
	{5555553133149889, false},
	{5555555555555503, true},
	{5555555555555555, false},
	{6666666666666666, false},
	{6666666666666719, true},
	{6666667141414921, false},
	{7777777536340681, false},
	{7777777777777753, true},
	{7777777777777777, false},
	{9999999999999917, true},
	{9999999999999999, false},
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	root := int(math.Sqrt(float64(n)))
	for i := 3; i <= root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func check(n int) PrimeTime {
	t0 := time.Now()
	res := PrimeResult{n, isPrime(n)}
	elapsed := time.Since(t0)
	return PrimeTime{res, elapsed}
}

func worker(jobs <-chan int, results chan<- PrimeTime, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		results <- check(n)
	}
}

func startJobs(workers int, jobs chan int, results chan PrimeTime) {
	var wg sync.WaitGroup
	wg.Add(workers)

	for _, n := range Sample {
		jobs <- n.n
	}
	close(jobs)

	for i := 0; i < workers; i++ {
		go worker(jobs, results, &wg)
	}
	wg.Wait()
	close(results)
}

func report(workers int, results chan PrimeTime) int {
	checked := 0
	for result := range results {
		if result.n != 0 {
			checked++
			label := " "
			if result.prime {
				label = "P"
			}
			fmt.Printf("%16d  %s %9.6fs\n", result.n, label, result.elapsed.Seconds())
		}
	}
	return checked
}

func main() {
	var workers int
	if len(os.Args) > 1 {
		var err error
		workers, err = strconv.Atoi(os.Args[1])
		if err != nil || workers < 1 {
			fmt.Printf("Invalid number of workers: %v\n", os.Args[1])
			os.Exit(1)
		}
	} else {
		workers = runtime.NumCPU()
	}

	fmt.Printf("Checking %d numbers with %d coroutines:\n", len(Sample), workers)
	t0 := time.Now()
	jobs := make(chan int, len(Sample))
	timings := make(chan PrimeTime, len(Sample))
	go startJobs(workers, jobs, timings)
	checked := report(workers, timings)
	elapsed := time.Since(t0)
	fmt.Printf("%d checks in %.2fs\n", checked, elapsed.Seconds())
}
