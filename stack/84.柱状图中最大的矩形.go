/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (42.48%)
 * Likes:    1267
 * Dislikes: 0
 * Total Accepted:    133.9K
 * Total Submissions: 312K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 * 
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 * 
 * 
 * 
 * 
 * 
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 * 
 * 
 * 
 * 
 * 
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 * 
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
    // 单调栈
    n := len(heights)
    left , right := make([]int, n), make([]int, n)
    stack := make([]int, n)

    for i:= 0 ; i < n ; i++ {
        for len(stack) >0 && heights[stack[len(stack) - 1]] >= heights[i] {
            stack = stack[:len(stack) - 1]
        }

        if len(stack) == 0 {
            left[i] = -1
        }else {
            left[i] = stack[len(stack) -1]
        }
        stack = append(stack, i)
    }



    stack = []int{}
    for i := n -1 ; i >= 0 ; i-- {
        for len(stack) > 0 && heights[stack[len(stack) -1]] >= heights[i] {
            stack = stack[:len(stack) -1]
        }

        if len(stack) == 0 {
            right[i] = n
        }else {
            right[i] = stack[len(stack) -1]
        }
        stack = append(stack, i)
    }

    fmt.Println(left)
    fmt.Println(right)
    
    ans := 0
    for i:= 0 ; i < n ; i++ {
        ans = max(ans, (right[i] - left[i] -1)*heights[i])
    }
    return ans
}

func max(nums1, nums2 int) int {
    if nums1 > nums2 {
        return nums1
    }
    return nums2
}
// @lc code=end

