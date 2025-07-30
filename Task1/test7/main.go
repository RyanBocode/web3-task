/*
以数组 intervals 表示若干个区间的集合，
其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，
该数组需恰好覆盖输入中的所有区间 。
*/
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按起始时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	for _, interval := range intervals {
		n := len(res)
		if n == 0 || res[n-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			// 合并区间
			res[n-1][1] = max(res[n-1][1], interval[1])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals2 := [][]int{{1, 4}, {4, 5}}

	fmt.Println("示例 1 合并结果:", merge(intervals1)) // [[1 6] [8 10] [15 18]]
	fmt.Println("示例 2 合并结果:", merge(intervals2)) // [[1 5]]
}
