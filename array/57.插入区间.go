/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode-cn.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (37.29%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    32.4K
 * Total Submissions: 85.2K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
 *
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 *
 *
 *
 * 示例 1：
 *
 * 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出：[[1,5],[6,9]]
 *
 *
 * 示例 2：
 *
 * 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出：[[1,2],[3,10],[12,16]]
 * 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 *
 *
 *
 *
 * 注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。
 *
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	// initiate the returnSlice
	var index int
	var returnSlice [][]int

	// loop, find the insorting postion-index
	if len(intervals) == 0 {
		return append(returnSlice, newInterval)
	}

	for index = 0; index < len(intervals); index++ {
		if intervals[index][0] > newInterval[0] {
			break
		}
	}

	if index == len(intervals) {
		intervals = append(intervals, newInterval)
	}

	// loop the intervals and put the merged interval into returnSlice
	// when i equals index, do the same thing to newInterval
	// Then i - 1

	var insertFn func([]int)

	insertFn = func(slice []int) {
		if len(returnSlice) == 0 || slice[0] > returnSlice[len(returnSlice)-1][1] {
			returnSlice = append(returnSlice, slice)
		} else if slice[1] > returnSlice[len(returnSlice)-1][1] {
			returnSlice[len(returnSlice)-1][1] = slice[1]
		}
	}

	flag := false
	for i := 0; i < len(intervals); i++ {
		if i == index && !flag {
			insertFn(newInterval)
			i, flag = i-1, true
		} else {
			insertFn(intervals[i])
		}
	}

	return returnSlice
}

// @lc code=end

