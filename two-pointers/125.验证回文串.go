/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.79%)
 * Likes:    325
 * Dislikes: 0
 * Total Accepted:    198.7K
 * Total Submissions: 423.4K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 * 
 * 示例 1:
 * 
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "race a car"
 * 输出: false
 * 
 * 
 */

// @lc code=start
func isPalindrome(s string) bool {
    for i,j := 0, len(s)-1 ; i<= j ; i, j = i+1, j-1{
        for i<len(s) && !isChar(s[i]) {
            i++
        } 
        for j>= 0 && !isChar(s[j]) {
            j--
        }
        if i<=j && strings.ToLower(string(s[i])) !=strings.ToLower(string(s[j])) {
            return false
        }
    }
    return true

}

func isChar(char byte) bool {
    return char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char >= '0' && char <= '9'
}
// @lc code=end

