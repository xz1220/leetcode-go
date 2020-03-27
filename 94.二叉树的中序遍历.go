/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (70.61%)
 * Likes:    438
 * Dislikes: 0
 * Total Accepted:    128.3K
 * Total Submissions: 181.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 * 
 * 输出: [1,3,2]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
// func inorderTraversal(root *TreeNode) []int {
// 	var list []int
// 	var dfs func(*TreeNode)

// 	dfs=func(root *TreeNode){
// 		if root==nil{
// 			return
// 		}
// 		dfs(root.Left)
// 		list=append(list,root.Val)
// 		dfs(root.Right)
// 	}

// 	dfs(root)

// 	return list

// }

// 非递归版本
func inorderTraversal(root *TreeNode) []int {
	var list []int
	var pre *TreeNode=root
	stack:=make([]*TreeNode,0)

	for pre!=nil || len(stack)!=0{
		if pre!=nil{
			stack=append(stack,pre)
			pre=pre.Left
		}else{
			node:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			list=append(list,node.Val)
			pre=node.Right
		}

	}
	
	return list

}
// @lc code=end

