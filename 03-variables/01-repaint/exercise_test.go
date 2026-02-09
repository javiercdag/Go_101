package repaint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepaintColor(t *testing.T) {
	c, err := repaintColor("red")
	assert.Equal(t, "green", c)
	assert.NoError(t, err)

	c, err = repaintColor("green")
	assert.Equal(t, "red", c)
	assert.NoError(t, err)

	c, err = repaintColor("xxx")
	assert.Equal(t, "", c)
	assert.Error(t, err)
}
