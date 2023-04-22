/*
 * @lc app=leetcode.cn id=1027 lang=golang
 *
 * [1027] 最长等差数列
 *
 * https://leetcode.cn/problems/longest-arithmetic-subsequence/description/
 *
 * algorithms
 * Medium (43.93%)
 * Likes:    298
 * Dislikes: 0
 * Total Accepted:    32.3K
 * Total Submissions: 67.5K
 * Testcase Example:  '[3,6,9,12]'
 *
 * 给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。
 *
 * 回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ...
 * < ik <= nums.length - 1。并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1)
 * 的值都相同，那么序列 seq 是等差的。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,6,9,12]
 * 输出：4
 * 解释：
 * 整个数组是公差为 3 的等差数列。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [9,4,7,2,10]
 * 输出：3
 * 解释：
 * 最长的等差子序列是 [4,7,10]。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [20,1,15,3,10,5,8]
 * 输出：4
 * 解释：
 * 最长的等差子序列是 [20,15,10,5]。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 1000
 * 0 <= nums[i] <= 500
 *
 *
 */

// @lc code=start
func longestArithSeqLength(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	var maxLen int
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, 1001)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j] + 500
			dp[i][d] = dp[j][d] + 1

			if dp[i][d] > maxLen {
				maxLen = dp[i][d]
			}
		}
	}

	return maxLen + 1
}

// @lc code=end

