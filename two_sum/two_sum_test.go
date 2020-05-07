package two_sum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	nums := []int{3,2,7,11,15,2}

	target := 5

	result := []int{0,1}

	//ret := twoSum(nums, target)
	//ret := twoSum2(nums, target) //has bug,in this case,return [0,5]
	ret := twoSum3(nums, target)

	fmt.Println("twoSum return ", ret)

	if ret[0] != result[0] || ret[1] != result[1] {
		t.Fail()
	}
}
