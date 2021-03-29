/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode-cn.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (37.13%)
 * Likes:    615
 * Dislikes: 0
 * Total Accepted:    168.9K
 * Total Submissions: 451.8K
 * Testcase Example:  '2.00000\n10'
 *
 * 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x^n）。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：x = 2.00000, n = 10
 * 输出：1024.00000
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：x = 2.10000, n = 3
 * 输出：9.26100
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：x = 2.00000, n = -2
 * 输出：0.25000
 * 解释：2^-2 = 1/2^2 = 1/4 = 0.25
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * -100.0 
 * -2^31 
 * -10^4 
 * 
 * 
 */

// @lc code=start
func myPow(x float64, n int) float64 {
    if n == 0 {
        return float64(1)
    }

    signed := 1
    if n <= 0 {
        signed = -1
        n = -n
    }

    var pow func(x float64, n int) float64 
    pow = func(x float64, n int) float64 {
        if n == 0 {
            return float64(1)
        }else if n == 1 {
            return x
        }
        res := n % 2
        next := (n-res)/2
        powNext := pow(x, next)
        return powNext*powNext*pow(x, res)
    }

    if signed == 1 {
        return pow(x, n)
    }else {
        return float64(1)/pow(x, n)
    }
}
// @lc code=end

