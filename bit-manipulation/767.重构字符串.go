/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 *
 * https://leetcode-cn.com/problems/reorganize-string/description/
 *
 * algorithms
 * Medium (47.94%)
 * Likes:    431
 * Dislikes: 0
 * Total Accepted:    47.8K
 * Total Submissions: 99.2K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s ，检查是否能重新排布其中的字母，使得两相邻的字符不同。
 *
 * 返回 s 的任意可能的重新排列。若不可行，返回空字符串 "" 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "aab"
 * 输出: "aba"
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "aaab"
 * 输出: ""
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= s.length <= 500
 * s 只包含小写字母
 *
 *
 */

// @lc code=start
func reorganizeString(s string) string {
	//
	n := len(s)
	if n <= 1 {
		return s
	}

	digitMap := [26]int{}
	maxCount := 0

	for i := range s {
		index := s[i] - 'a'
		digitMap[index] += 1

		if digitMap[index] > maxCount {
			maxCount = digitMap[index]
		}
	}

	if maxCount > (n+1)/2 {
		return ""
	}

	ans := make([]byte, n)
	evenIndex, oddIndex, halfLen := 0, 1, n/2

	for i, c := range digitMap {
		ch := byte(i + 'a')

		for c > 0 && c <= halfLen && oddIndex < n {
			ans[oddIndex] = ch
			c--
			oddIndex += 2
		}

		for c > 0 {
			ans[evenIndex] = ch
			c--
			evenIndex += 2
		}
	}

	return string(ans)
}

// @lc code=end

