/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 *
 * https://leetcode-cn.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (74.53%)
 * Likes:    396
 * Dislikes: 0
 * Total Accepted:    65.7K
 * Total Submissions: 88.1K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * 翻转一棵二叉树。
 * 
 * 示例：
 * 
 * 输入：
 * 
 * ⁠    4
 * ⁠  /   \
 * ⁠ 2     7
 * ⁠/ \   / \
 * 1   3 6   9
 * 
 * 输出：
 * 
 * ⁠    4
 * ⁠  /   \
 * ⁠ 7     2
 * ⁠/ \   / \
 * 9   6 3   1
 * 
 * 备注:
 * 这个问题是受到 Max Howell 的 原问题 启发的 ：
 * 
 * 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
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
func invertTree(root *TreeNode) *TreeNode {

	que:=make([]*TreeNode,0)
	if root!=nil{
		que=append(que,root)
	}

	for len(que)!=0{
		tempQue:=make([]*TreeNode,0)

		for _,node:=range que{
			if node.Left!=nil{
				tempQue=append(tempQue,node.Left)
			}

			if node.Right!=nil{
				tempQue=append(tempQue,node.Right)
			}

			node.Left,node.Right=node.Right,node.Left
		}

		que=tempQue
	}

	return root

}
// @lc code=end

