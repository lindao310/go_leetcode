package two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{2,7,11,15,2}

	target := 4

	result := []int{0,4}

	ret := twoSum(nums, target)

	fmt.Println("twoSum return ", ret)

	if ret[0] != result[0] || ret[1] != result[1] {
		t.Fail()
	}
}
