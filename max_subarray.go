package coding_kata

func MaxSumOfSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sum := make([]int, len(nums))
	min := nums[0]
	max := nums[0]
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
		if sum[i] > max{
			max = sum[i]
		}
		if sum[i]-min > max {
			max = sum[i] - min
		}
		if min > sum[i] {
			min = sum[i]
		}
	}
	return max
}
