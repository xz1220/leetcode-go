/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (66.80%)
 * Likes:    812
 * Dislikes: 0
 * Total Accepted:    75.9K
 * Total Submissions: 112.3K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：3
 * 输出：
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释：
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= n <= 8
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
 func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return nil
    }

    return helper(1, n)
}

func helper(start , end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }

    ans := []*TreeNode{}

    for i:=start ; i<= end ; i++ {
        leftTrees := helper(start, i -1)
        rightTrees := helper(i+1, end)

        for _, leftTree := range leftTrees {
            for _, rightTree := range rightTrees {
                currentTree := &TreeNode{i, nil, nil}
                currentTree.Left = leftTree
                currentTree.Right = rightTree
                ans = append(ans, currentTree)
            }
        }
    }

    return ans
}
// @lc code=end

