package wordcount

import (
	"runtime"
	"strings"

	"golang.org/x/sync/errgroup"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func CountWords(texts []string) map[string]int {
	finalCounts := make(map[string]int)

	if len(texts) == 0 {
		return finalCounts
	}

	jobs := make(chan string, len(texts))

	for _, text := range texts {
		jobs <- text
	}

	close(jobs)

	numWorkers := runtime.NumCPU()
	results := make(chan map[string]int, numWorkers)

	var eg errgroup.Group

	for w := 0; w < numWorkers; w++ {
		eg.Go(func() error {
			localMap := make(map[string]int)

			for text := range jobs {
				words := strings.Fields(text)

				for _, word := range words {
					localMap[word]++
				}
			}

			results <- localMap

			return nil
		})
	}

	go func() {
		_ = eg.Wait()
		close(results)
	}()

	for individualMap := range results {
		for word, count := range individualMap {
			finalCounts[word] += count
		}
	}

	return finalCounts
}
