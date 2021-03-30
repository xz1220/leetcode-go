/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (42.13%)
 * Likes:    742
 * Dislikes: 0
 * Total Accepted:    142.3K
 * Total Submissions: 306.6K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 1 
 * -100 
 * 
 * 
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
    rowMin, rowMax, colMin, colMax := 0, len(matrix) -1, 0, len(matrix[0]) -1
    // 输出：[rowMin, colMin: colMax +1], [rowMin +1:rowMax +1, colMax]
    //      [rowMax, colMin: colMax -1: -1], [rowMin -1: rowMax -1 : -1, colMin]
    // 更新: rowMin = rowMin +1, rowMax = rowMax -1, colMin = colMin +1, colMax = colMax -1
    ans := []int{}
    for rowMin <= rowMax && colMin <= colMax {
        fmt.Println(rowMin, rowMax, colMin, colMax)
        for i := colMin ; i <= colMax ; i ++ {
            ans = append(ans, matrix[rowMin][i])
        }
        for i := rowMin +1 ; i <= rowMax ; i ++ {
            ans = append(ans, matrix[i][colMax])
        }
        for i := colMax -1 ; i >= colMin && rowMax != rowMin ; i -- {   // 当rowMax == rowMin 只应该
            ans = append(ans, matrix[rowMax][i])
        }
        for i := rowMax -1 ; i >= rowMin + 1 && colMax != colMin ; i -- {   // 同上
            ans = append(ans, matrix[i][colMin])
        }
        rowMin, rowMax, colMin, colMax = rowMin +1, rowMax -1, colMin +1, colMax -1
    } 
    return ans
}
// @lc code=end

