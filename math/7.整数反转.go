/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
// func reverse(x int) int {
// 	tmp := 0
// 	for x != 0 {
// 		tmp = tmp*10 + x%10
// 		x = x / 10
// 	}
// 	if tmp > 1<<31-1 || tmp < -(1<<31) {
// 		return 0
// 	}
// 	return tmp
// }

func reverse(x int) int {
	var res int
	for x != 0 {
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		// if int32(res*10)/10 != int32(res) {
		//     return 0
		// }
		res = res*10 + x%10
		x = x / 10
	}
	return res
}

// @lc code=end

