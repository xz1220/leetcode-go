/*
 * @Author: your name
 * @Date: 2021-08-18 23:46:05
 * @LastEditTime: 2021-08-18 23:46:18
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /leetcode-go/graph/207.课程表.go
 */
/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (54.53%)
 * Likes:    909
 * Dislikes: 0
 * Total Accepted:    131.7K
 * Total Submissions: 241.5K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi]
 * ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * prerequisites[i].length == 2
 * 0 i, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	inOrder := make([]int, numCourses)

	// 初始化边和入度
	for _, item := range prerequisites {
		edges[item[1]] = append(edges[item[1]], item[0])
		inOrder[item[0]]++
	}

	// 初始化队列
	q := make([]int, 0)
	for index, in := range inOrder {
		if in == 0 {
			q = append(q, index)
		}
	}

	//
	var ans []int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, e := range edges[node] {
			if inOrder[e] > 0 {
				inOrder[e]--
				if inOrder[e] == 0 {
					q = append(q, e)
				}
			}
		}
		ans = append(ans, node)
	}

	if len(ans) == numCourses {
		return true
	}
	return false
}

// @lc code=end

