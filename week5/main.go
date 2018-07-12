package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	// goroutine1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// goroutine2
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				close(ch)
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	fmt.Println("ok")
	fmt.Println("原因：因channel關閉後不能再傳值或取值")
	time.Sleep(time.Second * 100)
}