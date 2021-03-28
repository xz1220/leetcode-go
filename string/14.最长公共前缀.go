/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (39.06%)
 * Likes:    1524
 * Dislikes: 0
 * Total Accepted:    476.8K
 * Total Submissions: 1.2M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 
 * 如果不存在公共前缀，返回空字符串 ""。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * 0 
 * strs[i] 仅由小写英文字母组成
 * 
 * 
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    for index := range strs[0] {
        for index2 := range strs {
            if index == len(strs[index2]) || strs[0][:index +1] != strs[index2][:index+1] {
                return strs[0][:index]
            } 
        }
    }
    return strs[0]
}
// @lc code=end

