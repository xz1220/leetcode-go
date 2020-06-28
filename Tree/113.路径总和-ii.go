/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (58.96%)
 * Likes:    190
 * Dislikes: 0
 * Total Accepted:    37.9K
 * Total Submissions: 64.3K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 * 
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 * 
 * 
 * 返回:
 * 
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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
//  func pathSum(root *TreeNode, sum int) [][]int {
// 	stack:=make([]*TreeNode,0)
// 	var sumNow int=0
// 	var parent int=0
// 	var flag=true
// 	var pre *TreeNode=root
// 	var list [][]int

// 	for pre!=nil || len(stack)!=0{
// 		if pre!=nil{
// 			sumNow+=pre.Val
// 			stack=append(stack,pre)
// 			pre=pre.Left
// 		}else{
			
// 			node:=stack[len(stack)-1]
// 			if node.Right==nil{
// 				if len(stack)==1{
// 					if sumNow==sum{
// 						list=storeList(stack,list)
// 					}
// 					break
// 				}
// 				if stack[len(stack)-2].Left==node{
// 					if sumNow==sum{
// 						list=storeList(stack,list)
// 					}
// 					if flag{
// 						parent=len(stack)-2
// 						flag=false
// 					}	
// 					sumNow-=stack[len(stack)-1].Val
// 					stack=stack[0:len(stack)-1]
// 					stack[len(stack)-2].Left=nil
// 					pre=node.Right
// 				}
// 				if stack[len(stack)-2].Right==node{
// 					if sumNow==sum{
// 						list=storeList(stack,list)
// 					}
// 					for _,n:=range stack[parent:]{
// 						sumNow-=n.Val
// 					}
//                     flag=true
// 					stack=stack[:parent]
// 					node=stack[parent-1]
// 					pre=node.Right
// 				}
// 			}else{
// 				// sumNow-=stack[len(stack)-1].Val
// 				// stack=stack[:len(stack)-1]
// 				pre=node.Right
// 			}
// 		}
// 	}
//     return list
// }

// func storeList(listNode []*TreeNode,l [][]int)[][]int{
// 	list:=[]int{}
// 	for _,node:=range listNode{
// 		list=append(list,node.Val)
// 	}
// 	return append(l,list)
// }

//非递归方法失败了
// func pathSum(root *TreeNode, sum int) [][]int {
// 	var list [][]int
// 	var storeList []int=[]int{}
// 	var dfs func(*TreeNode,int,[]int)

// 	dfs=func(root *TreeNode,val int,l []int){
// 		if root==nil{
// 			return
// 		}
// 		temp:=make([]int,len(l)+1)
// 		copy(temp,l)
// 		temp[len(l)]=root.Val

// 		if root.Left==nil && root.Right==nil{
// 			if val+root.Val==sum{
// 				list=append(list,temp)
// 			}
// 		}else if root.Left==nil && root.Right!=nil{
// 			dfs(root.Right,val+root.Val,temp)
// 		}else if root.Right==nil && root.Left!=nil{
// 			dfs(root.Left,val+root.Val,temp)
// 		}else{
// 			// temp2:=make([]int,len(temp))
// 			// copy(temp2,temp)

// 			dfs(root.Left,val+root.Val,temp)
// 			dfs(root.Right,val+root.Val,temp)

// 		}
// 	}

// 	dfs(root,0,storeList)
// 	return list
// }

//我的做法会在没一层递归都进行深copy 这是非常耗费时间也浪费内存，其实只要在找到满足条件的情况下进行深copy就可以

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level, sum int) {
		if node == nil {
			return
		}

		// 根据 level 来更新 path
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val

		// 到达 leaf
		// 并且，此条路径符合要求
		if node.Left == nil && node.Right == nil && sum == 0 {
			temp := make([]int, level+1)
			// copy 不会复制 path[len(temp):]，如果 len(path) > len(temp) 的话
			copy(temp, path)
			res = append(res, temp)
		}

		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}

	dfs(root, 0, sum)

	return res
}

// @lc code=end

