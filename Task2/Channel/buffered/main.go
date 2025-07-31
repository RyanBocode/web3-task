/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

package main

import (
	"fmt"
	"sync"
)

// 生产者：发送 100 个整数到缓冲通道中
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch) // 发送完成后关闭通道
}

// 消费者：从通道中接收数据并打印
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("接收到：", num)
	}
}

func main() {
	ch := make(chan int, 20) // 创建一个缓冲大小为 20 的通道
	var wg sync.WaitGroup

	wg.Add(2) // 两个协程：生产者 + 消费者

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("所有任务完成")
}
