/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 *
 * https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (33.40%)
 * Likes:    415
 * Dislikes: 0
 * Total Accepted:    55.9K
 * Total Submissions: 165.6K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 * 
 * 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * 输出：[0,9]
 * 解释：
 * 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 * 输出的顺序不重要, [9,0] 也是有效答案。
 * 
 * 
 * 示例 2：
 * 
 * 输入：
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * 输出：[]
 * 
 * 
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	// 算法： 双指针法
	// 首先根据words初始化一个hash表，
	// 循环遍历起点，按照单个word的长度：
	//     初始化一个计数器和一个计数Hash表
	//     定义双执政l、r，l为起点，r来读取单词，循环从起点逐个读取单词，
	//     判断在words 的hash 表中有无对应的单词; 若有：
	//          如果计数hash表中的单词总数已经超过words的hash表
    //          循环移动l直到满足 计数单词总数小于words中的Hash表，更新计数器和计数Hash表的值
	//          对应计数器和hash表+1
    //     若无：
    //          循环移动l到当前r
    //          l移动到下一个单词处
    //     如果 当前的计数器的大小等于words单词的长度：
    //          在结果中加入l
    // 返回结果
    // 边界条件：1. words为空，返回空 2. words 单词总长度大于s的长度，返回空

    sLen := len(s)
    numsWords := len(words)
    lenWord := len(words[0])
    result := []int{}
    wordsMap := make(map[string]int)
    distance := (numsWords-1)*lenWord

    if numsWords == 0 || numsWords*lenWord > sLen {
        return []int{}
    }

    for _,word := range words {
        wordsMap[word]++ 
    }

    for i := 0 ; i< lenWord ; i++ {
        counterMap := make(map[string]int)

        for l,r := i, i ; r<= sLen - lenWord ; r+=lenWord {
            word := s[r:r+lenWord]
            if num, find := wordsMap[word] ; find {
                for counterMap[word] >= num {
                    counterMap[s[l:l+lenWord]]--
                    l+=lenWord
                }

                counterMap[word]++
            }else {
                for l<r {
                    counterMap[s[l:l+lenWord]]--
                    l+=lenWord
                }
                l+=lenWord
            }

            if (r - l) == distance {
                result = append(result, l)
            }
        }
    }
    
    return result

}

// @lc code=end

