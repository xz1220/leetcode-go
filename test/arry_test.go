package test

import (
	"array"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	correct := []int{0, 1}
	output := array.TwoSum(nums, target)
	if !equalInt(output, correct) {
		t.Error("不一样")
	}
}

func equalInt(input []int, correct []int) bool {
	if (input == nil) != (correct == nil) {
		return false
	}
	if len(input) != len(correct) {
		return false
	}

	sort.Ints(input)
	sort.Ints(correct)

	for i := range input {
		if input[i] != correct[i] {
			return false
		}
	}
	return true
}
