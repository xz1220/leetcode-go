/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 *
 * https://leetcode-cn.com/problems/interleaving-string/description/
 *
 * algorithms
 * Hard (45.69%)
 * Likes:    408
 * Dislikes: 0
 * Total Accepted:    42.3K
 * Total Submissions: 92.3K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
 * 
 * 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
 * 
 * 
 * s = s1 + s2 + ... + sn
 * t = t1 + t2 + ... + tm
 * |n - m| 
 * 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 +
 * ...
 * 
 * 
 * 提示：a + b 意味着字符串 a 和 b 连接。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * 输出：false
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s1 = "", s2 = "", s3 = ""
 * 输出：true
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * 0 
 * s1、s2、和 s3 都由小写英文字母组成
 * 
 * 
 */

// @lc code=start
/** wrong answer **/
func isInterleave(s1 string, s2 string, s3 string) bool {
    if s1 == "" {
        if s3 == s2 {
            return true
        }
        return false
    }
    if s2 == "" {
        if s3 == s1 {
            return true
        }
        return false
    }
    s1Length, s2Length, s3Length := len(s1), len(s2), len(s3)

    index1, index2, index3 := 0, 0, 0
    flag1, flag2 := true, false
    flag3, flag4 := false, false
    for index3 < s3Length {
        if s3[index3]!=s2[index2] && s3[index3] != s1[index1] {
            // fmt.Println(s1[index1], s2[index2], s3[index3])
            fmt.Println(index1, index2, index3)
            return false
        }

        if flag1 && index1<s1Length && s3[index3] == s1[index1] {
            if index1 < s1Length -1 {
                index1 ++ 
            }
            index3 ++ 
            flag3 = true
            continue
        }else if flag1 == false {
            // do nothing
        }else if flag1 && index1<s1Length && s3[index3] != s1[index1] {
            flag1 = false
            flag2 = true
            continue
        }

        if flag2 && index2<s2Length && s3[index3] == s2[index2] {
            if index2 < s2Length-1 {
                index2 ++
            }
            index3 ++ 
            flag4 = true
            continue
        }else if flag2 == false {
            // do nothing
        }else if flag2 && index2<s2Length && s3[index3] != s2[index2] {
            flag1 = true
            flag2 = false
            continue
        }
        
    }

    return true && flag3 && flag4
}
// @lc code=end

