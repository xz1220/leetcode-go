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
	// 算法：从第十五题更改过来，找到绝对值由小变大的那个位置
	// 最小值只可能出现在这些点之中

	// 排序并定义一些变量
	var n int = len(nums)
	var distance int = 1<<63 - 1
	var result int
	var temp int = distance
	var nextTemp int = distance
	var isPositive1 bool
	var isPositive2 bool
	var abs = func(num int) (int, bool) {
		if num < 0 {
			return -num, false
		}
		return num, true
	}

	sort.Ints(nums)

	// 第一重循环
	for first := 0; first < n; first++ {
		// 排除相同的数
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := n - 1
		// 双指针，第一个
		for second := first + 1; second < third; second++ {

			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// distanceTemp := 1<<63 - 1
			for second < third {
				temp, isPositive1 = abs(nums[second] + nums[third] + nums[first] - target)
				// distanceTemp = temp
				nextTemp, isPositive2 = abs(nums[second] + nums[third-1] + nums[first] - target)
				if temp < nextTemp || (temp == nextTemp && isPositive1 != isPositive2) {
					break
				}
				third--
			}

			if temp < distance {
				distance = temp
				if second == third {
					result = nums[second] + nums[third+1] + nums[first]
					// fmt.Println("restlt:",result)
				} else {
					result = nums[second] + nums[third] + nums[first]
					// fmt.Println("restlt:",result)
				}
			}
		}
	}

	return result
}

// @lc code=end

