package pool_v1

/*** 示例
func main(){
    pool := pool_v1.New(100)
    println(runtime.NumGoroutine())
    for i := 0; i < 1000; i++ {
        pool.Add(1) // 协程达到上限了将阻塞
        go func() {
            time.Sleep(time.Second)
            println(runtime.NumGoroutine())
            pool.Done()
        }()
    }
    pool.Wait() // 等待所有的结束
    println(runtime.NumGoroutine())
}
*/

import "sync"

// 结构体1: 协程池
type Pool struct {
	queue chan int
	num   int
	wg    *sync.WaitGroup
}

// 名称：创建协程池
// @size 允许并发个数
func New(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
		num:   0,
	}
}

// 名称：添加协程
// @delta 个数
func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	p.num += delta
	p.wg.Add(delta)
}

// 名称：完成一个协程
func (p *Pool) Done() {
	p.num--
	<-p.queue
	p.wg.Done()
}

// 名称：完成一个协程
func (p *Pool) Len() int {
	return p.num
}

// 名称：等待所有协程完成
func (p *Pool) Wait() {
	p.wg.Wait()
}
