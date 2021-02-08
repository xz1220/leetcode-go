/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.86%)
 * Likes:    681
 * Dislikes: 0
 * Total Accepted:    184.3K
 * Total Submissions: 401.7K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 *
 *
 * 示例：
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 10^3
 * -10^3 <= nums[i] <= 10^3
 * -10^4 <= target <= 10^4
 *
 *
 */
// [-4, -1. 10, 20，100]
// @lc code=start
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	var distance int = 1<<63 - 1
	var result int

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		for second := first + 1; second < n; second++ {

			third := n - 1

			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			temp := 1<<63 - 1
			nextTemp := 1<<63 - 1
			for second < third {
				temp = abs(nums[second] + nums[third] + nums[first] - target)
				nextTemp = abs(nums[second] + nums[third-1] + nums[first] - target)
				if temp < nextTemp {
					break
				}
				third--
			}

			if temp < distance {
				distance = temp
				if second == third {
					result = nums[second] + nums[third+1] + nums[first]
				} else {
					result = nums[second] + nums[third] + nums[first]
				}
			}
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// @lc code=end

