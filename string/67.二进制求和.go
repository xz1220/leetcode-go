/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (54.45%)
 * Likes:    648
 * Dislikes: 0
 * Total Accepted:    181.4K
 * Total Submissions: 333.2K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串，返回它们的和（用二进制表示）。
 * 
 * 输入为 非空 字符串且只包含数字 1 和 0。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: a = "11", b = "1"
 * 输出: "100"
 * 
 * 示例 2:
 * 
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 每个字符串仅由字符 '0' 或 '1' 组成。
 * 1 <= a.length, b.length <= 10^4
 * 字符串如果不是 "0" ，就都不含前导零。
 * 
 * 
 */

// @lc code=start
func addBinary(as string, bs string) string {
	a := []byte(as)
	b := []byte(bs)
	ans := make([]byte, 0)

	var flag byte = 0
	var i, j int
	for i, j = len(a) - 1 , len(b) - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		temp := (a[i] - 48) + (b[j] - 48) + flag
		flag = (temp - temp %2) /2
		temp = temp % 2
		
		ans = append(ans, temp)
	}

	fmt.Println(ans, flag)

	if i < 0 && j >= 0 {
		for ; j >=0 ; j -- {
			temp := (b[j] - 48) + flag

			flag = (temp - temp %2) /2
			temp = temp % 2
			
			ans = append(ans, temp)
		}
	}

	if j < 0 && i >= 0{
		for ; i >= 0 ; i -- {
			temp := (a[i] - 48) + flag
			
			flag = (temp - temp %2) /2
			temp = temp % 2
			
			ans = append(ans, temp)
		}
	}

	if  flag != 0 {
		ans = append(ans, flag)
	}

	fmt.Println(ans)
	reverse(ans)
	fmt.Println(ans)

	return string(ans)
	
}

func reverse(ans []byte) []byte {
	length := len(ans)

	for i := 0 ; i < length / 2; i++ {
		ans[i], ans[length - 1 -i] = ans[length -1 - i] + 48, ans[i] + 48
	}

	if length %2 != 0 {
		ans[length / 2] += 48
	}

	return ans
}
// @lc code=end

