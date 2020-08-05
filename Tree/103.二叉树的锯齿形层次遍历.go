/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (54.01%)
 * Likes:    166
 * Dislikes: 0
 * Total Accepted:    39.5K
 * Total Submissions: 73.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回锯齿形层次遍历如下：
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	que:=make([]*TreeNode,0)
	flag:=true
	var list [][]int
	if root!=nil{
		que=append(que,root)
	}

	for len(que)!=0{
		temp:=make([]*TreeNode,0)
		var tempList []int

		for  i:=0;i<len(que);i++{
			if flag{
				tempList=append(tempList,que[i].Val)
			}else{
				tempList=append(tempList,que[len(que)-i-1].Val)
				
			}
			
			if que[i].Left!=nil{
				temp=append(temp,que[i].Left)
			}

			if que[i].Right!=nil{
				temp=append(temp, que[i].Right)
			}
		}

		que=temp
		list=append(list,tempList)
		flag=!flag
	}

	return list

}
// @lc code=end

