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
    // dp[i] = 0    if s[i] == '('
    //       = dp[i-2] +2   if s[i] == ")" && s[i-1] == '('
    //       = dp[i-1] +2 +dp[i-dp[i-1]-2] if s[i] == ")" && s[i-1] == ")" && s[i - dp[i-1] - 1] == '('
    //  (((()))))()())

    length := len(s)
    dp := make([]int, length +1)

    max := 0
    for i := 1; i <= length ; i++ {
        if i > 1 && s[i-1] == ')' && s[i-2] == '(' {
            dp[i] = dp[i-2] +2
        }else if i - dp[i-1] - 2 >=0 && s[i -1] == ')' && s[i-2] == ')' && s[i - dp[i-1] - 2] == '(' {
            dp[i] = dp[i-1] +2 + dp[i-dp[i-1] -2]
        }
        if dp[i] > max {
            max = dp[i]
        }
    }

    return max
}


func longestValidParentheses(s string) int {
    l, r , max := 0, 0, 0

    for i := range s {
        if s[i] == '(' {
            l ++
        }else if s[i] == ')' {
            r ++
        }
        if r == l {
            if max < l {
                max = l
            }
        }else if r > l {
            l, r = 0, 0
        }
    }

    l, r = 0, 0
    for i := len(s) -1 ; i >= 0 ; i-- {
                if s[i] == '(' {
            l ++
        }else if s[i] == ')' {
            r ++
        }
        if r == l {
            if max < l {
                max = l
            }
        }else if r < l {
            l, r = 0, 0
        }
    }
    return max*2
}


// @lc code=end

