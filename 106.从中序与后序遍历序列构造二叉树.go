/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (67.50%)
 * Likes:    176
 * Dislikes: 0
 * Total Accepted:    29.1K
 * Total Submissions: 43.1K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	//如若不相等，则是错误
	if len(inorder)!=len(postorder){
		return nil
	}

	if len(inorder)==0{
		return nil
	}

	root:=&TreeNode{
		Val:postorder[len(postorder)-1],
	}

	index:=indexOf(root.Val,inorder)
	root.Left=buildTree(inorder[:index],postorder[:index])
	root.Right=buildTree(inorder[index+1:],postorder[index:len(postorder)-1])

	return root
}

func indexOf(value int,order []int) int{
	for index,val:=range order{
		if val==value{
			return index
		}
	}
	return -1
}
// @lc code=end

