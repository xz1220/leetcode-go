/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (61.30%)
 * Likes:    419
 * Dislikes: 0
 * Total Accepted:    96.2K
 * Total Submissions: 156.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 * 
 * 例如:
 * 给定二叉树: [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其层次遍历结果：
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 * 
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
func levelOrder(root *TreeNode) [][]int {
	que:=make([]*TreeNode,0)
	var list [][]int //可能会出错，不确定是否这样可以
	if root!=nil{
		que=append(que,root)
	}

	for len(que)!=0{
		temp:=make([]*TreeNode,0)
		var tempList =[]int{}
		for _,node:=range que{
			tempList=append(tempList,node.Val)
			if node.Left!=nil{
				temp=append(temp,node.Left)
			}

			if node.Right!=nil{
				temp=append(temp,node.Right)
			}
		}
		list=append(list,tempList)
		que=temp
	}
	return list

}
// @lc code=end

