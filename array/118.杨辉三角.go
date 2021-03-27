/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (69.76%)
 * Likes:    463
 * Dislikes: 0
 * Total Accepted:    159.3K
 * Total Submissions: 226.5K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 * 
 * 
 * 
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 * 
 * 示例:
 * 
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 */

// @lc code=start
func generate(numRows int) [][]int {
    ans := make([][]int, numRows)
    for row:= 1 ; row <= numRows ; row ++ {
        temp:= make([]int, row)
        middle := 0
        if row % 2 == 1 {
            middle = row/2 +1
        }else {
            middle = row/2
        }

        for index:= 0 ; index < middle ; index ++ {
            temp[index] = cal(row -1, index)
        }
        for index:= middle ; index < row ; index ++ {
            if row % 2 == 1 {
                temp[index] = temp[middle - (index - middle) -2]
            }else {
                temp[index] = temp[middle - (index - middle) -1]
            }
        }
        ans[row -1] = temp
    }
    return ans
}

func cal(n, m int) int {
    result1, result2 := 1, 1
    for i:= n ; i > (n - m) ; i-- {
        result1 *= i
    }

    for i:= m; i >0 ; i-- {
        result2 *=i
    }

    return result1/result2
}
// @lc code=end

