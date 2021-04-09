func missingNumber(nums []int) int {
    ans := nums[0]
    for i:= 1 ; i< len(nums) ; i ++ {
        ans ^= nums[i]
    }
    for i:= 0 ; i<= len(nums) ; i ++ {
        ans ^=i
    }
    return ans
}