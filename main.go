package main

import (
	"fmt"
)

func main() {
	s := "a)b(c)d"

	stack := make([]int, 0)
	var res []byte

	for i, c := range []byte(s) {
		if c == '(' {
			stack = append(stack, i)
			res = append(res, c)
		} else if c == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				res = append(res, c)
			} else {
				res = append(res, ' ')
			}

		} else {
			res = append(res, c)
		}
	}
	fmt.Println("res:", res)
	fmt.Println("stack:", stack)
	for _, v := range stack {
		res[v] = ' '
	}
	fmt.Println("res:", res)
	var res2 string
	for i := range res {
		if res[i] != ' ' {
			res2 += string(res[i])
		}
	}

	fmt.Println(res2)
}
