/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
 *
 * https://leetcode-cn.com/problems/binary-search-tree-iterator/description/
 *
 * algorithms
 * Medium (71.87%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    15.8K
 * Total Submissions: 21.9K
 * Testcase Example:  '["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]\n[[[7,3,15,null,null,9,20]],[null],[null],[null],[null],[null],[null],[null],[null],[null]]'
 *
 * 实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
 * 
 * 调用 next() 将返回二叉搜索树中的下一个最小的数。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * BSTIterator iterator = new BSTIterator(root);
 * iterator.next();    // 返回 3
 * iterator.next();    // 返回 7
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 9
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 15
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 20
 * iterator.hasNext(); // 返回 false
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
 * 你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
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
type BSTIterator struct {
	now int
	nums []int
}


func Constructor(root *TreeNode) BSTIterator {

	var nums []int
	var dfs func(*TreeNode)

	dfs=func(root *TreeNode){
		if root==nil{
			return
		}
		dfs(root.Left)
		nums=append(nums,root.Val)
		dfs(root.Right)
	}

	dfs(root)

	return BSTIterator{nums:nums}
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.now<len(this.nums){
		defer func(){
			this.now+=1
		}()
		return this.nums[this.now]
	}
	panic("到头了！")
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.now<len(this.nums){
		return true
	}
	return false
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

