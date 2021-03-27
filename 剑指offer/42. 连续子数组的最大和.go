func maxSubArray(nums []int) int {
    max := nums[0]
    for index := range nums {
        if index > 0 {
            if nums[index] < nums[index] + nums[index -1] {
                nums[index] = nums[index] + nums[index -1]
            }
            if nums[index] > max {
                max = nums[index]
            }
        }
    }
    return max
}