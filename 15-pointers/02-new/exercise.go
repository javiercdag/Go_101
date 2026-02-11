package pointernew

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func newValue() *float64 {
	newFloat := new(float64)
	*newFloat = 12.05
	return newFloat
}