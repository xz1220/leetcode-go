/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (62.58%)
 * Likes:    272
 * Dislikes: 0
 * Total Accepted:    108.7K
 * Total Submissions: 166.9K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 * 
 * 
 * 
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 * 
 * 示例:
 * 
 * 输入: 3
 * 输出: [1,3,3,1]
 * 
 * 
 * 进阶：
 * 
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 * 
 */

// @lc code=start
func getRow(rowIndex int) []int {
    ans := make([]int, rowIndex +1)
    ans[0] =1
    for i := range ans {
        if i>0 {
            ans[i] = ans[i-1] * (rowIndex - i +1)/i 
        }
    }

    return ans
}
// @lc code=end

