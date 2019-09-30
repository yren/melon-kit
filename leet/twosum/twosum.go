package twosum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		tmp := target - v
		if j, ok := m[tmp]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}