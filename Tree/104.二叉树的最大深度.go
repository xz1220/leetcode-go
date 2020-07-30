/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (72.63%)
 * Likes:    490
 * Dislikes: 0
 * Total Accepted:    147.2K
 * Total Submissions: 202.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 * 
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最大深度 3 。
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

 
 //有多种方法，深度有限搜索一下，或者进行一下层次遍历
 //dfs
//  func maxDepth(root *TreeNode) int {
// 	var maxDepth=0
// 	var depth=0
// 	var dfs func(*TreeNode,int)

// 	dfs=func(root *TreeNode,depth int){
// 		if root==nil{
// 			return
// 		}

// 		depth+=1
// 		left:=depth
// 		Right:=depth
// 		if depth>maxDepth{
// 			maxDepth=depth
// 		}
// 		dfs(root.Left,left)
// 		dfs(root.Right,Right)
// 	}

// 	dfs(root,depth)

// 	return maxDepth
// }


//非递归版本，运行速度较慢，但是内存占用率低
func maxDepth(root *TreeNode) int {

	var depth=0

	que:=make([]*TreeNode,0)
	if root!=nil{
		que=append(que,root)
	}

	for len(que)!=0{
		depth+=1
		tempQue:=make([]*TreeNode,0)

		for _,node:=range que{
			if node.Left!=nil{
				tempQue=append(tempQue,node.Left)
			}

			if node.Right!=nil{
				tempQue=append(tempQue,node.Right)
			}
		}

		que=tempQue
		
	}

	return depth

}
// @lc code=end

