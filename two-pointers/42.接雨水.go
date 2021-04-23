/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (53.85%)
 * Likes:    2017
 * Dislikes: 0
 * Total Accepted:    192.7K
 * Total Submissions: 355.7K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == height.length
 * 0 
 * 0 
 * 
 * 
 */

// @lc code=start
func trap(height []int) int {
    // 算法：双指针法
    // 维护两个指针，left_max right_max
    // 记录两者之间的最小值 minimize
    // 当左指针小于等于右指针时 移动左指针
    // 当左指针大于右指针时 移动右指针
    // 每移动一个位置 记录差值并累加

    n:= len(height)
    if n== 0 || n == 1 {
        return 0
    }
    left_max, right_max := height[0], height[n-1]
    minimize := 0
    result := 0

    for i,j := 0, n-1 ; i != j ; {
        if left_max> right_max {
            minimize = right_max
        }else {
            minimize = left_max
        }

        if height[i] <= height[j] {
            i++
            if height[i] > left_max {
                left_max = height[i]
            }
            if height[i] < minimize {
                result += (minimize - height[i])
            }
        }else {
            j--
            if height[j] > right_max {
                right_max = height[j]
            }
            if height[j] < minimize {
                result += (minimize - height[j])
            }
        }
    }

    return result
}

func trap(height []int) int {
    if len(height) <= 2 {
        return 0
    }
    ans :=0
    l, lbig, r, rbig := 0, height[0], len(height) -1, height[len(height) -1]

    for l < r {
        if lbig <= rbig {
            l ++
            if height[l] <= lbig {
                ans += (lbig - height[l])
            }else {
                lbig = height[l]
            }
        }else {
            r -- 
            if height[r] <= rbig {
                ans += (rbig - height[r])
            }else {
                rbig = height[r]
            }
        }
    }
    return ans
}
// @lc code=end

