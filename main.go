package main

import (
	"log"
)

func lengthOfLongestSubstring(s string) int {

	//将字符串转换成rune
	r := []rune(s)
	i, j := 0, 0
	maxLength := 0

	//make 和new的区别详解 ？ 重写new?
	count := make(map[rune]int)

	for _, char := range r {
		count[char] = 0
	}

	for j < len(r) {
		count[r[j]] += 1
		if count[r[j]] == 2 {
			maxLength = max(maxLength, j-i)

			for {
				count[r[i]] -= 1
				i++
				if r[i-1] == r[j] {
					break
				}
			}
		}
		j++
	}

	maxLength = max(maxLength, j-i)
	return maxLength

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

	s := "aab"
	count := lengthOfLongestSubstring(s)
	log.Println(count)
}
