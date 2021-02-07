/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (34.42%)
 * Likes:    926
 * Dislikes: 0
 * Total Accepted:    133K
 * Total Submissions: 365.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 *
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 *
 * 必须 原地 修改，只允许使用额外常数空间。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,1]
 * 输出：[1,2,3]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,1,5]
 * 输出：[1,5,1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 *
 *
 */

// @lc code=start
func nextPermutation(nums []int) {
	// 1 5 3 2
	// 2 1 3 5
	// 说明：很巧妙，从左边找到一个较小数，然后在右边找到一个较大数。对于这两个数，我们的要求是：
	// - 较小数：尽量靠近右边
	// - 较大数：比较小数大但是尽量小
	// 算法：nextPermutation(nums slice)

	// 初始化并处理边界条件
	var i = len(nums) - 1
	var j = len(nums) - 1

	if len(nums) <= 1 {
		return
	}

	// 找出较小数，即从右往左找出不满足降序的连续两个数，下标较小的就是较小数
	for ; i > 0 && nums[i-1] >= nums[i]; i -= 1 {

	}

	// 找出较大数，由于从下标i到数组末尾的数是降序的，即从右往左找到第一个大于nums[i-1]的数即可，i!=0 处理边界条件
	for ; i != 0 && j >= i; j -= 1 {
		if nums[j] > nums[i-1] {
			nums[j], nums[i-1] = nums[i-1], nums[j]
			break
		}
	}

	// 由于交换后的[i,...]仍旧是逆序的，交换后已经保证了此时的序列是大于之前的，这时候我们想要这个序列尽量的小，而最小的极为升序
	for reverseIndex := len(nums) - 1; reverseIndex >= i; reverseIndex, i = reverseIndex-1, i+1 {
		nums[reverseIndex], nums[i] = nums[i], nums[reverseIndex]
	}
}

// @lc code=end

