package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
	assert.NotEqual(t, 123, 456, "they should not equal")

	var obj []int
	assert.Nil(t, obj)

}
