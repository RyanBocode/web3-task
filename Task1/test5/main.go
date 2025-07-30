/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	// 创建副本避免修改原数组
	result := make([]int, n)
	copy(result, digits)

	for i := n - 1; i >= 0; i-- {
		if result[i] < 9 {
			result[i]++
			return result
		}
		result[i] = 0 // 当前是9，加1进位
	}

	// 如果走到这说明全是9，比如 [9,9,9] → [1,0,0,0]
	newResult := make([]int, n+1)
	newResult[0] = 1
	return newResult
}

func main() {
	// Test the plusOne function
	testCases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{9},
		{9, 9, 9},
	}

	for _, digits := range testCases {
		result := plusOne(digits)
		fmt.Printf("%v + 1 = %v\n", digits, result)
	}
}
