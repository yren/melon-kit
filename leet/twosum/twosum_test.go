package twosum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expect := []int{0, 1}
	result := twoSum(nums, target)
	assert.Equal(t, expect, result)
}
