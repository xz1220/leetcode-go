/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (70.13%)
 * Likes:    645
 * Dislikes: 0
 * Total Accepted:    94.1K
 * Total Submissions: 129.3K
 * Testcase Example:  '"aab"'
 *
 * 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
 * 
 * 回文串 是正着读和反着读都一样的字符串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "aab"
 * 输出：[["a","a","b"],["aa","b"]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "a"
 * 输出：[["a"]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 仅由小写英文字母组成
 * 
 * 
 */

// @lc code=start
func partition(s string) [][]string {
    // dp + 回溯

    // 初始化矩阵
    var ans [][]string
    n := len(s)
    f := make([][]bool, n)
    for i := range f {
        f[i] = make([]bool, n)
        for j := range f[i] {
            f[i][j] = true
        }
    }

    // 记录dp举证
    // f[i][j] = true           if i >= j 
    //           f[i+1][j-1] && (s[i] == s[j])  otherwise
    for i:= n-1 ; i >=0 ; i-- {                               // 自下向上更新
        for j := i+1 ; j < n ; j++ {
            f[i][j] = f[i+1][j-1] && (s[i] == s[j])           // 应该先更新i+1层
        }
    }

    // 回溯生成答案
    var split []string
    var dfs func(int) 
    dfs = func(i int) {
        if i == n {
            ans = append(ans, append([]string(nil), split...))
        }

        for j := i ; j < n ; j++ {
            if f[i][j] {
                split = append(split, s[i:j+1])
                dfs(j+1)
                split = split[:len(split) -1]
            }
        }
    }

    dfs(0)
    return ans
}
// @lc code=end

