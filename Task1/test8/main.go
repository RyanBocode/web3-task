/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
*/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // value -> index

	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	return nil // 题目说一定有解，不会执行到这里
}

func main() {
	fmt.Println("示例 1：", twoSum([]int{2, 7, 11, 15}, 9)) // 输出: [0 1]
	fmt.Println("示例 2：", twoSum([]int{3, 2, 4}, 6))      // 输出: [1 2]
	fmt.Println("示例 3：", twoSum([]int{3, 3}, 6))         // 输出: [0 1]
}
