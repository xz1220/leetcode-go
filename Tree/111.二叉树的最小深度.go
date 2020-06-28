/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (41.84%)
 * Likes:    231
 * Dislikes: 0
 * Total Accepted:    61.5K
 * Total Submissions: 147K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 * 
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最小深度  2.
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


 //用非递归的方法做起来非常容易，其实就是层次遍历
func minDepth(root *TreeNode) int {

	var depth=0

	que:=make([]*TreeNode,0)
	if root!=nil{
		que=append(que,root)
	}

	for len(que)!=0{
		depth+=1
		tempQue:=make([]*TreeNode,0)

		for _,node:=range que{
			if node.Left==nil && node.Right==nil{
				goto end
			}

			if node.Left!=nil{
				tempQue=append(tempQue,node.Left)
			}
			if node.Right!=nil{
				tempQue=append(tempQue,node.Right)
			}

		}

		que=tempQue
		
	}
end:

	return depth
}
// @lc code=end

