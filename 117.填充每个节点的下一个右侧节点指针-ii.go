/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 *
 * https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (46.45%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    16.4K
 * Total Submissions: 35.4K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * 给定一个二叉树
 * 
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 * 
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
 * 
 * 初始状态下，所有 next 指针都被设置为 NULL。
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 你只能使用常量级额外空间。
 * 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 * 
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * 输入：root = [1,2,3,4,5,null,7]
 * 输出：[1,#,2,3,#,4,5,7,#]
 * 解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的节点数小于 6000
 * -100 <= node.val <= 100
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 //很显然这一题我没有找到进阶的办法，还是用的老方法
 //这种方法无法保证一定是在常量级别的额外内存空间，因为需要队列来保存某一层的节点，
 //可能会有新的非递归的方法
 //可能会有新的递归方法
func connect(root *Node) *Node {

	que:=make([]*Node,0)
	if root!=nil{
		que=append(que,root)
		root.Next=nil
	}

	for len(que)!=0{
		temp:=make([]*Node,0)
		for _,node:=range que{
			if node.Left!=nil{
				temp=append(temp,node.Left)
			}
			if node.Right!=nil{
				temp=append(temp,node.Right)
			}
		}
		for i:=0;i<len(temp)-1;i++{
			if i==len(temp)-1{
				temp[i].Next=nil
				break
			}
			temp[i].Next=temp[i+1]
		}
		que=temp
	}

	return root
	
}
// @lc code=end

