package main

// Time Complexity: O(a+b)
// Space Complexity: O(a+b)
// func backspaceCompare(s string, t string) bool {
// 	fs := stringBuild(s)
// 	ft := stringBuild(t)

// 	if len(fs) != len(ft) {
// 		return false
// 	}

// 	for i := 0; i < len(fs); i++ {
// 		if fs[i] != ft[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func stringBuild(str string) string {
// 	var res []byte //rune

// 	for i := 0; i < len(str); i++ {
// 		if str[i] != '#' {
// 			res = append(res, str[i])
// 		} else {
// 			if len(res) > 0 {
// 				res = res[:len(res)-1]
// 			}
// 		}
// 	}

// 	return string(res)
// }

// Time Complexity: O(a+b)
// Space Complexity: O(1)
func backspaceCompare(s string, t string) bool {
	index1, index2 := len(s)-1, len(t)-1
	for index1 >= 0 || index2 >= 0 {
		i1 := getNextValidCharIndex(s, index1)
		i2 := getNextValidCharIndex(t, index2)
		if i1 < 0 && i2 < 0 {
			return true
		}
		if i1 < 0 || i2 < 0 {
			return false
		}
		if s[i1] != t[i2] {
			return false
		}

		index1 = i1 - 1
		index2 = i2 - 1
	}

	return true
}

func getNextValidCharIndex(s string, index int) int {
	backSpaceCount := 0
	for index >= 0 {
		if s[index] == '#' {
			backSpaceCount++
		} else if backSpaceCount > 0 {
			backSpaceCount--
		} else {
			break
		}
		index--
	}
	return index
}
