/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (29.23%)
 * Likes:    2932
 * Dislikes: 0
 * Total Accepted:    412.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 *
 *
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

// func threeSum(nums []int) [][]int {
//     // algorithm:
//     // sort the nums
//     // init the result
//     // for first = 0,..., n-1 then
//     //      third = n-1
//     //      sum = -nums[first]
//     //      for second = first + 1; second< third ; second ++ then
//     //          if nums[second] == nums[second+1] then
//     //              continue
//     //           for ; nums[third] + nums[second]  > sum ; third -- ;
//     //           if nums[third] + nums[second] == sum then
//     //              result.appen([first,second,third])
//     // return result
//     //
//     sort.Ints(nums)
//     n:=len(nums)
//     result:=make([][]int,0)
//     for first:=0; first <= n-1; first++ {
//         // 第一个数也不能重复
//         if first>0 && nums[first] == nums[first-1] {
//             continue
//         }
//         third := n-1
//         sum := -1*nums[first]
//         for second:= first+1 ; second < third ; second++ {
//             if second> first+1 && nums[second] == nums[second-1] {
//                 continue
//             }

//             for second< third-1 && nums[third] + nums[second] > sum {
//                 third--
//             }

//             if nums[third]+ nums[second] == sum {
//                 result = append(result,[]int{nums[first],nums[second],nums[third]})
//             }
//         }
//     }

//     return result

// }
// @lc code=end

