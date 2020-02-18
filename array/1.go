package array

//TwoSum is A func written by others
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{i, idx}
		}
		m[v] = i
	}
	return nil
}
