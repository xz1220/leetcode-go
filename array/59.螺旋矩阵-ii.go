/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (78.52%)
 * Likes:    399
 * Dislikes: 0
 * Total Accepted:    95K
 * Total Submissions: 118.6K
 * Testcase Example:  '3'
 *
 * 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3
 * 输出：[[1,2,3],[8,9,4],[7,6,5]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：[[1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 
 * 
 */

// @lc code=start
func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }

    num := 1
    rowMin, rowMax, colMin, colMax := 0, len(matrix) -1, 0, len(matrix) -1
    for rowMin <= rowMax && colMin <= colMax {
        // fmt.Println(rowMin, rowMax, colMin, colMax)
        for i := colMin ; i <= colMax ; i ++ {
            matrix[rowMin][i] = num
            num += 1
        }
        for i := rowMin +1 ; i <= rowMax ; i ++ {
            matrix[i][colMax] = num
            num += 1
        }
        for i := colMax -1 ; i >= colMin && rowMax != rowMin ; i -- {   // 当rowMax == rowMin 只应该
            matrix[rowMax][i] = num
            num += 1
        }
        for i := rowMax -1 ; i >= rowMin + 1 && colMax != colMin ; i -- {   // 同上
            matrix[i][colMin] = num
            num += 1
        }
        rowMin, rowMax, colMin, colMax = rowMin +1, rowMax -1, colMin +1, colMax -1
    }
    return matrix
}
// @lc code=end

