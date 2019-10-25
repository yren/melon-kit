package sum

func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, value := range numbersToSum {
		sums = append(sums, Sum(value))
	}
	return
}