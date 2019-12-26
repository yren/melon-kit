package arrays

import (
	"reflect"
	"testing"
)

const outputFormat = "expected: %d, but got: %d, given: %v"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf(outputFormat, expected, got, numbers)
		}

	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf(outputFormat, expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{5, 6})
	expected := []int{3, 11}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v, but got: %v", expected, got)
	}
}
