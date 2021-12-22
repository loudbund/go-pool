package main

import (
	"github.com/loudbund/go-pool/pool_v1"
	"runtime"
	"time"
)

func main() {
	// 创建100个并发池
	pool := pool_v1.New(10)
	println(pool.Len(), runtime.NumGoroutine())

	// 循环执行1千个协程
	for i := 0; i < 100; i++ {
		pool.Add(1) // 协程达到上限了将阻塞
		go func(n int) {
			time.Sleep(time.Second)
			println(pool.Len(), runtime.NumGoroutine(), n)
			pool.Done()
		}(i)
	}

	// 等待所有的结束
	pool.Wait()
	println(pool.Len(), runtime.NumGoroutine())
}
