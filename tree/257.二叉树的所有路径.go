/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (64.05%)
 * Likes:    413
 * Dislikes: 0
 * Total Accepted:    87.9K
 * Total Submissions: 132.5K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归版本
func binaryTreePaths(root *TreeNode) []string {
	var list []string
	var dfs func(*TreeNode, string)

	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			path = path + strconv.Itoa(root.Val)
			list = append(list, path)
			return
		}
		path = path + strconv.Itoa(root.Val)
		if root.Left != nil {
			dfs(root.Left, path+"->")
		}
		if root.Right != nil {
			dfs(root.Right, path+"->")
		}
	}

	dfs(root, "")
	return list
}

// @lc code=end

