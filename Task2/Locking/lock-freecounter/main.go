/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // 使用 int64 类型用于原子操作
	var wg sync.WaitGroup

	numGoroutines := 10
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增
			}
		}()
	}

	wg.Wait()
	fmt.Println("最终计数器的值：", counter)
}
