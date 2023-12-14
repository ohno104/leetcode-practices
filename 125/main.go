package main

import (
	"unicode"
)

// Time Complixity: O(2N)
// Space Complixity: O(1)
// func isPalindrome(s string) bool {
// 	f := func(r rune) rune {
// 		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
// 			return -1
// 		}

// 		return unicode.ToLower(r)
// 	}
// 	s = strings.Map(f, s)

// 	i, j := 0, len(s)-1
// 	for i < j {
// 		if s[i] != s[j] {
// 			return false
// 		}

// 		i = i + 1
// 		j = j - 1
// 	}

// 	return true
// }

// Time Complixity: O(N)
// Space Complixity: O(1)
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left++
		} else if !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
		} else if unicode.ToLower(rune(s[left])) == unicode.ToLower(rune(s[right])) {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}
