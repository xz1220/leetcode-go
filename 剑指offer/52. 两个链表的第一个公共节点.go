/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    head := headA
    for headA.Next != nil {
        headA = headA.Next
    }
    headA.Next = headB

    // 快慢指针
    slow, fast := head, head
    for fast != nil && fast.Next != nil && slow != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            temp := head
            for temp != slow {
                temp = temp.Next
                slow = slow.Next
            }
            headA.Next = nil
            return temp
        }
    }

    headA.Next = nil
    return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    
    a, b := headA, headB
    for a != b {
        if a == nil {
            a = headB
        }else {
            a = a.Next
        }

        if b == nil {
            b = headA
        }else {
            b = b.Next
        }
    }

    return a
}