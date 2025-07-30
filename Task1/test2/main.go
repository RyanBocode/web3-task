/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/
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
	// Test the isPalindrome function
	testCases := []int{121, -121, 10, 12321, 12345}

	for _, num := range testCases {
		result := isPalindrome(num)
		fmt.Printf("%d is palindrome: %t\n", num, result)
	}
}
