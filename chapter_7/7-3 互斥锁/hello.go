package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64

	wg sync.WaitGroup

	mutex sync.Mutex
)

/*
原子函数能够以很底层的加锁机制来同步访问整型变量和指针
*/

func main() {

	wg.Add(2)

	go inCounter(1)

	go inCounter(2)

	wg.Wait()

	fmt.Println("Final Counter:", counter) //4

}

func inCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {

		mutex.Lock()

		{
			value := counter

			runtime.Gosched()

			value++

			counter = value
		}

		mutex.Unlock()
	}
}
