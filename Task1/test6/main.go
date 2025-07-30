/*
给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
由于在某些语言中不能改变数组的长度，因此必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
将最终结果插入 nums 的前 k 个位置后返回 k 。
不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
判题标准:
系统会用下面的代码来测试你的题解:
int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案
int k = removeDuplicates(nums); // 调用
*/
package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}

	for idx, nums := range testCases {
		original := append([]int(nil), nums...) // 复制原数组用于打印
		k := removeDuplicates(nums)
		fmt.Printf("示例 %d: 输入: %v\n       输出: 长度=%d, nums=%v\n",
			idx+1, original, k, nums[:k])
	}
}
