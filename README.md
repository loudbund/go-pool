# go-pool
go-pool是用来控制并发协程数量和等待协程执行完成。4部可以实现并发控制和等待执行结束
#####1、初始化设置并发数量
#####2、启动协程时并发+1
#####3、结束一个协程时并发-1
#####4、等待所有协程结束

# 安装
go get github.com/loudbund/go-pool

# 使用
```golang
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

```