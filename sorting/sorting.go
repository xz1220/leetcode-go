package sorting

import (
	"fmt"
	"sort"
)

/*
 *
 * [56] 合并区间

 * algorithms
 * Medium (40.33%)
 * Likes:    288
 * Dislikes: 0
 * Total Accepted:    53.9K
 * Total Submissions: 133.6K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

func Merge(intervals [][]int) [][]int {
	l := len(intervals) * 2
	//const length = l
	newlist := [][]int{}
	i := 0
	j := 0
	for k := 0; k < l; k++ {
		if i < l/2 && intervals[i][0] <= intervals[j][1] {
			newlist = append(newlist, []int{intervals[i][0], 0})
			//fmt.Println("i:", i)
			i++
		} else {
			newlist = append(newlist, []int{intervals[j][1], 1})
			//fmt.Println("j:", j)
			j++
		}
	}
	index := [][]int{}
	flag := true
	fmt.Println(newlist)
	for k := 0; k < l; k++ {
		if newlist[k][1] == 0 && flag {
			i = newlist[k][0]
			flag = false
			//fmt.Println("i:", i)
		}
		if newlist[k][1] == 1 && !flag {
			if ((k+1) != l && newlist[k+1][1] == 0) || (k+1) == l {
				j = newlist[k][0]
				flag = true
				//fmt.Println("j:", j)
				index = append(index, []int{i, j})
			}

		}
	}
	//fmt.Println(index)
	return index

}

/*
上面的版本缺陷很大，首先是因为它基于了一个并不存在的假设，就是每个pair的第一个数，第二个数是按顺序排列，
可以组成两个顺序的数组。但这个假设显然是错误的。
导致上面这个解法无法处理两种情况：
1. [[1,4],[0,4]]
2. [[1,4],[2,3]]

*/

/*
下面这个方法的逻辑是，先对按照每个pair的第一个数的大小进行排序，
从第一个pair开始，将它的第二个数与后面的pair的第一个和第二个数进行比较，那么会有如下三种情况。
1. 大于某个pair的第一个数和第二个数 那么说明这个pair包含在其中
2. 大于等于某个pair的第一个数，小于某个pair的第二个数，说明这两个pair的取键可以进行合并，此时更新第一个pair的值
3. 小于某个pair的第一个数，此时说明往后再也找不到可以合并的区间，于是将存储合并后的pair

注意以上都需要判断 临界条件

时间复杂度：
	排序：O(?) ==> O(nlog(n)) or O(n^2)
	搜索比较pair：O(n)

** 还有优化的空间
*/

func Merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	index := [][]int{}
	flag := 0
	for k := 0; k < len(intervals); k++ {
		if intervals[flag][1] >= intervals[k][0] && intervals[flag][1] >= intervals[k][1] {
			fmt.Println("case 1")
			if k == len(intervals)-1 {
				index = append(index, []int{intervals[flag][0], intervals[flag][1]})

			}

		} else if intervals[flag][1] >= intervals[k][0] && intervals[flag][1] < intervals[k][1] {
			fmt.Println("case 2")
			intervals[flag][1] = intervals[k][1]
			if k == len(intervals)-1 {
				index = append(index, []int{intervals[flag][0], intervals[flag][1]})
			}
		} else if intervals[flag][1] < intervals[k][0] {
			fmt.Println("case 3")
			index = append(index, []int{intervals[flag][0], intervals[flag][1]})
			flag = k
			if k == len(intervals)-1 {
				index = append(index, []int{intervals[flag][0], intervals[flag][1]})
			}
		}
	}
	fmt.Println(index)

	return index
}

// BubbleSort is a //
func BubbleSort(list []int) []int {
	length := len(list)
	for i := 0; i < length; i++ {
		flag := true
		for j := 0; j < length-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	//fmt.Println(list)
	return list
}

//MergeSort is
func MergeSort(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}

	midddle := length / 2
	//递归的写法
	return merge(MergeSort(list[0:midddle]), MergeSort(list[midddle:]))

}

func merge(left []int, right []int) []int {
	//result := make([]int, len(left)+len(right))
	result := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
