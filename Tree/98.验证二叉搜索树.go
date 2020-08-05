/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (29.54%)
 * Likes:    478
 * Dislikes: 0
 * Total Accepted:    84.9K
 * Total Submissions: 287.1K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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
// func isValidBST(root *TreeNode) bool {

// 	var flag=true
// 	var dfs func(*TreeNode)(int,int)

// 	if root==nil{
// 		return true
// 	}

// 	//因为不确定是否有负数的存在
// 	//其实也不需要在return的时候那么麻烦，因为我们可以假定if条件语句执行通过
// 	dfs=func(root *TreeNode)(int,int){
// 		if root.Left==nil && root.Right==nil{
// 			return root.Val,root.Val
// 		}
// 		if root.Left!=nil && root.Right==nil{
// 			leftmin,leftmax:=dfs(root.Left)
// 			if leftmax>=root.Val{
// 				flag=false
// 			}

// 			return min(root.Val,leftmin),max(root.Val,leftmax)
// 		}
// 		if root.Right!=nil && root.Left==nil{
// 			rightmin,rightmax:=dfs(root.Right)
// 			if root.Val>=rightmin{
// 				flag=false
// 			}

// 			return min(root.Val,rightmin),max(root.Val,rightmax)

// 		}
// 		leftmin,leftmax:=dfs(root.Left)
// 		rightmin,rightmax:=dfs(root.Right)

// 		if leftmax>=root.Val || rightmin<=root.Val{
// 			flag=false
// 		}

// 		return min(leftmin,min(root.Val,rightmin)),max(leftmax,max(root.Val,rightmax))

// 	}

// 	dfs(root)

// 	return flag

// }

// func max(a,b int)int{
// 	if a>b{
// 		return a
// 	}else{
// 		return b
// 	}
// }

// func min(a,b int)int{
// 	if a<b{
// 		return a
// 	}
// 	return b
// }

// 官方题解1：如果是二叉搜索树，则满足各个节点都要在一个区间之内
// 非常优雅

// func isValidBST(root *TreeNode) bool {

// 	return helper(root,math.MinInt64, math.MaxInt64)

// }

// func helper(node *TreeNode,low, high int) bool{
// 	if node == nil {
// 		return true
// 	}

// 	if node.Val <= low || node.Val >= high {
// 		return false
// 	}

// 	return helper(node.Left,low,node.Val) && helper(node.Right,node.Val,high)
// }

// 官方题解2： 二叉搜索树的中序遍历是一个递增的数列 可以使用非递归实现

// func isValidBST(root *TreeNode) bool {
// 	stack := []*TreeNode{}
// 	// var pre *TreeNode=root
// 	var temp int = math.MinInt64

// 	for root!=nil || len(stack)!=0{
// 		if root!=nil{
// 			stack=append(stack,root)
// 			root = root.Left
// 		}else{
// 			root =stack[len(stack)-1]
// 			stack=stack[:len(stack)-1]
// 			if root.Val <=temp {
// 				return false
// 			}
// 			temp = root.Val
// 			root = root.Right
// 		}
// 	}
// 	return true
// }
// Accepted
// Your runtime beats 78.32 % of golang submissions
// Your memory usage beats 80 % of golang submissions (5.5 MB)

// 另一种中序遍历的实现 
func isValidBST(root *TreeNode) bool {
    stack := []*TreeNode{}
    inorder := math.MinInt64
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Val <= inorder {
            return false
        }
        inorder = root.Val
        root = root.Right
    }
    return true
}
// Accepted
// Your runtime beats 78.32 % of golang submissions
// Your memory usage beats 80 % of golang submissions (5.5 MB)

// @lc code=end

