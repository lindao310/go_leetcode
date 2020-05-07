package two_sum

//暴力解法 时间复杂福O(n^2)
func twoSum(nums []int, target int) []int {
	for i, val := range nums {
		for j, val2 := range nums {
			if (j != i) && (val + val2 == target) {
				return []int{i, j}
			}
		}
	}

	return nil
}
