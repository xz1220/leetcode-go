/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode-cn.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (35.54%)
 * Likes:    394
 * Dislikes: 0
 * Total Accepted:    30.9K
 * Total Submissions: 86.8K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
 * 
 * 示例：
 * 
 * 输入: S = "ADOBECODEBANC", T = "ABC"
 * 输出: "BANC"
 * 
 * 说明：
 * 
 * 
 * 如果 S 中不存这样的子串，则返回空字符串 ""。
 * 如果 S 中存在这样的子串，我们保证它是唯一的答案。
 * 
 * 
 */

// @lc code=start
func minWindow(s string, t string) string {


	// if len(s)<len(t){
	// 	return ""
	// }
	flag:=false
	begin, end := 0, len(s)
	i, j := 0, 0
	count := 0
	countMax:=0

	sRune := []rune(s)
	tRune := []rune(t)

	mapT := make(map[rune]int)
	mapS:=make(map[rune]int)
	for _, r := range tRune {
		if _,ok:=mapT[r];!ok{
			countMax+=1
		}
		mapT[r] +=1
	}

	for j < len(sRune) {
		log.Println("外部循环  ", j)
		mapS[sRune[j]] += 1
		if _, ok := mapT[sRune[j]]; ok {
			log.Println("判断是否存在键")
			if mapT[sRune[j]] == mapS[sRune[j]]{
				log.Println("键值+1")
				count++
			}
		}

		if count == countMax {
			flag=true
			log.Println("缩减开始")
			for {
				i++
				mapS[sRune[i-1]] -= 1
				if _, ok := mapT[sRune[i-1]]; ok {
					if mapS[sRune[i-1]] < mapT[sRune[i-1]] {
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

	if flag{
		return string(sRune[begin:end])
	}else{
		return ""
	}

	

}
// @lc code=end

