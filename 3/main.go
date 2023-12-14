package main

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	max := 0
	seen := make(map[byte]int)
	for l, r := 0, 0; r < len(s); r++ {
		//XXX: idx<l 重要！
		if idx, ok := seen[s[r]]; !ok || idx < l {
			seen[s[r]] = r
			curr := r - l + 1 //XXX
			if curr > max {
				max = curr
			}

		} else {
			l = idx + 1
			seen[s[r]] = r
		}
	}

	return max
}
