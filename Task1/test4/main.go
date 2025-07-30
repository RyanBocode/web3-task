/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	fmt.Println("示例 1：", longestCommonPrefix([]string{"flower", "flow", "flight"})) // 输出: "fl"
	fmt.Println("示例 2：", longestCommonPrefix([]string{"dog", "racecar", "car"}))    // 输出: ""
}
