package grades

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradeExam(t *testing.T) {
	assert.Equal(t, 5, gradeExam(100))
	assert.Equal(t, 5, gradeExam(80))
	assert.Equal(t, 4, gradeExam(60))
	assert.Equal(t, 3, gradeExam(40))
	assert.Equal(t, 2, gradeExam(20))
	assert.Equal(t, 0, gradeExam(0))
}
