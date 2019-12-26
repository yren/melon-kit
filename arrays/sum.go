package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	numbersLength := len(numbersToSum)
	sums = make([]int, numbersLength)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
