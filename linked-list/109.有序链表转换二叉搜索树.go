/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (76.14%)
 * Likes:    491
 * Dislikes: 0
 * Total Accepted:    75.7K
 * Total Submissions: 99.3K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
 * 
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 * 
 * 示例:
 * 
 * 给定的有序链表： [-10, -3, 0, 5, 9],
 * 
 * 一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
 * 
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
 * 
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var globalHead *ListNode

 func sortedListToBST(head *ListNode) *TreeNode {
	 globalHead = head
	 return build(0, getLen(head) -1)
 }
 
 func getLen(head *ListNode) int {
	 length := 0
	 for head != nil {
		 length ++
		 head = head.Next
	 }
	 return length
 }
 
 func build(left , right int) *TreeNode {
	 if left > right {
		 return nil
	 }
	 root := &TreeNode{}
	 middle := (right + left)/2 
	 root.Left = build(left, middle-1)
	 root.Val = globalHead.Val
	 globalHead = globalHead.Next
	 root.Right = build(middle +1, right)
	 return root
 }
// @lc code=end

