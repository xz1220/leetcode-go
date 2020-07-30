/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (64.58%)
 * Likes:    378
 * Dislikes: 0
 * Total Accepted:    52.3K
 * Total Submissions: 80.9K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0 || len(preorder)!=len(inorder){
		return nil
	}
	root:=&TreeNode{
		Val:preorder[0],
	}

	index:=indexOf(root.Val,inorder)
	root.Left=buildTree(preorder[1:index+1],inorder[:index])
	root.Right=buildTree(preorder[index+1:],inorder[index+1:])

	return root

}

func indexOf(value int,inorder []int) int{
	for index,val:=range inorder{
		if val==value{
			return index
		}
	}

	panic("匹配错误")
}
// @lc code=end

