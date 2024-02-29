package main

// Time Complexity: O(N)
// Space Complexity: O(1)
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	start := 0
	res := len(nums) + 1 // +1用來判斷有沒有符合條件

	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			subLength := end - start + 1
			if subLength < res {
				res = subLength
			}
			sum -= nums[start] // 先減才能移動
			start++
		}
	}

	if res > len(nums) {
		return 0
	}

	return res
}
