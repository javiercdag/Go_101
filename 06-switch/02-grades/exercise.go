package grades

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// gradeExam returns the grade of an exam with the given percentage
func gradeExam(percent float32) int {
	grade := 0

	switch {
	case percent < 20:
		grade = 0
	case percent < 40:
		grade = 2
	case percent < 60:
		grade = 3
	case percent < 80:
		grade = 4
	case percent >= 80:
		grade = 5
	}

	return grade
}
