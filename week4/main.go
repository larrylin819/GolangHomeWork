package main

import (
	"time"
	"fmt"
	"sync"
)

type factory struct {
	name  string //功能
	mp   int    //消耗產生紙漿速率
	pp   int    //產生印製紙速率
}

func main() {
	//伐木場
	wood := factory{
		name:  "伐木場",
		mp:   1000,
	}
	//造紙場1
	paper1 := factory{
		name:  "造紙場1",
		mp:   500,
		pp:   5000,
	}
	//造紙場2
	paper2 := factory{
		name:  "造紙場2",
		mp:   300,
		pp:   3000,
	}
	//印刷廠
	print := factory{
		name:  "印刷廠",
		pp:   6000,
	}

	//紙漿
	pulp := 0
	//白紙
	blank := 0
	//成品
	finishPaper := 0

	//建立WaitGroup
	var wg sync.WaitGroup

	//建立互斥鎖
	var mutex sync.Mutex

	wg.Add(1)

	//使用不同工廠
	go woodFac(wood, &pulp, &finishPaper, mutex)
	go paperFac(paper1, &pulp, &blank, &finishPaper, mutex)
	go paperFac(paper2, &pulp, &blank, &finishPaper, mutex)
	go printFac(print, &blank, &finishPaper, &wg, mutex)

	//等待所有工作完成
	wg.Wait()

	//所有工作都完成
	fmt.Printf("%c[0;40;36m已全部製作完成%c[0m\n", 0x1B, 0x1B)


}

func woodFac(w factory, p *int, fp *int, mutex sync.Mutex) {
	for i := *fp; i < 60000; i = *fp {

		//紙漿數量上鎖
		mutex.Lock()

		*p = *p + w.mp

		fmt.Printf("%s產生%d公斤紙漿,目前數量 紙漿:%d公斤 成品:%d張\n", w.name, w.mp, *p, *fp)

		//解鎖
		mutex.Unlock()

		time_s := time.Duration(1) * time.Second
		time.Sleep(time_s)
	}
}


func paperFac(w factory, p *int, b *int, fp *int, mutex sync.Mutex) {
	for i := *fp; i < 60000; i = *fp {

		//紙漿數量上鎖
		mutex.Lock()

		if *p < w.mp {
			mutex.Unlock()
			continue
		}

		*p = *p - w.mp
		*b = *b + w.pp

		fmt.Printf("%s產生%d白紙,目前數量 紙漿:%d公斤 白紙:%d張 成品:%d張\n", w.name, w.pp, *p, *b, *fp)

		//解鎖
		mutex.Unlock()

		time_s := time.Duration(1) * time.Second
		time.Sleep(time_s)
	}
}


func printFac(w factory,b *int, fp *int, wg *sync.WaitGroup, mutex sync.Mutex) {
	for i := *fp; i < 60000; i = *fp {

		//紙漿數量上鎖
		mutex.Lock()

		if *b < w.pp {
			mutex.Unlock()
			continue
		}

		*b = *b - w.pp
		*fp = *fp + w.pp

		fmt.Printf("%s產生%d張成品,目前數量 白紙:%d張 成品:%d張\n", w.name, w.pp, *b, *fp)

		//解鎖
		mutex.Unlock()

		time_s := time.Duration(1) * time.Second
		time.Sleep(time_s)
	}
	//工作完成 回報WaitGroup -1
	wg.Done()
}