/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (50.59%)
 * Likes:    678
 * Dislikes: 0
 * Total Accepted:    113.8K
 * Total Submissions: 224.9K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 说明:
 * 
 * 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
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
// func isSymmetric(root *TreeNode) bool {
// 	que1:=make([]*TreeNode,0)
// 	que2:=make([]*TreeNode,0)
	
// 	//对根节点进行判断
// 	if root!=nil {
// 		que1=append(que1,root)
// 		que2=append(que2,root)
// 	}

// 	//层次便利，对没一层都进行判断
// 	for len(que1)!=0{
// 		tempQue:=make([]*TreeNode,0)
// 		tempQue2:=make([]*TreeNode,0)

// 		for i:=0;i<len(que1);i++{

// 			if que[i].Left !=nil && que[len(que)-i-1]!=nil{
// 				if que[i].Left.Val !=que[len(que)-i-1].Val{
// 					return false
// 				} 
// 				tempQue1=append(tempQue1,que[i].Left)
// 				tempQue2=append(tempQue2,que[len(que)-i-1])
// 			}else if que[i].Left !=nil && que[len(que)-i-1]==nil{
// 				return false
// 			}else if que[i].Left ==nil && que[len(que)-i-1]!=nil{
// 				return false
// 			}



// 			if que[i].Right !=nil && que[len(que)-i-1].Left!=nil{
// 				if que[i].Right.Val !=que[len(que)-i-1].Left.Val{
// 					return false
// 				} 
// 				tempQue1=append(tempQue1,que[i].Right)
// 				tempQue2=append(tempQue2,que[len(que)-i-1].Left)
// 			}else if que[i].Right !=nil && que[len(que)-i-1].Left==nil{
// 				return false
// 			}else if que[i].Right ==nil && que[len(que)-i-1].Left!=nil{
// 				return false
// 			}
// 		}

		
// 		que1=tempQue1
// 		que2=tempQue2
		
// 	}

// 	return true

// }

//在内存使用上没什么提升
func isSymmetric(root *TreeNode) bool {
	que:=make([]*TreeNode,0)
	
	//对根节点进行判断
	if root!=nil {
		que=append(que,root)
	}

	//层次便利，对没一层都进行判断
	for len(que)!=0{
		tempQue:=make([]*TreeNode,0)

		for i:=0;i<len(que);i++{

			if que[i].Left !=nil && que[len(que)-i-1].Right!=nil{
				if que[i].Left.Val !=que[len(que)-i-1].Right.Val{
					//fmt.Println("case 1")
					return false
				} 
				tempQue=append(tempQue,que[i].Left)
			}else if que[i].Left !=nil && que[len(que)-i-1].Right==nil{
				//fmt.Println("case 2")
				return false
			}else if que[i].Left ==nil && que[len(que)-i-1].Right!=nil{
				fmt.Println("case 3")
				return false
			}



			if que[i].Right !=nil && que[len(que)-i-1].Left!=nil{
				if que[i].Right.Val !=que[len(que)-i-1].Left.Val{
					//fmt.Println("case 4")
					return false
				} 
				tempQue=append(tempQue,que[i].Right)
			}else if que[i].Right !=nil && que[len(que)-i-1].Left==nil{
				//fmt.Println("case 5")
				return false
			}else if que[i].Right ==nil && que[len(que)-i-1].Left!=nil{
				//fmt.Println("case 6")
				return false
			}
		}

		que=tempQue
		fmt.Println(len(que))
		
	}

	return true

}

// @lc code=end

