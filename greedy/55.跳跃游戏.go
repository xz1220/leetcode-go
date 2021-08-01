/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (42.94%)
 * Likes:    1281
 * Dislikes: 0
 * Total Accepted:    280.6K
 * Total Submissions: 653.1K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 判断你是否能够到达最后一个下标。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,3,1,1,4]
 * 输出：true
 * 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,2,1,0,4]
 * 输出：false
 * 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 0 
 * 
 * 
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	dic := make([]int, len(nums))

	var current int
	for index, val := range nums {
		maxIndex := index + val
		if val == 0 || maxIndex < current {
			continue
		}

		distance := dic[index]
		for j := current ; j <= maxIndex && j <= len(nums) -1; j++ {
			dic[j] = distance +1
		}
		fmt.Println(current, maxIndex, dic)

		current = maxIndex + 1
		if current >= len(nums) {
			break
		}
	}

	return dic[len(dic) -1] != 0
}

// 只需要判断是否可达即可
func canJump(nums []int) bool {
    if len(nums) == 1 {
		return true
	}
	rightmax := 0
    
    for index, val := range nums {
        if index > rightmax {
            return false
        }

        if rightmax < index + val {
            rightmax = index + val
        }
    }

	return true
}

// @lc code=end

