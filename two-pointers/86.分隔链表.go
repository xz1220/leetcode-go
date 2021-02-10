/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (62.23%)
 * Likes:    360
 * Dislikes: 0
 * Total Accepted:    87.6K
 * Total Submissions: 140.4K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
 * 
 * 你应当 保留 两个分区中每个节点的初始相对位置。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [1,4,3,2,5,2], x = 3
 * 输出：[1,2,2,4,3,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [2,1], x = 2
 * 输出：[1,2]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目在范围 [0, 200] 内
 * -100 
 * -200 
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
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func partition(head *ListNode, x int) *ListNode {
	// 算法：
	// 定义快慢指针：lowPoint highPoint 指向哑节点
	// 如果下一个节点的值小于x，两者都指向下一个节点
	// 如果下一个节点的值大于或等于x，highPoint 指向下一个节点，直到下一个节点小于x
	// 指针交换，并都向下一个节点移动 

	dump := &ListNode{x-1,head}
	lowPoint := dump
	for lowPoint.Next!= nil && lowPoint.Next.Val < x {
		lowPoint = lowPoint.Next
	}
	highPoint := lowPoint

	for highPoint.Next != nil {
		for highPoint.Next!=nil && highPoint.Next.Val >= x {
			highPoint = highPoint.Next
		}

		if highPoint.Next == nil {
			continue
		}
		temp := highPoint.Next
		highPoint.Next = temp.Next
		temp.Next = lowPoint.Next
		lowPoint.Next = temp

		lowPoint = lowPoint.Next
	   //  highPoint = highPoint.Next
	}

	return dump.Next
}
// @lc code=end

