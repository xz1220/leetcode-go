package array

import (
	"array"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	correct := []int{1, 0}
	output := array.TwoSum(nums, target)
	if !reflect.DeepEqual(correct, output) {
		t.Error("不一样")
	}
}
