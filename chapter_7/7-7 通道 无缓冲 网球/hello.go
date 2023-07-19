package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
在网球比赛中，两位选手会把球在两个人之间来回传递。选手总是处在以下两种状态之一：
要么在等待接球，要么将球打向对方。可以使用两个 goroutine 来模拟网球比赛，并使用无缓冲
的通道来模拟球的来回
*/

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	court := make(chan int)

	wg.Add(2)

	// 启动两个选手
	go player("Nadal", court)

	go player("Dijav", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()

}

func player(name string, court chan int) {

	defer wg.Done()

	for {

		ball, ok := <-court

		if !ok {

			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数 判断我们是否丢球

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道 输了
			close(court)

			return
		}

		// 显示击球数
		fmt.Printf("Player %s Hit %d\n", name, ball)

		ball++

		// 将球打向对手
		court <- ball

	}

}
