/*
编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信
*/

package main

import (
	"fmt"
	"sync"
)

// 生产者：将 1 到 10 的整数发送到通道中
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成时通知 WaitGroup

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // 发送完成后关闭通道
}

// 消费者：从通道接收数据并打印
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 协程完成时通知 WaitGroup

	for num := range ch {
		fmt.Println("接收到：", num)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2) // 添加两个等待任务

	go producer(ch, &wg) // 启动生产者
	go consumer(ch, &wg) // 启动消费者

	wg.Wait() // 等待所有协程完成
	fmt.Println("所有任务完成")
}
