/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 *
 * https://leetcode-cn.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (76.57%)
 * Likes:    809
 * Dislikes: 0
 * Total Accepted:    133K
 * Total Submissions: 173.1K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 * 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：S = "ababcbacadefegdehijhklij"
 * 输出：[9,7,8]
 * 解释：
 * 划分结果为 "ababcbaca", "defegde", "hijhklij"。
 * 每个字母最多出现在一个片段中。
 * 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
 *
 *
 *
 *
 * 提示：
 *
 *
 * S的长度在[1, 500]之间。
 * S只包含小写字母 'a' 到 'z' 。
 *
 *
 */

// @lc code=start
func soluction1(s string) []int {
	leftDict := map[byte]int{}
	for index := range s {
		leftDict[s[index]]++
	}

	stack := make([]byte, 0)
	instack := [26]bool{}
	ans := []int{}

	for index := range s {
		digit := s[index]

		// try add it to the stack
		stack = append(stack, digit)
		instack[digit-'a'] = true
		leftDict[digit] -= 1

		// 判断是否满足条件
		gogo := false
		for num, exist := range instack {
			if exist && leftDict[byte(num)+'a'] != 0 {
				gogo = true
				break
			}
		}

		// 右边存在重复
		if gogo {
			continue
		}

		// 不存在重复
		ans = append(ans, len(stack))
		stack = stack[:0]
		instack = [26]bool{}
	}

	return ans
}

func soluction2(s string) []int {
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i
	}

	ans := []int{}
	start, end := 0, 0

	for i, c := range s {
		end = max(end, last[c-'a'])

		if i == end {
			ans = append(ans, end-start+1)
			start = i + 1
		}
	}

	return ans
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}

func partitionLabels(s string) []int {
	return soluction1(s)
}

// @lc code=end

