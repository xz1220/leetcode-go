/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.62%)
 * Likes:    1043
 * Dislikes: 0
 * Total Accepted:    316.6K
 * Total Submissions: 490K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 * 
 * 示例 1:
 * 
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 * 
 * 说明: 
 * 
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 * 
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
    var heapify func(int,int)
    var buildMax func(int)
    
    heapify = func(i, length int) {
        l, r := 2*i +1, 2*i +2
        largest := i

        if l < length && nums[l] > nums[largest] {
            largest = l
        }
        if r < length && nums[r] > nums[largest] {
            largest = r
        }

        if largest != i {
            nums[largest], nums[i] = nums[i], nums[largest]
            heapify(largest, length)
        }
    }

    buildMax = func(length int) {
        for i:=  length /2 ; i >= 0 ;i -- {
            heapify(i, length)
        }
    }

    buildMax(len(nums))
    for i:= len(nums) -1 ; i >= len(nums) -k +1 ; i -- {
        nums[i], nums[0] = nums[0], nums[i]
        heapify(0, i)
    }
    return nums[0]
}
// @lc code=end

