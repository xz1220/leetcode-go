/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (40.33%)
 * Likes:    288
 * Dislikes: 0
 * Total Accepted:    53.9K
 * Total Submissions: 133.6K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 * 
 * 示例 1:
 * 
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 * 
 * 
 * 示例 2:
 * 
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 * 
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	index := [][]int{}
	flag := 0
	for k := 0; k < len(intervals); k++ {
		if intervals[flag][1] >= intervals[k][0] && intervals[flag][1] >= intervals[k][1] {
			fmt.Println("case 1")
			if k == len(intervals)-1 {
				index = append(index, []int{intervals[flag][0], intervals[flag][1]})

			}

		} else if intervals[flag][1] >= intervals[k][0] && intervals[flag][1] < intervals[k][1] {
			fmt.Println("case 2")
			intervals[flag][1] = intervals[k][1]
			if k == len(intervals)-1 {
				index = append(index, []int{intervals[flag][0], intervals[flag][1]})
			}
		} else if intervals[flag][1] < intervals[k][0] {
			fmt.Println("case 3")
			index = append(index, []int{intervals[flag][0], intervals[flag][1]})
			flag = k
			if k == len(intervals)-1 {
				index = append(index, []int{intervals[flag][0], intervals[flag][1]})
			}
		}
	}
	fmt.Println(index)

	return index
	

}
// @lc code=end

