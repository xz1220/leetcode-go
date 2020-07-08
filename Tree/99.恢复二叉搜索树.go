/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
 *
 * https://leetcode-cn.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Hard (55.67%)
 * Likes:    170
 * Dislikes: 0
 * Total Accepted:    13.9K
 * Total Submissions: 24.8K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * 二叉搜索树中的两个节点被错误地交换。
 * 
 * 请在不改变其结构的情况下，恢复这棵树。
 * 
 * 示例 1:
 * 
 * 输入: [1,3,null,null,2]
 * 
 * 1
 * /
 * 3
 * \
 * 2
 * 
 * 输出: [3,1,null,null,2]
 * 
 * 3
 * /
 * 1
 * \
 * 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,1,4,null,null,2]
 * 
 * ⁠ 3
 * ⁠/ \
 * 1   4
 * /
 * 2
 * 
 * 输出: [2,1,4,null,null,3]
 * 
 * ⁠ 2
 * ⁠/ \
 * 1   4
 * /
 * ⁠ 3
 * 
 * 进阶:
 * 
 * 
 * 使用 O(n) 空间复杂度的解法很容易实现。
 * 你能想出一个只使用常数空间的解决方案吗？
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


func RecursiveRecoverTree(root *TreeNode)  {

	//两个指针是必要的，因为根据题目而言，只存在一对错误的节点，也就是在中序遍历的过程中
	//只存在两个pair不符合从小到大的规律
	//我之前的那种方法类似于一次冒泡，只能保证最大的节点的位置是正确的

	var first,second,pre *TreeNode
	var dfs func(*TreeNode)

	//中序遍历
	dfs=func(root *TreeNode){
		if root.Left!=nil{
			dfs(root.Left)
		}

		//比较当前节点
		if pre!=nil && pre.Val>root.Val{
			if first==nil{
				first=pre
			}

			if first!=nil{
				second=root
			}
		}


		//标记当前的节点
		pre=root

		if root.Right!=nil{
			dfs(root.Right)
		}

	}

	dfs(root)

	first.Val,second.Val=second.Val,first.Val
}

/*
非递归
*/

func NonRecursiveRecoverTree(root *TreeNode)  {

	stack :=make([]*TreeNode,0)
	var (
		pre *TreeNode
		x *TreeNode
		y *TreeNode
	)

	for root!=nil || len(stack)!=0{
		if root!=nil{
			stack = append(stack, root)
			root = root.Left
		}else{
			root =stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			if pre!=nil && pre.Val>=root.Val{

                y = root
                if x == nil {
                    x = pre
                }
			}

            pre = root
            root = root.Right

		}
	}

    x.Val,y.Val=y.Val,x.Val

}

// Morris 算法
func recoverTree(root *TreeNode) {
	var (
		predecessor *TreeNode
		x *TreeNode
		y *TreeNode
		pred *TreeNode
	)


	for root!=nil{
		if root.Left == nil {
			if pred!=nil && root.Val < pred.Val{
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root

			root = root.Right
		}else{
			predecessor = GetPreDecessor(root)

			if predecessor.Right==nil{
				predecessor.Right = root
				root = root.Left
			}else{
				if pred!=nil && root.Val < pred.Val{
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root

				predecessor.Right = nil
				root = root.Right
			}
		}
	}

	x.Val, y.Val=y.Val,x.Val
}

func GetPreDecessor(node *TreeNode) *TreeNode{
	var pre *TreeNode=node
	if node.Left != nil {
		pre = pre.Left
		for pre.Right!=nil && pre.Right != node{
			pre = pre.Right
		}
	}

	return pre
}



// @lc code=end

