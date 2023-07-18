package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int

	wg sync.WaitGroup
)

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

		value := counter

		runtime.Gosched()

		value++

		counter = value
	}
}
