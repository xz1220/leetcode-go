/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	var (
		left int
		right int
	)

	for right = len(s) - 1; right >= 0; right-- {
		if IsLetter(s[right]) {
			break
		}
	}

	if right < 0 {
		return 0
	}

	for left = right; left >= 0; left-- {
		if !IsLetter(s[left]) {
			break
		}
	}

	return right - left
}

func IsLetter(s byte) bool {
	if (s >= 65 && s <= 90) || (s >= 97 && s <= 122) {
		return true
	}
	return false
}
// @lc code=end

