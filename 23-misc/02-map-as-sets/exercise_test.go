package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"bytes"
)

func TestContain(t *testing.T) {
	var stdin bytes.Buffer
	stdin.WriteString("One by one he draws us out of the world. The Empire never ended. The Head Apollo is about to return. St. Sophia is going to be born again; she was not acceptable before. The Buddha is in the park. Siddhartha sleeps (but is going to awaken). The time you have waited for has come.")

	assert.Equal(t, true, contain(&stdin, "ONE"), "The result should be: true")
}
