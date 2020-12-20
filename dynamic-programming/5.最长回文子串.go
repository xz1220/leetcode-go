/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (31.67%)
 * Likes:    3020
 * Dislikes: 0
 * Total Accepted:    440.4K
 * Total Submissions: 1.4M
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {

	/*
	 * 状态转移方程：
	 * P(i,i) = true
	 * P(i,i+1) = true
	 * P(i,j) = P(i+1,j-1) && S[i] == S[j]
	 *
	 */

	n := len(s)
	answer := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] > 0 && l+1 > len(answer) {
				answer = s[i : i+l+1]
			}
		}
	}
	return answer

}

// @lc code=end

