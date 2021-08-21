/*
 * @Author: your name
 * @Date: 2021-08-21 15:21:14
 * @LastEditTime: 2021-08-21 19:12:05
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /leetcode-go/array/40.组合总和-ii.go
 */
/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (63.00%)
 * Likes:    670
 * Dislikes: 0
 * Total Accepted:    187.1K
 * Total Submissions: 297.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * 1
 * 1
 *
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var freq [][2]int
	for _, item := range candidates {
		if freq == nil || freq[len(freq)-1][0] != item {
			freq = append(freq, [2]int{item, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	col := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(int, int)
	dfs = func(left int, index int) {
		if left == 0 {
			ans = append(ans, append([]int{}, col...))
			return
		}
		if index == len(freq) || left < freq[index][0] {
			return
		}

		dfs(left, index+1)

		max := min(left/freq[index][0], freq[index][1])
		for i := 1; i <= max; i++ {
			col = append(col, freq[index][0])
			dfs(left-i*freq[index][0], index+1)
		}
		col = col[:len(col)-max]
	}

	dfs(target, 0)
	return ans
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}

// @lc code=end

