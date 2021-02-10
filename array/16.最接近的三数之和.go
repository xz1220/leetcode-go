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
	// 算法：从第十五题更改过来
	// 要寻找三数之后最接近这个点的数，即找出相减后绝对值最小的点
	// 1. 数组任意三数之和大于或者小于 target ， 找到最接近的即可
	// 2. 任意三数之和有的大于有的小于，双指针法逼近

	// 排序并定义一些变量
	var n int = len(nums)
	var distance int = 1<<63 - 1
	var result int                        //
	var temp int = distance               // 查找绝对值变化的点，第二重循环
	var nextTemp int = distance           //
	var isPositive1 bool                  // 判断是否为正数
	var isPositive2 bool                  //
	var abs = func(num int) (int, bool) { // 返回绝对值和是否为正数
		if num < 0 {
			return -num, false
		}
		return num, true
	}

	sort.Ints(nums) // 排序

	// 第一重循环
	for first := 0; first < n; first++ {
		// 排除相同的数
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		// 双指针法
		third := n - 1
		for second := first + 1; second < third; second++ {
			// 排除相同的数
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 寻找距离突变的点：
			// 1. 无突变的点，不发生break
			// 2. 有绝对值突变的点，break
			for second < third {
				temp, isPositive1 = abs(nums[second] + nums[third] + nums[first] - target)
				nextTemp, isPositive2 = abs(nums[second] + nums[third-1] + nums[first] - target)
				if temp < nextTemp || (temp == nextTemp && isPositive1 != isPositive2) {
					break
				}
				third--
			}

			// 判断是否需要更新
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

// @lc code=end

