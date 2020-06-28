/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 *
 * https://leetcode-cn.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (56.95%)
 * Likes:    332
 * Dislikes: 0
 * Total Accepted:    73.8K
 * Total Submissions: 129.5K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * 给定两个二叉树，编写一个函数来检验它们是否相同。
 * 
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 * 
 * 示例 1:
 * 
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 * 
 * ⁠       [1,2,3],   [1,2,3]
 * 
 * 输出: true
 * 
 * 示例 2:
 * 
 * 输入:      1          1
 * ⁠         /           \
 * ⁠        2             2
 * 
 * ⁠       [1,2],     [1,null,2]
 * 
 * 输出: false
 * 
 * 
 * 示例 3:
 * 
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 * 
 * ⁠       [1,2,1],   [1,1,2]
 * 
 * 输出: false
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

 //主要是写逻辑写的烦
func isSameTree(p *TreeNode, q *TreeNode) bool {

	que1:=make([]*TreeNode,0)
	que2:=make([]*TreeNode,0)
	
	//对根节点进行判断
	if p!=nil && q!=nil && p.Val==q.Val {
		que1=append(que1,p)
		que2=append(que2,q)
	}else if p==nil && q ==nil{
		return true
	}else{
		return false
	}

	//层次便利，对没一层都进行判断
	for len(que1)!=0{
		tempQue1:=make([]*TreeNode,0)
		tempQue2:=make([]*TreeNode,0)

		for i:=0;i<len(que1);i++{

			if que1[i].Left !=nil && que2[i].Left!=nil{
				if que1[i].Left.Val !=que2[i].Left.Val{
					return false
				} 
				tempQue1=append(tempQue1,que1[i].Left)
				tempQue2=append(tempQue2,que2[i].Left)
			}else if que1[i].Left !=nil && que2[i].Left==nil{
				return false
			}else if que1[i].Left ==nil && que2[i].Left!=nil{
				return false
			}



			if que1[i].Right !=nil && que2[i].Right!=nil{
				if que1[i].Right.Val !=que2[i].Right.Val{
					return false
				} 
				tempQue1=append(tempQue1,que1[i].Right)
				tempQue2=append(tempQue2,que2[i].Right)
			}else if que1[i].Right !=nil && que2[i].Right==nil{
				return false
			}else if que1[i].Right ==nil && que2[i].Right!=nil{
				return false
			}
		}

		
		que1=tempQue1
		que2=tempQue2
		
	}

	return true

}
// @lc code=end

