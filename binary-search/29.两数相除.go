/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (20.33%)
 * Likes:    538
 * Dislikes: 0
 * Total Accepted:    83.4K
 * Total Submissions: 408.4K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 * 
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 * 
 * 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) =
 * -2
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
 * 
 * 示例 2:
 * 
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = truncate(-2.33333..) = -2
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 * 
 * 
 */

// @lc code=start
func divide(dividend int, divisor int) int {
    // init ans = 1
    // divisor 左移 1 位 ， ans = ans << 1,
    // 若超过，减去并

    // 边界条件处理
    if dividend == 0 {
        return 0
    }else if divisor == 1 {
        return dividend
    }else if divisor == -1 {
        if dividend == -(1<<31) {
            return (1<<31) -1 
        }
        return -dividend
    }else if dividend == divisor {
        return 1
    }

    sign := 1
    dividend64 := int64(dividend)
    divisor64 := int64(divisor)
    if (dividend64 > 0 && divisor64< 0) || (dividend64 < 0 && divisor64 > 0) {
        sign = -1
    }
    if dividend64 <= 0 {
        dividend64 = - dividend64
    }
    if divisor64 <=0 {
        divisor64 = -divisor64
    }

    maxMove := 0
    for divisor64 <= dividend64 {
        divisor64 = divisor64 << 1
        maxMove ++
    }
    var result int64 = 0
    for maxMove > 0 {
        divisor64 = divisor64 >> 1
        maxMove --
        if divisor64 <= dividend64 {
            result += int64(1 << maxMove)
            dividend64 -= divisor64
        }
    }
    return int(result)*sign
}
// @lc code=end

