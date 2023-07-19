package main

import (
	"fmt"
	"sync"
	"time"
)

/*
在接力比赛里，4 个跑步者围绕赛道轮流跑（如代码清单 6-22 所示）。第二个、第三个和
第四个跑步者要接到前一位跑步者的接力棒后才能起跑。比赛中最重要的部分是要传递接力棒，
要求同步传递。在同步接力棒的时候，参与接力的两个跑步者必须在同一时刻准备好交接
*/

var (
	wg sync.WaitGroup
)

func init() {

}

func main() {
	baton := make(chan int)

	//
	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {

	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d Running with Baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1

		fmt.Printf("Runner %d To The Line\n", newRunner)

		go Runner(baton)
	}

	time.Sleep(100 * time.Microsecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)

		wg.Done()

		return
	}

	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner

}
