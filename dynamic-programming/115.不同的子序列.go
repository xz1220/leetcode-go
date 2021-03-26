/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 *
 * https://leetcode-cn.com/problems/distinct-subsequences/description/
 *
 * algorithms
 * Hard (49.77%)
 * Likes:    491
 * Dislikes: 0
 * Total Accepted:    47.3K
 * Total Submissions: 87.1K
 * Testcase Example:  '"rabbbit"\n"rabbit"'
 *
 * 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
 * 
 * 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE"
 * 的一个子序列，而 "AEC" 不是）
 * 
 * 题目数据保证答案符合 32 位带符号整数范围。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "rabbbit", t = "rabbit"
 * 输出：3
 * 解释：
 * 如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
 * (上箭头符号 ^ 表示选取的字母)
 * rabbbit
 * ^^^^ ^^
 * rabbbit
 * ^^ ^^^^
 * rabbbit
 * ^^^ ^^^
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "babgbag", t = "bag"
 * 输出：5
 * 解释：
 * 如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
 * (上箭头符号 ^ 表示选取的字母)
 * babgbag
 * ^^ ^
 * babgbag
 * ^^    ^
 * babgbag
 * ^    ^^
 * babgbag
 * ⁠ ^  ^^
 * babgbag
 * ⁠   ^^^
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * s 和 t 由英文字母组成
 * 
 * 
 */

// @lc code=start
func numDistinct(s string, t string) int {
    // dp[i][j] = 0                             if i < j or (i == j and s[i] != t[j])
    //          = bp[i-1][j]                    if i >= j and s[i] != t[j]
    //          = bp[i-1][j] + bp[i-1][j-1]     if i>=j and s[i] == t[j]
    // 边界条件：
    //       1. i >= 0 and j == 0       bp[i][j] = 1   
    //       2. i == 0 and j > 0       bp[i][j] = 0

    sLength, tLength := len(s), len(t)
    bp := make([][]int, sLength +1)
    for index := range bp {
        bp[index] = make([]int, tLength+1)
    }

    for i:= 0 ; i <=sLength ; i ++ {
        for j := 0 ; j <= tLength ; j ++ {
            if j == 0 {
                bp[i][j] = 1
            }else if i > 0 && i >= j && s[i-1] != t[j-1] {
                bp[i][j] = bp[i-1][j]
            }else if i > 0 && i >= j && s[i-1] == t[j-1] {
                bp[i][j] = bp[i-1][j] + bp[i-1][j-1]
            }
        }
    } 

    return bp[sLength][tLength]
    
}
// @lc code=end

