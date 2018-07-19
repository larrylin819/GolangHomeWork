package main

import (
	r "rocket"
	"sync"
	"fmt"
)

func main() {
	//建立WaitGroup
	var wg sync.WaitGroup
	r1 := r.RocketDetail{Width: 10, Height: 20, Ready: 5, Version: "天弓三型飛彈"}
	r2 := r.RocketDetail{Width: 30, Height: 100, Ready: 10, Version: "巡弋飛彈"}
	wg.Add(2)
	go r.Luncher(r1, &wg)
	go r.Luncher(r2, &wg)
	//等待所有工作完成
	wg.Wait()
	fmt.Println("所有飛彈皆已成功發射！")
}
