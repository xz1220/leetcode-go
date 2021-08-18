/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 *
 * https://leetcode-cn.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (44.11%)
 * Likes:    594
 * Dislikes: 0
 * Total Accepted:    119.5K
 * Total Submissions: 270.9K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X'
 * 填充。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
 * 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
 * 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O'
 * 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["X"]]
 * 输出：[["X"]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n == board[i].length
 * 1
 * board[i][j] 为 'X' 或 'O'
 *
 *
 *
 *
 */

// @lc code=start
func solve(board [][]byte) {
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	rows, columns := len(board), len(board[0])
	if rows == 0 || columns == 0 {
		return
	}
	// 标记边缘
	q := make([][]int, 0)
	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			board[i][0] = 'A'
			q = append(q, []int{i, 0})
		}
		if board[i][columns-1] == 'O' {
			board[i][columns-1] = 'A'
			q = append(q, []int{i, columns - 1})
		}
	}

	for i := 0; i < columns; i++ {
		if board[0][i] == 'O' {
			board[0][i] = 'A'
			q = append(q, []int{0, i})
		}
		if board[rows-1][i] == 'O' {
			board[rows-1][i] = 'A'
			q = append(q, []int{rows - 1, i})
		}
	}

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx < 0 || nx >= rows || ny < 0 || ny >= columns || board[nx][ny] != 'O' {
				continue
			}
			q = append(q, []int{nx, ny})
			board[nx][ny] = 'A'
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

// @lc code=end

