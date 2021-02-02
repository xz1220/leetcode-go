/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (64.05%)
 * Likes:    2153
 * Dislikes: 0
 * Total Accepted:    358K
 * Total Submissions: 553.7K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 * 示例 2：
 *
 *
 * 输入：height = [1,1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：height = [4,3,2,1,4]
 * 输出：16
 *
 *
 * 示例 4：
 *
 *
 * 输入：height = [1,2,1]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * n = height.length
 * 2
 * 0
 *
 *
 */

// @lc code=start
func maxArea(height []int) int {
	// 算法：
	// 1. 初始化两个下标，一个指向头部，一个指向尾部； 初始化maxV
	// 2. 循环：数字小的那个下标向中间移动到下一个比他大的那个数字，计算体积并决定是否更新最大值，直到两个下标相遇

	var head, tail, maxV int
	head, tail = 0, len(height)-1
	maxV = (tail - head) * min(height[head], height[tail])

	for tail > head {
		if height[tail] >= height[head] {
			head = nextIndex(height, head, tail, true)
			tempV := (tail - head) * min(height[head], height[tail])
			if tempV > maxV {
				maxV = tempV
			}
		} else {
			tail = nextIndex(height, head, tail, false)
			tempV := (tail - head) * min(height[head], height[tail])
			if tempV > maxV {
				maxV = tempV
			}
		}
	}

	return maxV
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func nextIndex(nums []int, head, tail int, ToRight bool) int {
	if ToRight {
		for index := head; index < tail; index++ {
			if nums[index] > nums[head] {
				return index
			}
		}
		return tail
	} else {
		for index := tail; index > head; index-- {
			if nums[index] > nums[tail] {
				return index
			}
		}
		return head
	}
}

// @lc code=end

