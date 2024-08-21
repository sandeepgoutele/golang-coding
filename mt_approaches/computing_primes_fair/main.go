package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var (
	MAX_INT           = 100000000
	totalPrimes int32 = 0
	CONCURRENCY       = 10
	currentNum  int32 = 0
)

func countPrime(num int) {
	if num <= 1 || (num%2 == 0 || num%3 == 0) {
		return
	}

	if num <= 3 {
		atomic.AddInt32(&totalPrimes, 1)
		return
	}

	idx := 5
	for idx*idx <= num {
		if num%idx == 0 || num%(idx+2) == 0 {
			return
		}
		idx += 6
	}

	atomic.AddInt32(&totalPrimes, 1)
}

func processOptimally(wg *sync.WaitGroup, idx int) {
	defer wg.Done()
	startTime := time.Now()
	for {
		num := atomic.AddInt32(&currentNum, 1)
		if num > int32(MAX_INT) {
			break
		}
		countPrime(int(num))
	}
	log.Printf("Thread: %d took time: %f", idx, time.Since(startTime).Seconds())
}

func main() {
	startTime := time.Now()
	idx := 1
	var wg sync.WaitGroup
	for idx = 0; idx < CONCURRENCY; idx++ {
		wg.Add(1)
		go processOptimally(&wg, idx)
	}

	wg.Wait()
	log.Printf("Total primes: %d\nTotal time required: %f", totalPrimes, time.Since(startTime).Seconds())
}
