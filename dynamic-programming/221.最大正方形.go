/*
 * @Author: your name
 * @Date: 2021-08-21 23:36:33
 * @LastEditTime: 2021-08-21 23:36:39
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /leetcode-go/dynamic-programming/221.最大正方形.go
 */
/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 *
 * https://leetcode-cn.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (47.11%)
 * Likes:    846
 * Dislikes: 0
 * Total Accepted:    129.2K
 * Total Submissions: 274K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [["0","1"],["1","0"]]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：matrix = [["0"]]
 * 输出：0
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
 * matrix[i][j] 为 '0' 或 '1'
 *
 *
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	var maxside int
	sideMatrix := make([][]int, len(matrix))
	for index, _ := range sideMatrix {
		sideMatrix[index] = make([]int, len(matrix[index]))
		for index2, _ := range sideMatrix[index] {
			sideMatrix[index][index2] = int(matrix[index][index2] - '0')
			if sideMatrix[index][index2] == 1 {
				maxside = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if sideMatrix[i][j] == 1 {
				sideMatrix[i][j] = min(min(sideMatrix[i-1][j], sideMatrix[i][j-1]), sideMatrix[i-1][j-1]) + 1
				if sideMatrix[i][j] > maxside {
					maxside = sideMatrix[i][j]
				}
			}
		}
	}

	return maxside * maxside
}

func min(nums1, nums2 int) int {
	if nums1 > nums2 {
		return nums2
	}
	return nums1
}

// @lc code=end

