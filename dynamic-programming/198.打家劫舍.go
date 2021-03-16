/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 *
 * https://leetcode-cn.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (47.29%)
 * Likes:    1318
 * Dislikes: 0
 * Total Accepted:    252.9K
 * Total Submissions: 528.3K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 * 
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[1,2,3,1]
 * 输出：4
 * 解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 * 
 * 示例 2：
 * 
 * 输入：[2,7,9,3,1]
 * 输出：12
 * 解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 400
 * 
 * 
 */

// @lc code=start
func rob(nums []int) int {
    // 状态转移矩阵
    // nums[i] = max(nums[i-2]+nums[i], nums[i-1])

    if len(nums) == 0 {
        return 0
    }else if len(nums) == 1 {
        return nums[0]
    }else if len(nums) == 2 {
        nums[1] = max(nums[0], nums[1])
        return nums[1]
    }

    for i:=1 ; i < len(nums) ; i++ {
        if i == 1 {
            nums[1] = max(nums[0], nums[1])
            continue
        }
        nums[i] = max(nums[i-2]+nums[i], nums[i-1])
    }

    return nums[len(nums) -1]
}

func max(nums1, nums2 int) int {
    if nums1 > nums2 {
        return nums1
    }
    return nums2
}
// @lc code=end

