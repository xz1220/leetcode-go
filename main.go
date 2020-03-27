package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	for i, index := range a {
		fmt.Println(i, index)
	}

	type test struct {
		val int
	}

	b := make([]test, 0)
	b = append(b, nil)
	fmt.Println(len(b))
}
