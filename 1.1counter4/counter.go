package main

import (
	"fmt"
	"sync"
)

// 可以对外隐藏锁的逻辑
// 把获取锁、释放锁、计数的逻辑封装成一个方法
func main() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.incr() // 受到锁的保护的方法
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

// 线程安全的计数器类型
type Counter struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

// 加1的方法,内部使用互斥锁保护
func (c *Counter) incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值,也需要锁的保护
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
