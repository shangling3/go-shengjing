package main

import (
	"sync"
)

var sum = 0
var mutex sync.RWMutex

func readSum() int {
	// 只获取读锁
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
func run() {
	var wg sync.WaitGroup
	wg.Add(110)
	go func() {
		defer wg.Done()
		add(10)
	}()
	for i := 0; i < 100; i++ {
		go add(10)
	}

	for i := 0; i < 10; i++ {
		go func() {
			// 计数器值减 1
			defer wg.Done()
			println("和为：", readSum())
		}()
		//go fmt.Println("和为；", readSum())
	}
	//time.Sleep(2 * time.Second)
	wg.Wait()
}
func main() {
	run()
}
