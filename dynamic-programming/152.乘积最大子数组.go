/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (40.81%)
 * Likes:    982
 * Dislikes: 0
 * Total Accepted:    124.4K
 * Total Submissions: 301.8K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 * 
 */

// @lc code=start
// 要注意这一道题，并不满足最优子结构，因为当前的状态的最优不一定是前一个状态的最优转化过来的

func maxProduct(nums []int) int {
// 需要维护两个变量， fmax, fmin
	fmax, fmin, ans := nums[0], nums[0], nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		tmax, tmin := fmax, fmin
		fmax = max(tmax*nums[i], max(tmin*nums[i], nums[i]))
		fmin = min(tmax*nums[i], min(tmin*nums[i], nums[i]))
		ans = max(fmax, ans)
	}

	return ans
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
} 

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
// @lc code=end

