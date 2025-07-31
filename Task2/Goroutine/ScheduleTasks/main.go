/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，
同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义任务类型：无参数、无返回值的函数
type Task func()

// 调度器函数，接收任务切片并并发执行，记录每个任务的执行时间
func ScheduleTasks(tasks []Task) {
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)

		// 为每个任务启动一个协程
		go func(index int, t Task) {
			defer wg.Done()

			start := time.Now()
			t() // 执行任务
			elapsed := time.Since(start)

			fmt.Printf("任务 %d 执行时间：%v\n", index+1, elapsed)
		}(i, task)
	}

	wg.Wait() // 等待所有任务完成
}

func main() {
	// 定义多个模拟任务
	tasks := []Task{
		func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("任务 1 完成")
		},
		func() {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("任务 2 完成")
		},
		func() {
			time.Sleep(400 * time.Millisecond)
			fmt.Println("任务 3 完成")
		},
	}

	// 调用调度器执行任务
	ScheduleTasks(tasks)

	fmt.Println("所有任务执行完毕")
}
