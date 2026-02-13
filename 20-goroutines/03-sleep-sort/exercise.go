package sleepSort

import (
	"time"

	"golang.org/x/sync/errgroup"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// reverseSleepSort returns the input uint-slice sorted in the reverse order.
func reverseSleepSort(input []uint) []uint {
	if len(input) <= 0 {
		return input
	}

	reverseSorted := []uint{}

	results := make(chan uint, len(input))

	var eg errgroup.Group

	for w := 0; w < len(input); w++ {
		eg.Go(func() error {
			inputNumber := input[w]
			delay := time.Duration(500-inputNumber*10) * time.Millisecond
			time.Sleep(delay)
			results <- inputNumber

			return nil
		})
	}

	go func() {
		_ = eg.Wait()
		close(results)
	}()

	for individualNumber := range results {
		reverseSorted = append(reverseSorted, individualNumber)
	}

	return reverseSorted
}
