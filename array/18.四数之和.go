/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (39.62%)
 * Likes:    739
 * Dislikes: 0
 * Total Accepted:    152.1K
 * Total Submissions: 382.1K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	// 算法：
	// 计算数组长度，排序数组
	// 外双重循环，并保证不选用一样的数字
	// 内双重循环使用双指针法

	sort.Ints(nums)
	var n int = len(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			fourth := n - 1
			sum := target - nums[first] - nums[second]
			for third := second + 1; third < n; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}

				for third < fourth && nums[third]+nums[fourth] > sum {
					fourth--
				}

				if third == fourth {
					continue
				}

				if nums[third]+nums[fourth] == sum {
					ans = append(ans, []int{nums[first], nums[second], nums[third], nums[fourth]})
				}
			}

		}
	}

	return ans
}

// @lc code=end

