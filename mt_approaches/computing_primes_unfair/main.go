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

func processBatch(wg *sync.WaitGroup, idx, start, end int) {
	defer wg.Done()
	startTime := time.Now()
	for num := start; num <= end; num++ {
		countPrime(num)
	}
	log.Printf("Thread: %d took time: %f", idx, time.Since(startTime).Seconds())
}

func main() {
	startTime := time.Now()
	batchSize := MAX_INT / CONCURRENCY
	startNum, idx := 1, 0
	var wg sync.WaitGroup
	for idx = 0; idx < CONCURRENCY-1; idx++ {
		wg.Add(1)
		go processBatch(&wg, idx, startNum, startNum+batchSize-1)
		startNum += batchSize
	}

	wg.Add(1)
	go processBatch(&wg, idx, startNum, startNum+batchSize)
	wg.Wait()
	log.Printf("Total primes: %d\nTotal time required: %f", totalPrimes, time.Since(startTime).Seconds())
}
