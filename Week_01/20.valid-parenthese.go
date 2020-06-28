package main

import "math"

var a = math.MaxInt32

func isValid(s string) bool {
	n := len(s)

	m := map[byte]byte{')': '(', '}': '{', ']': '['}

	stack := make([]byte, n)
	p := 0

	for i := 0; i < n; i++ {
		switch s[i] {
		case '(', '{', '[':
			stack[p] = s[i]
			p++
		default:
			if p == 0 {
				return false
			}

			p--
			left := stack[p] // 返回栈顶元素

			if left != m[s[i]] {
				return false
			}
		}

	}

	return p == 0
}


