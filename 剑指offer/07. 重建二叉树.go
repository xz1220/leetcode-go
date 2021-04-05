/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if preorder == nil || len(preorder) == 0 {
        return nil
    }
    node := &TreeNode{
        Val: preorder[0],
    }
    var index int 
    for ; index < len(inorder) ; index ++ {
        if inorder[index] == preorder[0] {
            break
        }
    }
    node.Left = buildTree(preorder[1:index +1], inorder[0:index])
    node.Right = buildTree(preorder[index+1:], inorder[index+1:])
    return node
}