/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (63.90%)
 * Likes:    917
 * Dislikes: 0
 * Total Accepted:    220.1K
 * Total Submissions: 341.4K
 * Testcase Example:  '3\n7'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 * 
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 * 
 * 问总共有多少条不同的路径？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：m = 3, n = 7
 * 输出：28
 * 
 * 示例 2：
 * 
 * 
 * 输入：m = 3, n = 2
 * 输出：3
 * 解释：
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向下
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：m = 7, n = 3
 * 输出：28
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：m = 3, n = 3
 * 输出：6
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 题目数据保证答案小于等于 2 * 10^9
 * 
 * 
 */

// @lc code=start
// 解法1: DP
func uniquePaths(m int, n int) int {
    // 初始化状态转化方程 
    state := make([]int, n)
    for index, _ := range state {
        state[index] = 1
    }

    // 状态转移
    for i:= 0 ; i< m-1 ; i++ {
        for j:=1 ; j < n ; j++ {
            state[j] += state[j-1]
        }
    }

    return state[len(state) -1]
}

// 组合数学
func uniquePaths(m, n int) int {
    return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
// @lc code=end

