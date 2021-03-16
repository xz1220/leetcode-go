/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode-cn.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (67.98%)
 * Likes:    815
 * Dislikes: 0
 * Total Accepted:    189.6K
 * Total Submissions: 277.9K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 * 
 * 说明：每次只能向下或者向右移动一步。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 
 * 0 
 * 
 * 
 */

// @lc code=start
func minPathSum(grid [][]int) int {
    m, n := len(grid[0]), len(grid)
    state := make([]int, m)
    for index, _ := range state {
        if index == 0 {
            state[0] = grid[0][0]
        }else{
            state[index] = grid[0][index] + state[index -1]
        }
    }
    
    for i:= 1 ; i < n ; i++ {
        for j:=0 ; j< m ; j++ {
            if j == 0 {
                state[j] += grid[i][j]
            }else if state[j -1] > state[j] {
                state[j] += grid[i][j]
            }else {
                state[j] = state[j-1] + grid[i][j]
            }
        }
    }

    return state[m-1]
}
// @lc code=end

