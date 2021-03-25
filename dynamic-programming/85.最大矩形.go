/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 *
 * https://leetcode-cn.com/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (51.59%)
 * Likes:    850
 * Dislikes: 0
 * Total Accepted:    72.4K
 * Total Submissions: 140.1K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix =
 * [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
 * 输出：6
 * 解释：最大矩形如上图所示。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix = []
 * 输出：0
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：matrix = [["0"]]
 * 输出：0
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：matrix = [["1"]]
 * 输出：1
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：matrix = [["0","0"]]
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * rows == matrix.length
 * cols == matrix[0].length
 * 0 
 * matrix[i][j] 为 '0' 或 '1'
 * 
 * 
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0{
        return 0
    }
    m, n := len(matrix), len(matrix[0])
    tempArray := make([]int, n)
    var maxArea int

    for i:=0 ; i< m ; i++ {
        // 更新每一行
        for j:= 0 ; j < n ; j++ {
            if matrix[i][j] == '1' {			// 注意byte的比较方式
                tempArray[j] += 1
            }else {
                tempArray[j] = 0
            }
        }

        // fmt.Println("Temparray: ", tempArray)

        tempAns := 0
        stack := make([]int, n)
        left, right := make([]int, n), make([]int, n)

        // 单调栈
        for j:=0; j< n ; j++ {
            for len(stack) > 0 && tempArray[stack[len(stack) -1]]  >= tempArray[j] {
                stack = stack[:len(stack) -1]
            }

            if len(stack) == 0 {
                left[j] = -1
            }else {
                left[j] = stack[len(stack) -1]
            }
            stack = append(stack, j)
        }
        // fmt.Println("left: ", left)

        stack = []int{}
        for j:=n-1 ; j >=0 ; j -- {
            for len(stack) > 0 && tempArray[stack[len(stack) -1]]  >= tempArray[j] {
                stack = stack[:len(stack) -1]
            }

            if len(stack) == 0 {
                right[j] = n
            }else {
                right[j] = stack[len(stack) -1]
            }
            stack = append(stack, j)
        }
        // fmt.Println("Right: ", right)

        for j:= 0 ; j < n ; j ++ {
            tempAns = max(tempAns, (right[j] - left[j] -1)*tempArray[j])
        }

        // fmt.Println("Current MaxArea: ", maxArea, "  tempAns:",tempAns)
        if tempAns > maxArea {
            maxArea = tempAns
        }
    }
    return maxArea
}

func max(nums1, nums2 int) int {
    if nums1 > nums2 {
        return nums1
    }
    return nums2
}
// @lc code=end

