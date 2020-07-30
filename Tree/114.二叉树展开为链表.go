/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (68.09%)
 * Likes:    298
 * Dislikes: 0
 * Total Accepted:    33K
 * Total Submissions: 48.5K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给定一个二叉树，原地将它展开为链表。
 * 
 * 例如，给定二叉树
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 * 
 * 将其展开为：
 * 
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
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

 //前序便利(注意是要在原地便利，分情况讨论即可)
func flatten(root *TreeNode)  {

	var dfs func(*TreeNode) *TreeNode

	dfs=func(root *TreeNode) *TreeNode{

		if root==nil{
			return nil
		}

		if  root.Left==nil && root.Right==nil{
			return root
		}

		if root.Left!=nil && root.Right==nil{
			root.Right=root.Left
			root.Left=nil
			return dfs(root.Right)
		}

		if root.Left==nil && root.Right!=nil{
			return dfs(root.Right)
		}

		if root.Left!=nil && root.Right!=nil{
			left:=dfs(root.Left)
			right:=dfs(root.Right)
			left.Right=root.Right
			root.Right=root.Left
			root.Left=nil
			return right
		}
		return nil
	}

	dfs(root)
}
// @lc code=end

