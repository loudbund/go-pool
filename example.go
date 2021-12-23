package main

import (
	"github.com/loudbund/go-pool/pool_v1"
	"runtime"
	"time"
)

func main() {
	// step1、创建10个并发池
	pool := pool_v1.New(10)
	println(pool.Len(), runtime.NumGoroutine())

	// 循环执行1千个协程
	for i := 0; i < 100; i++ {
		// step2、启动协程前加1，协程达到上限了将阻塞
		pool.Add(1)
		go func(n int) {
			// step3、协程结束时减1
			defer func() { pool.Done() }()
			time.Sleep(time.Second)
			println(pool.Len(), runtime.NumGoroutine(), n)
		}(i)
	}

	// step4、等待所有的结束
	pool.Wait()
	println(pool.Len(), runtime.NumGoroutine())
}
