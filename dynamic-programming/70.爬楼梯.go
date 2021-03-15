/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (51.21%)
 * Likes:    1525
 * Dislikes: 0
 * Total Accepted:    383.5K
 * Total Submissions: 743.6K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * 
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 
 * 注意：给定 n 是一个正整数。
 * 
 * 示例 1：
 * 
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 * 
 * 示例 2：
 * 
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 * 
 * 
 */

// @lc code=start
// 斐波那契数列方法
// func climbStairs(n int) int {
//     // 我们发现：f(x) = f(x-1) + f(x-2)
//     if n == 0 || n ==1 {
//         return 1
//     }

//     var pre, current = 1, 1
//     for i:=2 ; i <= n ; i++ {
//         pre, current = current, pre + current
//     }

//     return current
// }

// 特征方程法
// func climbStairs(n int) int {
//     sqrt5 := math.Sqrt(5)
//     pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
//     pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
//     return int(math.Round((pow1 - pow2) / sqrt5))
// }

// 矩阵快速幂
type matrix [2][2]int

func mul(a, b matrix) (c matrix) {
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
        }
    }
    return c
}

func pow(a matrix, n int) matrix {
    res := matrix{{1, 0}, {0, 1}}  // I
    for ; n > 0; n >>= 1 {
        if n&1 == 1 {
            res = mul(res, a)
        }
        a = mul(a, a)
    }
    return res
}

func climbStairs(n int) int {
    res := pow(matrix{{1, 1}, {1, 0}}, n)
    return res[0][0]
}


// @lc code=end

