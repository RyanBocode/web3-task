/*
题目 ：编写一个Go程序，使用协程实现两个协程交替打印奇数和偶数，
要求在主函数中启动两个协程，并等待它们完成。
考察点 ：协程的创建与调度、通道的使用。
*/
package main

import (
	"fmt"
	"time"
)

// 打印奇数的协程
func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数：", i)
		time.Sleep(30 * time.Millisecond) // 模拟延时，增强交替打印效果
	}
}

// 打印偶数的协程
func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数：", i)
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	go printOdd()  // 启动奇数打印协程
	go printEven() // 启动偶数打印协程

	time.Sleep(1 * time.Second) // 等待所有协程执行完，防止主程序提前退出
	fmt.Println("主程序结束")
}
