package rocket

import (
	"time"
	"fmt"
	"sync"
)

type rocket interface {
	Name() string
	Size() int
	ReadyTime() int
}

type RocketDetail struct {
	Version       string
	Width, Height int
	Ready       int
}

func (r RocketDetail) Name() string {
	return r.Version
}

func (r RocketDetail) Size() int {
	return r.Height * r.Width
}

func (r RocketDetail) ReadyTime() int {
	return r.Ready
}

func Luncher (r rocket, wg *sync.WaitGroup) {
	fmt.Printf("火箭： %v 準備就緒 尺寸： %v 等待時間： %v 秒 \n", r.Name(), r.Size(), r.ReadyTime())
	ready := r.ReadyTime()
	for {
		if(ready == 0){
			wg.Done()
			fmt.Printf("火箭： %v 完成發射！ \n", r.Name())
			break
		}
		fmt.Printf("火箭： %v 倒數 %v 秒發射！ \n", r.Name(), ready)
		time_s := time.Duration(1) * time.Second
		time.Sleep(time_s)
		ready--
	}

}