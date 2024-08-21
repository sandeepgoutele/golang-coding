package main

import (
	"log"
	"time"
)

var (
	MAX_INT           = 100000000
	totalPrimes int32 = 0
)

func countPrime(num int) {
	if num <= 1 || (num%2 == 0 || num%3 == 0) {
		return
	}

	if num <= 3 {
		totalPrimes += 1
		return
	}

	idx := 5
	for idx*idx <= num {
		if num%idx == 0 || num%(idx+2) == 0 {
			return
		}
		idx += 6
	}

	totalPrimes += 1
}

func main() {
	startTime := time.Now()
	for num := 1; num <= MAX_INT; num++ {
		countPrime(num)
	}

	log.Printf("Total primes: %d\nTotal time required: %f", totalPrimes, time.Since(startTime).Seconds())
}
