package main

// Time Complexity: O(N)
// Space Complexity: O(1)
func trap(height []int) int {
	left, right := 0, len(height)-1
	maxL, maxR := height[left], height[right]
	sum := 0

	for left < right {
		if maxL < maxR {
			left++
			if height[left] > maxL {
				maxL = height[left]
			}
			sum += maxL - height[left]
		} else {
			right--
			if height[right] > maxR {
				maxR = height[right]
			}
			sum += maxR - height[right]
		}
	}

	return sum
}
