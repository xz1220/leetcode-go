package main

import "fmt"

func main() {
	slice := make([]int, 2, 6)
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(len(slice))
	slice = slice[1:]
	fmt.Println(len(slice))
	a := []int{1}
	slice = append(slice, a...)
	slice = append(slice, 2)
	fmt.Println(len(slice))
}
