package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x) // 把整数转成字符串

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	// 测试回文数函数
	testCases := []int{121, -121, 10, 12321, 12345}

	for _, num := range testCases {
		result := isPalindrome(num)
		fmt.Printf("%d 是回文数: %t\n", num, result)
	}
} 