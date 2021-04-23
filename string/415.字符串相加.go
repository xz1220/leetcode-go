/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (51.86%)
 * Likes:    360
 * Dislikes: 0
 * Total Accepted:    111.1K
 * Total Submissions: 212.2K
 * Testcase Example:  '"11"\n"123"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * num1 和num2 的长度都小于 5100
 * num1 和num2 都只包含数字 0-9
 * num1 和num2 都不包含任何前导零
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
 * 
 * 
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
    ans := ""
    i, j , flag := len(num1) -1, len(num2) -1, 0
    for i>=0 || j >=0 || flag != 0 {
        x, y := 0, 0
        if i >= 0 {
            x = int(num1[i] - '0')
            i --
        }
        if j >= 0 {
            y = int(num2[j] - '0')
            j --
        }

        temp := x +y +flag
        flag = temp / 10
        num := strconv.Itoa(temp % 10)
        ans = num + ans
    }
    return ans
}
// @lc code=end

