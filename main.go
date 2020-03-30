package main

import (
	"fmt"
	"log"
)

func minWindow(s string, t string) string {

	begin, end := 0, len(s)
	i, j := 0, 0
	count := 0

	sRune := []rune(s)
	tRune := []rune(t)

	mapT := make(map[rune]int)
	for _, r := range tRune {
		mapT[r] = 0
	}

	for j < len(sRune) {
		log.Println("外部循环  ", j)
		if _, ok := mapT[sRune[j]]; ok {
			log.Println("判断是否存在键")
			mapT[sRune[j]] += 1
			if mapT[sRune[j]] == 1 {
				log.Println("键值+1")
				count++
			}
		}

		if count == len(tRune) {
			log.Println("缩减开始")
			for {
				i++
				if _, ok := mapT[sRune[i-1]]; ok {
					mapT[sRune[i-1]] -= 1
					if mapT[sRune[i-1]] == 0 {
						if end-begin > (j+1)-(i-1) {
							end = j + 1
							begin = i - 1
						}
						count--
						break
					}
				}
			}
		}
		j++

	}

	log.Println("\nbegin:", begin, "\nend:", end, "\ni：", i, "\nj:", j)

	return string(sRune[begin:end])

}

func main() {

	s := "ADOBECODEBANC"
	t := "ABC"
	re := minWindow(s, t)
	fmt.Println(re)

}
