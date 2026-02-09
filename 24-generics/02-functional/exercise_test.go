package functional

import (

	"strconv"

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenericFunction(t *testing.T) {
	
	
	valuesInt := []int{2,5,8,88,12}
	valuesStr := []string{"22","3","4","56"}
	reducedInt := reduce(valuesInt, 1, func(a ,b int) int {return a * b})
	reducedStr := reduce(valuesStr, 0,func(a int, b string) int {
		parsed, _ := strconv.Atoi(b)
		return parsed + a
	})
	assert.Equal(t, 84480,reducedInt)
	assert.Equal(t, 85,reducedStr)
	
	
	
}
