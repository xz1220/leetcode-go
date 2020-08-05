/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (70.66%)
 * Likes:    257
 * Dislikes: 0
 * Total Accepted:    62.1K
 * Total Submissions: 87.8K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
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
 * 输出: [3,2,1]
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
func postorderTraversal(root *TreeNode) []int {
	var top,pre *TreeNode
	var list []int
	stack:=make([]*TreeNode,0)

	//如果root==nil 则append会panic
	if root!=nil{
		stack=append(stack,root)
	}

	//
	for len(stack) != 0{
		top=stack[len(stack)-1]
		if (top.Left==nil && top.Right==nil) || (pre!=nil && (pre==top.Left || pre==top.Right)){
			list=append(list,top.Val)
			pre=top
			stack=stack[:len(stack)-1]
			
		}else{
			if top.Right!=nil{
				stack=append(stack,top.Right)
			}
			if top.Left!=nil{
				stack=append(stack,top.Left)
			}
		}
	}

	return list

}
// @lc code=end

