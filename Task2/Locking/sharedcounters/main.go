/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int       // 共享计数器
	var mu sync.Mutex     // 定义互斥锁
	var wg sync.WaitGroup // 用于等待所有协程完成

	numGoroutines := 10
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines) // 添加10个协程任务

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				mu.Lock()   // 加锁
				counter++   // 安全地递增
				mu.Unlock() // 解锁
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
	fmt.Println("最终计数器的值：", counter)
}
