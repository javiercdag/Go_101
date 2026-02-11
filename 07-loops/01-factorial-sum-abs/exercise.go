package factorial

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func calcSum(n int) int {
	if n == 1 {
		return 1
	}

	return n + calcSum(n-1)
}
