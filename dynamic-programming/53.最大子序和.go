/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (52.17%)
 * Likes:    2981
 * Dislikes: 0
 * Total Accepted:    436.9K
 * Total Submissions: 816.5K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
 * 输出：6
 * 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1]
 * 输出：1
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [0]
 * 输出：0
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：nums = [-1]
 * 输出：-1
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：nums = [-100000]
 * 输出：-100000
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10^5 
 * 
 * 
 * 
 * 
 * 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 * 
 */

// @lc code=start
// func maxSubArray(nums []int) int {
//     /* f(i) = max(f(i-1)+ nums[i], nums[i]) 
//      * f(i) 为以右端点为起始点的最大和连续子数组
//      *
//     */
//     var max int = nums[0]
//     for i:= 1; i < len(nums) ; i++ {
//         if nums[i-1] + nums[i] > nums[i] {
//             nums[i] += nums[i-1]
//         }

//         if nums[i] > max {
//             max = nums[i]
//         }
//     }

//     return max
// }


// 线段树搜索，分治算法
type stateNums struct {
    totalSum int    // 整个子数组的总和
    leftSum int
    rightSum int
    maxSum int
}

func max(num1 , num2 int) int {
    if num1 > num2 {
        return num1
    }
    return num2
}

func putUp(left, right stateNums) stateNums {
    totalSum := left.totalSum + right.totalSum
    leftSum := max(left.leftSum, left.totalSum + right.leftSum)
    rightSum := max(right.rightSum, left.rightSum + right.totalSum)
    maxSum := max(max(left.maxSum, right.maxSum), left.rightSum + right.leftSum)
    return stateNums{totalSum, leftSum, rightSum, maxSum}
}

func get(nums []int, left, right int) stateNums {
    if left == right {
        return stateNums{nums[left], nums[left], nums[left], nums[left]}
    }

    middleIndex := (left + right) >>1
    leftSub := get(nums, left, middleIndex)
    rightSub := get(nums, middleIndex+1, right)

    return putUp(leftSub, rightSub)
}


func maxSubArray(nums []int) int {
    return get(nums, 0, len(nums) -1).maxSum
}

func maxSubArray(nums []int) int {
    ans := nums[0]
    for i := range nums {
        if i == 0 {
            continue
        }
        nums[i] = max(nums[i], nums[i] + nums[i-1])
        ans = max(nums[i], ans)
    }
    return ans
}

func max(num1, num2 int) int {
    if num1 > num2 {
        return num1
    }
    return num2
}
// @lc code=end

