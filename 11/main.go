package main

// Time Complexity: O(N)
// Space Complexity: O(1)
func maxArea(height []int) int {
	max := 0
	idxL := 0
	idxR := len(height) - 1

	for idxL <= idxR {
		area := 0
		if height[idxL] < height[idxR] {
			area = height[idxL] * (idxR - idxL)
			idxL++
		} else {
			area = height[idxR] * (idxR - idxL)
			idxR--
		}

		if area > max {
			max = area
		}
	}

	return max
}
