package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 10
	fmt.Println(b, a)
}
