/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (64.61%)
 * Likes:    208
 * Dislikes: 0
 * Total Accepted:    48.5K
 * Total Submissions: 75K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其自底向上的层次遍历为：
 * 
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
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
 //非递归可以使用栈来做，打上标记啥的，先全部节点按层次顺序入栈，接着出栈
func levelOrderBottom(root *TreeNode) [][]int {
	res:=[][]int{}

	var dfs func(*TreeNode,int)
	dfs=func(root *TreeNode,level int){
		if root==nil{
			return
		}

		//出现新的level
		if level>=len(res){
			res=append([][]int{[]int{}},res...)
		}
		n:=len(res)
		res[n-level-1]=append(res[n-level-1],root.Val)

		dfs(root.Left,level+1)
		dfs(root.Right,level+1)
	}

	dfs(root,0)

	return res

}
// @lc code=end

