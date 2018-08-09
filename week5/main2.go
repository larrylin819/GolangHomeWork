package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	total := 1
	ch := make(chan int, 20)
	// goroutine1
	go func() {
		for i := 1; i <= 20; i++ {
			total = total * i
			//fmt.Printf("%d!: %d\n", i, total)
			ch <- total
		}
	}()

	tmp2 := 0
	total2 := 0
	for x := 1; x <= 20; x++ {
		tmp2 = <-ch
		total2 += tmp2
	}
	fmt.Println(total2)

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("time: %s\n", delta)

	fmt.Println("ok")
}
