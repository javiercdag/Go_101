package digits

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func multiplyDigits(n int) int {
	if n == 0 {
		return 0
	}

	product := 1

	for n != 0 {
		digit := n % 10
		product *= digit
		n /= 10 // robust go: n stays as int
	}

	return product
}
