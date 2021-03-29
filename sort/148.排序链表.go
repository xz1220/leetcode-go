/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (67.73%)
 * Likes:    1065
 * Dislikes: 0
 * Total Accepted:    151.9K
 * Total Submissions: 225.7K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
 * 
 * 进阶：
 * 
 * 
 * 你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [4,2,1,3]
 * 输出：[1,2,3,4]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [-1,5,3,4,0]
 * 输出：[-1,0,3,4,5]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：head = []
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目在范围 [0, 5 * 10^4] 内
 * -10^5 
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
// var dummyNode *ListNode
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    temp := slow.Next
    slow.Next = nil
    slow = temp
    left := sortList(head)
    right := sortList(slow)
    left = mergeSortedList(left, right)
    return left
}

func mergeSortedList(l, r *ListNode) *ListNode {
    if l == nil {
        return r
    }
    if r == nil {
        return l
    }
    dummyNode := &ListNode{0, l}
    l = dummyNode
    for l.Next != nil && r != nil {
        if r.Val >= l.Next.Val {
            l = l.Next
        }else {
            temp := r.Next
            r.Next = l.Next
            l.Next = r
            r = temp
            l = l.Next
        }
    }

    if r != nil {
        l.Next = r
    }
    return dummyNode.Next
}
// @lc code=end

