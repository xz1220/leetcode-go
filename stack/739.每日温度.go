/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
    // 单调栈
    stack := make([]int, 0)
    ans := make([]int, len(temperatures))

    for i:=0 ; i< len(temperatures); i++ {
        if len(stack) == 0 {
            stack = append(stack, i)
            continue
        }

        if temperatures[i] <= temperatures[stack[len(stack) -1]] {
            stack = append(stack, i)
        }else {
            var j int
            for j = len(stack) -1; j >=0 ; j -- {
                if temperatures[stack[j]] >= temperatures[i] {
                    break
                }
                ans[stack[j]] = i - stack[j]
            }
            stack = stack[:j+1]
            stack = append(stack, i)
        }
    }

    return ans
}

//  标准解法

func dailyTemperatures(temperatures []int) []int {
    length := len(temperatures)
    ans := make([]int, length)
    stack := []int{}
    for i := 0; i < length; i++ {
        temperature := temperatures[i]
        for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
            prevIndex := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans[prevIndex] = i - prevIndex
        }
        stack = append(stack, i)
    }
    return ans
}

// @lc code=end

