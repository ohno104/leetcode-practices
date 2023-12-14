package main

// Time Complexity: O(N)
// Space Complexity: O(N)
func twoSum(nums []int, target int) []int {
	rec := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if j, ok := rec[target-nums[i]]; ok {
			return []int{i, j}
		}
		rec[nums[i]] = i
	}

	return []int{}
}
