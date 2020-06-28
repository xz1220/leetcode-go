/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (33.52%)
 * Likes:    3373
 * Dislikes: 0
 * Total Accepted:    408.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 * 
 * 示例 1:
 * 
 * 输入: "abcabcbb"
 * 输出: 3 
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 * 
 * 
 * 示例 2:
 * 
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 * 
 * 
 * 示例 3:
 * 
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 * 
 * 
 */

// @lc code=start

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

	maxLength = max(maxLength, j-i)  //必要：由于循环中需要进入if语句才能进行max，如果最后找到的子串最大但是又没有进入循环就完了
	return maxLength

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

