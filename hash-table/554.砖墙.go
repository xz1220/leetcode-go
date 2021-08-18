/*
 * @lc app=leetcode.cn id=554 lang=golang
 *
 * [554] 砖墙
 *
 * https://leetcode-cn.com/problems/brick-wall/description/
 *
 * algorithms
 * Medium (49.92%)
 * Likes:    241
 * Dislikes: 0
 * Total Accepted:    43.1K
 * Total Submissions: 86.3K
 * Testcase Example:  '[[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]'
 *
 * 你的面前有一堵矩形的、由 n 行砖块组成的砖墙。这些砖块高度相同（也就是一个单位高）但是宽度不同。每一行砖块的宽度之和相等。
 * 
 * 你现在要画一条 自顶向下 的、穿过 最少
 * 砖块的垂线。如果你画的线只是从砖块的边缘经过，就不算穿过这块砖。你不能沿着墙的两个垂直边缘之一画线，这样显然是没有穿过一块砖的。
 * 
 * 给你一个二维数组 wall ，该数组包含这堵墙的相关信息。其中，wall[i] 是一个代表从左至右每块砖的宽度的数组。你需要找出怎样画才能使这条线
 * 穿过的砖块数量最少 ，并且返回 穿过的砖块数量 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：wall = [[1],[1],[1]]
 * 输出：3
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == wall.length
 * 1 
 * 1 
 * 1 
 * 对于每一行 i ，sum(wall[i]) 是相同的
 * 1 
 * 
 * 
 */

// @lc code=start
// 妙蛙
func leastBricks(wall [][]int) int {
    dic := make(map[int]int, 0)

    for _, layer := range wall {
        sum := 0
        for _, item := range layer[:len(layer) -1] {
            sum += item
            dic[sum] ++
        }
    }

    var maxSide int
    for _, val := range dic {
        if val > maxSide {
            maxSide = val
        }
    }

    return len(wall) - maxSide
}
// @lc code=end

