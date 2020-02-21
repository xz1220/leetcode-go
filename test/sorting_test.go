package test

import (
	"sorting"
	"testing"
)

func TestMerge(t *testing.T) {
	var example = [][]int{{2, 6}, {8, 10}, {1, 3}, {15, 18}}
	var example2 = [][]int{{1, 4}, {4, 5}}
	var example3 = [][]int{{1, 4}, {2, 3}}
	sorting.Merge2(example)
	sorting.Merge2(example2)
	sorting.Merge2(example3)
	//fmt.Println("projPath is:")

}
