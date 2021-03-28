/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.68%)
 * Likes:    1668
 * Dislikes: 0
 * Total Accepted:    249.6K
 * Total Submissions: 324.1K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：["()"]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 
 * 
 */

// @lc code=start
func generateParenthesis(n int) []string {

    ans := make([]string, 0)
    var dfs func (l, r int, curStr string) 
    dfs = func (l, r int, curStr string) {
        if l == 0 && r == 0 {
            ans = append(ans, curStr)
        }

        if l > 0 {
            dfs(l-1, r,curStr+"(")
        }
        if l < r {
            dfs(l, r -1, curStr+")")
        }
    }

    dfs(n,n,"")
    
    return ans 
}
// @lc code=end

