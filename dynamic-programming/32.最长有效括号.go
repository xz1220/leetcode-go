/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (34.29%)
 * Likes:    1228
 * Dislikes: 0
 * Total Accepted:    135.5K
 * Total Submissions: 390.6K
 * Testcase Example:  '"(()"'
 *
 * 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "(()"
 * 输出：2
 * 解释：最长有效括号子串是 "()"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = ")()())"
 * 输出：4
 * 解释：最长有效括号子串是 "()()"
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = ""
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * s[i] 为 '(' 或 ')'
 * 
 * 
 * 
 * 
 */

// @lc code=start
func longestValidParentheses(s string) int {
	// dp[i] =  dp[i-2] + 2      if s[i-1] == "(" && s[i] == ")"
	//          dp[i-1] + 2 + dp[i - dp[i-1] -2]       if s[i-1] == s[i] == ")" && dp[i - dp[i-1] -1] == "("
	
		m := len(s)
		if m<=1 {
			return 0
		}
		dp := make([]int, m+1)
		dp[0] = 0
		maxLength := 0
	
		for i := 1 ; i <= m ; i ++ {
			if s[i-1] == ')' {
				if s[i-1] == '(' {
					dp[i] = dp[i-2] + 2
					// fmt.Println("i:", i, "  dp[i]:",dp[i])
					maxLength = max(maxLength, dp[i])
				}else {
					if i - dp[i-1] -2>=0 && s[i - dp[i-1] -2] == '(' {
						dp[i] = dp[i-1] + dp[i - dp[i-1] - 2] + 2
						maxLength = max(maxLength, dp[i])
						// fmt.Println("i:", i, "  dp[i]:",dp[i])
					}
				}
			}
		}
	
		return maxLength
	}
	
	func max(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}
		return num2
	}
// @lc code=end

