package pipeline

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func generator(nums []int) <-chan int {
	results := make(chan int)

	go func() {
		for _, n := range nums {
			results <- n
		}
		close(results)
	}()

	return results
}

func adder(in <-chan int) <-chan float32 {
	results := make(chan float32)

	go func() {
		for n := range in {
			results <- float32(n + 2)
		}
		close(results)
	}()

	return results
}

func collector(in <-chan float32) []float32 {
	collected := []float32{}

	for n := range in {
		collected = append(collected, n)
	}

	return collected
}
