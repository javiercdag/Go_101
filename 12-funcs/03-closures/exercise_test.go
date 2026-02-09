package closures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxy(t *testing.T) {
	
	
	limitedOne := proxy(2, func() int { return 2 })
	limitedTwo := proxy(3, func() int { return 3 })
	for i := 0; i < 2; i++ {
		ret, err := limitedOne()
		assert.Nil(t, err)
		assert.Equal(t, 2, ret)
	}
	ret, err := limitedOne()
	assert.NotNil(t, err)
	assert.Equal(t, 0, ret)
	for i := 0; i < 3; i++ {
		ret, err = limitedTwo()
		assert.Nil(t, err)
		assert.Equal(t, 3, ret)
	}
	ret, err = limitedTwo()
	assert.NotNil(t, err)
	assert.Equal(t, 0, ret)
	
	
	
}