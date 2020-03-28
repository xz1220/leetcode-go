/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (63.93%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    22.2K
 * Total Submissions: 34.7K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 * 
 * 示例:
 * 
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 * 
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
func rightSideView(root *TreeNode) []int {
	que:=make([]*TreeNode,0)
	var list []int
	if root!=nil{
		que=append(que,root)
	}

	for len(que)!=0{
		temp:=make([]*TreeNode,0)
		for _,node:=range que{
			if node.Left!=nil{
				temp=append(temp,node.Left)
			}
			if node.Right!=nil{
				temp=append(temp,node.Right)
			}
		}
		list=append(list,que[len(que)-1].Val)
		que=temp
	}

	return list

}
// @lc code=end

