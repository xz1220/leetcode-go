/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 *
 * https://leetcode-cn.com/problems/scramble-string/description/
 *
 * algorithms
 * Hard (48.65%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    18K
 * Total Submissions: 36.9K
 * Testcase Example:  '"great"\n"rgeat"'
 *
 * 给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。
 * 
 * 下图是字符串 s1 = "great" 的一种可能的表示形式。
 * 
 * ⁠   great
 * ⁠  /    \
 * ⁠ gr    eat
 * ⁠/ \    /  \
 * g   r  e   at
 * ⁠          / \
 * ⁠         a   t
 * 
 * 
 * 在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。
 * 
 * 例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。
 * 
 * ⁠   rgeat
 * ⁠  /    \
 * ⁠ rg    eat
 * ⁠/ \    /  \
 * r   g  e   at
 * ⁠          / \
 * ⁠         a   t
 * 
 * 
 * 我们将 "rgeat” 称作 "great" 的一个扰乱字符串。
 * 
 * 同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。
 * 
 * ⁠   rgtae
 * ⁠  /    \
 * ⁠ rg    tae
 * ⁠/ \    /  \
 * r   g  ta  e
 * ⁠      / \
 * ⁠     t   a
 * 
 * 
 * 我们将 "rgtae” 称作 "great" 的一个扰乱字符串。
 * 
 * 给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。
 * 
 * 示例 1:
 * 
 * 输入: s1 = "great", s2 = "rgeat"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: s1 = "abcde", s2 = "caebd"
 * 输出: false
 * 
 */

// @lc code=start

/**
 * 用例超出时间限制
 */

// func isScramble(s1 string, s2 string) bool {
//     s1Length, s2Length := len(s1), len(s2)

//     if s1Length != s2Length {
//         return false
//     }

//     if s1 == s2 {
//         return true
//     }

//     // 剪枝
//     // s3 := []rune(s1)
//     // sort.Slice(s3, func(i, j int) bool {
//     //     return s3[i] < s3[j]
//     // })
//     // s4 := []rune(s2)
//     // sort.Slice(s4, func(i, j int) bool {
//     //     return s4[i] < s4[j]
//     // })

//     // for index := range s3 {
//     //     if s3[index] != s4[index] {
//     //         return false
//     //     }
//     // }

//     for index:=1 ; index < s1Length ; index++ {
//         x1 := s1[0:index]
//         y1 := s1[index:]
//         x2 := s2[0:index]
//         y2 := s2[index:]
//         x3 := s2[0:s1Length - index]
//         y3 := s2[s1Length-index:]

//         if isScramble(x1,x2) && isScramble(y1,y2) {
//             return true
//         }
//         if isScramble(y1,x3) && isScramble(x1,y3) {
//             return true
//         }
//     }

//     return false
// }

var memo map[string]bool
func isScramble(s1 string, s2 string) bool {
	memo = make(map[string]bool)
	return dfs(s1, s2)
}
func dfs(s1 string, s2 string) bool {
	key := fmt.Sprintf("%s_%s", s1, s2)
	if value, ok := memo[key]; ok {
		return value
	}
	// 空字符 或者 长度不一，直接返回false
	if s1 == "" || s2 == "" || len(s1) != len(s2) {
		memo[key] = false
		return false
	}
	if s1 == s2 {
		memo[key] = true
		return true
	}
	letters := make([]int, 26)
	for k, v := range s1 {
		letters[v-'a']++
		letters[s2[k]-'a']--
	}
	// 字母统计不一致，直接返回false
	for _, v := range letters {
		if v != 0 {
			memo[key] = false
			return false
		}
	}
	l := len(s1)
	for i := 1; i < l; i++ {
		// 分割点： s1左==s2左 && s1右==s2右
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			memo[key] = true
			return true
		}
		// 分割点： s1左==s2右 && s1右==s2左
		if isScramble(s1[:i], s2[l-i:]) && isScramble(s1[i:], s2[0:l-i]) {
			memo[key] = true
			return true
		}
	}
	memo[key] = false
	return false
}
// @lc code=end

