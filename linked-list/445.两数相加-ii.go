/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (58.23%)
 * Likes:    358
 * Dislikes: 0
 * Total Accepted:    66.1K
 * Total Submissions: 113.2K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
 * 
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 * 
 * 
 * 
 * 进阶：
 * 
 * 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 8 -> 0 -> 7
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

 // 这样做会溢出
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    number1, number2 := 0, 0
    for l1 != nil {
        number1 = number1 *10 + l1.Val
        l1 = l1.Next
    }
    for l2 != nil {
        number2 = number2 *10 + l2.Val
        l2 = l2.Next
    }

    sum := number1 + number2
    if sum == 0 {
        return &ListNode{
            Val: 0,
            Next: nil,
        }
    }

    var current *ListNode = nil
    for sum > 0 {
        temp := sum % 10
        tempNode := &ListNode{temp,current}

        sum = (sum - temp)/10
        current = tempNode    
    }
    return current
}
// @lc code=end

