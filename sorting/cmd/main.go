package main

import (
	"fmt"
	"github.com/AlexMykhailov1/cautious-enigma/sorting/sorts"
	"math/rand"
	"sync"
	"time"
)

const n = 100_000

func main() {
	// Generate a random array of length n
	randInts := rand.Perm(n)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		ints := make([]int, n)
		copy(ints, randInts)

		start := time.Now()

		_ = sorts.Simple(ints)

		elapsed := time.Since(start)

		fmt.Printf("Sorted using simple sort. Time elapsed: %v\n", elapsed)
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		ints := make([]int, n)
		copy(ints, randInts)

		start := time.Now()

		_ = sorts.Bubble(ints)

		elapsed := time.Since(start)

		fmt.Printf("Sorted using bubble sort. Time elapsed: %v\n", elapsed)
	}(wg)

	wg.Wait()
}

// Notes:
// 1. Use variable for elapsed time to mark the time right after the sort.
// 2. Copy slice to not use the same underlying array for different sorts.
