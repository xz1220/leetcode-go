/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (53.94%)
 * Likes:    1231
 * Dislikes: 0
 * Total Accepted:    232.6K
 * Total Submissions: 425.7K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 * 
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 * 
 * 
 * 示例 2：
 * 
 * 输入：lists = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 输入：lists = [[]]
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
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
 func mergeKLists(lists []*ListNode) *ListNode {

    minHeap := new(MinHeap)
    for index := range lists {
        if lists[index] != nil {
            heap.Push(minHeap, lists[index])
        }
    }

    var dummyHead *ListNode = &ListNode{}
    var current *ListNode = dummyHead
    for minHeap.Len() > 0 {
        node := heap.Pop(minHeap).(*ListNode)
        if node.Next != nil {
            heap.Push(minHeap, node.Next)
        }

        current.Next = node
        current = node
    }

    return dummyHead.Next
}

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }else if len(lists) == 1 {
        return lists[0]
    }else if len(lists) == 2 {
        return merge(lists[0], lists[1])
    }
    
    length := len(lists)
    mergeList1 := mergeKLists(lists[:length/2])
    mergeList2 := mergeKLists(lists[length/2:])
    return merge(mergeList1, mergeList2)
}

func merge(list1, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }else if list2 == nil {
        return list1
    }

    dummyNode := &ListNode{0, list1}
    pre := dummyNode
    node1, node2 := list1, list2
    for node1 != nil && node2 != nil {
        if node2.Val >= node1.Val {
            node1 = node1.Next
            pre = pre.Next
        }else {
            temp := node2
            node2 = node2.Next
            temp.Next = node1
            pre.Next = temp
            pre = pre.Next  //  插入之后，pre 后移
        }
    }

    if node1 == nil {
        pre.Next = node2
    }
    return dummyNode.Next
}
// @lc code=end

