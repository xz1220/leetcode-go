/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (79.44%)
 * Likes:    1075
 * Dislikes: 0
 * Total Accepted:    216.6K
 * Total Submissions: 272.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 * 
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0]
 * 输出：[[],[0]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10 
 * nums 中的所有元素 互不相同
 * 
 * 
 */

// @lc code=start
func subsets(nums []int) [][]int {
    ans := make([][]int, 0)
    ans = append(ans, []int{})

    for index := range nums {
        for _, val := range ans {
            temp := append(val, nums[index])  
            copyed := make([]int, len(temp))
            copy(copyed, temp)                      // 在 len < cap 如果不进行copy. 可能会在同一块内存中进行操作，导致答案错误
            sort.Ints(copyed)
            ans = append(ans, copyed)
        }
    }

    return ans
}
// @lc code=end

