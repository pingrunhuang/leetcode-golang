package utils

// NewSet
func NewSet(arr []int) []int {
	m := make(map[int]bool)
	result := make([]int, 0, len(arr))
	for _, val := range arr {
		if _, ok := m[val]; ok {
			m[val] = true
			result = append(result, val)
		}
	}
	return result
}
