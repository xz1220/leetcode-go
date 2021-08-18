/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (39.30%)
 * Likes:    732
 * Dislikes: 0
 * Total Accepted:    349.8K
 * Total Submissions: 890K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

// @lc code=start
func mySqrt(x int) int {
	l, r := 0, x

	for l < r {
		m := (l + r) /2
		if m*m == x {
			return m
		}else if m*m < x {
			l = m +1
		}else {
			r = m - 1
		}
	}

	if l * l > x {
		return l - 1
	}
	return l
}
// @lc code=end

