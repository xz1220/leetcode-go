/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (67.01%)
 * Likes:    710
 * Dislikes: 0
 * Total Accepted:    134.5K
 * Total Submissions: 200.1K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形 triangle ，找出自顶向下的最小路径和。
 * 
 * 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1
 * 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * 输出：11
 * 解释：如下面简图所示：
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：triangle = [[-10]]
 * 输出：-10
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -10^4 
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
 * 
 * 
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
    layer := len(triangle)

    for currentLayer := 1; currentLayer < layer ; currentLayer ++ {
        for currentIndex := 0; currentIndex <= currentLayer ; currentIndex ++ {
            if currentLayer == 1 {
                triangle[currentLayer][currentIndex] += triangle[0][0]
            }else if currentIndex == 0 {
                triangle[currentLayer][currentIndex] += triangle[currentLayer -1][0]
            }else if currentIndex == currentLayer {
                triangle[currentLayer][currentIndex] += triangle[currentLayer -1][currentIndex -1]
            }else if triangle[currentLayer -1][currentIndex] > triangle[currentLayer-1][currentIndex -1]{
                triangle[currentLayer][currentIndex] += triangle[currentLayer -1][currentIndex-1]
            }else {
                triangle[currentLayer][currentIndex] += triangle[currentLayer-1][currentIndex]
            }

        }
    }

    var min int = int(^uint(0) >> 1)
    
    for _, value := range triangle[layer -1] {
        if value < min {
            min = value
        }
    }

    return min
}
// @lc code=end

