package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if m[target-v] != 0 && m[target-v]-1 != k {
			return []int{m[target-v] - 1, k}
		}
		m[v] = k + 1
	}
	return nil
}
