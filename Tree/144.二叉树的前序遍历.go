/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (64.68%)
 * Likes:    228
 * Dislikes: 0
 * Total Accepted:    83.4K
 * Total Submissions: 128.7K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
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
 * 输出: [1,2,3]
 * 
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

 //非递归版本
func preorderTraversal(root *TreeNode) []int {
	var list []int
	var pre *TreeNode=root
	stack:=make([]*TreeNode,0)

	for pre!=nil || len(stack)!=0{
		if pre!=nil{
			stack=append(stack,pre)
			list=append(list,pre.Val)
			pre=pre.Left
		}else{
			node:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			pre=node.Right
		}

	}
	
	return list
}
// @lc code=end

