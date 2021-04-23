/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (49.12%)
 * Likes:    756
 * Dislikes: 0
 * Total Accepted:    266.6K
 * Total Submissions: 540.4K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 * 
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自
 * nums2 的元素。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
 * 输出：[1,2,2,3,5,6]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums1 = [1], m = 1, nums2 = [], n = 0
 * 输出：[1]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * nums1.length == m + n
 * nums2.length == n
 * 0 
 * 1 
 * -10^9 
 * 
 * 
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    for p:= m+n ; m>0 && n>0  ; p-- {
        if nums1[m-1] > nums2[n-1] {
            nums1[p-1] = nums1[m-1]
            m = m -1
        }else {
            nums1[p-1] = nums2[n-1]
            n = n -1
        }
    }

    if m == 0 {
        for n>0 {
            nums1[n -1] = nums2[n -1]
            n -- 
        }
    }

}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    p1, p2, pMax := m -1, n -1, len(nums1) -1
    for p1 >=0 && p2 >=0 {
        if nums1[p1] > nums2[p2] {
            nums1[pMax] = nums1[p1]
            p1 --
            pMax --
        }else {
            nums1[pMax] = nums2[p2]
            p2 --
            pMax --
        }
    }

    if p1 <= 0 {                            // < :---> [0] 0 [1] 1
        for pMax >=0 && p2 >=0 {
            nums1[pMax] = nums2[p2]
            p2 --
            pMax --
        }
    }
}
// @lc code=end

