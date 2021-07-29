func findContinuousSequence(target int) [][]int {
    var ans [][]int

    if target == 1 {
        return ans 
    }

    left := 0
    right := 1
    sum := 1
    endRight := target/2 + 1

    for right < endRight {
        // 移动右指针
        for right <= endRight && sum < target {
            right = right +1
            sum += right
        }

        if right > endRight {
            break
        }
        if sum == target {
            ans = append(ans, buildAns(left, right))
            // update
            right = right +1
            sum += right
        }

        for left + 1 < right && sum > target {
            left = left + 1
            sum -= left
        }

        if sum == target {
            ans = append(ans, buildAns(left, right))

            left = left +1
            sum -= left
        }
    }

    return ans
}

func buildAns(left, right int) []int {
    ans := make([]int, right - left)
    
    for i:= 0 ; i < len(ans) ; i ++ {
        ans[i] = left +1
        left += 1
    }
    return ans
}