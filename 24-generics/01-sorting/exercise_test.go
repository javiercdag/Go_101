package sorting

import (
	
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	

	

	

	
	arrTypeOne := []string{"banan","alma","64","cseresznye","jojo","toki","wa","ugoki","desu"}
	arrTypeTwo := []uint8{1,2,7,5,4,6,3,9,8}
	assert.Equal(t,[]string{"64","alma","banan","cseresznye","desu","jojo","toki","ugoki","wa"}, sortSlice(arrTypeOne))
	assert.Equal(t, []uint8{1,2,3,4,5,6,7,8,9}, sortSlice(arrTypeTwo))
	

	
}
