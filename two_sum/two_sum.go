package two_sum

import "fmt"

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

//两遍hash 时间复杂福O(n)
func twoSum2(nums []int, target int) []int {
	hash := make(map[int]int)

	//先存一遍
	for j, num := range nums {
		hash[num] = j
	}

	fmt.Println(hash)

	for i, val := range nums {
		another := target - val
		index, ok := hash[another] //读一遍
		if ok && (index != i) {
			return []int{i,index}
		}
	}

	return []int{0,0}
}

//两遍hash 时间复杂福O(n)
func twoSum3(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, val := range nums {
		another := target - val
		index, ok := hash[another] //读一遍
		if ok && (index != i) {
			return []int{index, i} //hash exist index before current i
		}

		hash[val] = i
	}

	return []int{0,0}
}
