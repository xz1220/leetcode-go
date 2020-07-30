/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}

func maxPathSum(root *TreeNode) int {
	if root==nil{
		return 0
	}

	//在这里不可以随意初始化，要考虑到只有一个节点的情况
	maxSum:=root.Val

	var dfs func(*TreeNode)int
	dfs=func(root *TreeNode)int{
		if root==nil{
			return 0
		}
		left:=max(0,dfs(root.Left))
		right:=max(0,dfs(root.Right))
		sum:=left+root.Val+right
		if sum>maxSum{
			maxSum=sum
		}
		return max(left,right)+root.Val
	}

	dfs(root)

	return maxSum

}
// @lc code=end

