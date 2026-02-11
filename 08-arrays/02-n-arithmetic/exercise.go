package narithmetic

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// nArithmetic returns the result of an arithmetic operation over "n" elements.
func nArithmetic(elems [10]int) int {
	accumulator := elems[0]

	for i := 1; i < len(elems); i++ {
		accumulator -= elems[i]
	}

	return accumulator
}
