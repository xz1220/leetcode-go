/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (36.34%)
 * Likes:    309
 * Dislikes: 0
 * Total Accepted:    59.4K
 * Total Submissions: 159.3K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 * 
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 * 
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 * 
 * 示例 1:
 * 
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 * 
 * 进阶:
 * 
 * 
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 * 
 * 
 */

// @lc code=start
func search(nums []int, target int) bool {
    left, right, middle := 0, len(nums) -1, (len(nums) -1)/2
    ans := false

    for left < right {
        if nums[left] == nums[middle] {
           ans = search(nums[left:middle+1],target) || search(nums[middle+1:], target)
           break
        }else if nums[left] < nums[middle] {
            if nums[left] <= target && target <= nums[middle] {
                right = middle
                middle = (left + right) /2
            }else {
                left = middle + 1
                middle = (left + right) /2
            }
        }else {
            if nums[middle+1] <= target && target <= nums[right] {
                left = middle +1
                middle = (left + right) /2
            }else {
                right = middle
                middle = (left + right) /2
            }
        }
    }

    if nums[left] != target {
        ans = ans || false
    }else {
        ans = ans || true
    }

    return ans
}
// @lc code=end

