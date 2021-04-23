/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (51.83%)
 * Likes:    1117
 * Dislikes: 0
 * Total Accepted:    240.7K
 * Total Submissions: 450K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
 * 
 * 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
 * 
 * 此外，你可以假设该网格的四条边均被水包围。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * 输出：1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * 输出：3
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
 * grid[i][j] 的值为 '0' 或 '1'
 * 
 * 
 */

// @lc code=start
func numIslands(grid [][]byte) int {
    if len(grid) ==0 {
        return 0
    }

    m, n := len(grid), len(grid[0])
    var dfs func(int, int)
    dfs = func(i ,j int) {
        grid[i][j] = '0'
        if (i -1 >=0 && grid[i-1][j] == '1') {
            dfs(i-1, j)
        }
        if (i +1 < m && grid[i+1][j] == '1') {
            dfs(i+1, j)
        }
        if (j - 1 >=0 && grid[i][j-1] == '1') {
            dfs(i, j-1)
        }
        if (j +1 < n && grid[i][j+1] == '1') {
            dfs(i, j+1)
        }
    }

    ans := 0
    for i := 0 ; i< m ; i ++ {
        for j := 0 ; j < n ; j++ {
            if grid[i][j] == '1' {
                dfs(i, j)
                ans ++
            }
        }
    }
    return ans
}
// @lc code=end

