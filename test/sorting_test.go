package test

import (
	"fmt"
	"math/rand"
	"sorting"
	"testing"
	"time"
)

func randSlice(seed int64, length int) []int {
	rand.Seed(seed)
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(1000)
	}
	//fmt.Println(slice)
	return slice
}

func isOrder(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}

// func TestMerge(t *testing.T) {
// 	var example = [][]int{{2, 6}, {8, 10}, {1, 3}, {15, 18}}
// 	var example2 = [][]int{{1, 4}, {4, 5}}
// 	var example3 = [][]int{{1, 4}, {2, 3}}
// 	sorting.Merge2(example)
// 	sorting.Merge2(example2)
// 	sorting.Merge2(example3)
// 	//fmt.Println("projPath is:")

// }

func TestBubble(t *testing.T) {
	slice := randSlice(time.Now().Unix(), 10)
	slice = sorting.BubbleSort(slice)
	if !isOrder(slice) {
		t.Error("Error")
	}
	sliceForMerge := randSlice(time.Now().Unix(), 10)
	fmt.Println(sliceForMerge)
	sliceForMerge = sorting.MergeSort(sliceForMerge)
	fmt.Println(sliceForMerge)
	if !isOrder(sliceForMerge) {
		t.Error("Error")
	}

}
