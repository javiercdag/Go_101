package closures

import "errors"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate
func proxy(limit int, f func() int) func() (int, error) {
	count := 0

	return func() (int, error) {
		if count < limit {
			count++
			return f(), nil
		}

		return 0, errors.New("limit exceeded")
	}
}