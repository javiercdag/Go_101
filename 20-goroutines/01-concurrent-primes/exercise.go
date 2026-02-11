package concurrentprimes

import (
	"math"
	"runtime"
	"sort"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func GeneratePrimes(n int) []int {
	if n <= 0 {
		return []int{}
	}

	jobs := make(chan int, n)
	results := make(chan int, n)

	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for num := range jobs {
				if isPrime(num) {
					results <- num
				}
			}
		}()
	}

	for i := 2; i <= n; i++ {
		jobs <- i
	}

	close(jobs)

	go func() {
		wg.Wait() // blocking while wg working (jobs open OR workers working)
		close(results)
	}()

	var primes []int

	for p := range results { // blocking while results is open
		primes = append(primes, p)
	}

	sort.Ints(primes)

	return primes
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	// We only need to check up to the square root of the number
	sq := int(math.Sqrt(float64(num)))

	for i := 2; i <= sq; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
