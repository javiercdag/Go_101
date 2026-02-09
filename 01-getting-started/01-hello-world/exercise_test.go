package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "Helló Világ!", helloWorld(), "hello world in Hungarian")
}
