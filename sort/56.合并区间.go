/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (42.67%)
 * Likes:    662
 * Dislikes: 0
 * Total Accepted:    156.2K
 * Total Submissions: 361.5K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 * 
 * 
 * 示例 2:
 * 
 * 输入: intervals = [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 * 
 * 注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * intervals[i][0] <= intervals[i][1]
 * 
 * 
 */

// @lc code=start
func merge(intervals [][]int) [][]int {

	sort.Slice(intervals, func(i,j int) bool {
		return intervals[i][0] < intervals[j][0] /*满足接口要求，要注意是小于号而不能是小于等于.*/
		/*TODO：为什么在小于等于时，无法保证排序的稳定性？*/
    })

    var returnIntervals [][]int

    for i:=0 ; i<len(intervals) ;i++ {
        if len(returnIntervals) == 0 /* 数组元素的第一个加入结果*/ || intervals[i][0] > returnIntervals[len(returnIntervals)-1][1]{
            returnIntervals = append(returnIntervals,intervals[i])
        }else if intervals[i][1] > returnIntervals[len(returnIntervals)-1][1]{
            returnIntervals[len(returnIntervals)-1][1] = intervals[i][1]
        }
    }

    return returnIntervals

}
// @lc code=end

