func translateNum(num int) int {
	// cou nt[i] = count[i-1]    if count[i-1:i+1] in the range
	//            count[i-1] + count[i-2]   if not
	
	if num < 10 {
		return 1
	}

	nums := []int{}
	for num != 0{
		temp := num %10
		nums = append(nums, temp)
		num = (num - temp)/10
	}

	for i, j := 0, len(nums) -1; i<j ; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	length := len(nums)
	count := make([]int, length)

	count[0] = 1
	for i:=1 ; i< length ; i++ {
		if i == 1 {
			if nums[0]*10 + nums[1] <= 25 {
				count[1] = 2
			}else {
				count[1] = 1
			}
			continue
		}

		temp := nums[i-1]*10 + nums[i]
		if temp > 25 || temp < 10 {
			count[i] = count[i-1]
		}else {
			count[i] = count[i-1] + count[i-2]
		}
	}

	return count[length -1]

}

// 滚动数组进行优化
func translateNumV2(num int) int {
    src := strconv.Itoa(num)
    p, q, r := 0, 0, 1
    for i := 0; i < len(src); i++ {
        p, q, r = q, r, 0
        r += q
        if i == 0 {
            continue
        }
        pre := src[i-1:i+1]
        if pre <= "25" && pre >= "10" {
            r += p
        }
    }
    return r
}