package phone

import (
	"fmt"
)

type phone interface {
	CheckName() string
	CheckPass() string
	CheckVolume() int
}

func CheckUser (p phone) {
	if(p.CheckName() == "Larry" && p.CheckPass() == "0000") {
		fmt.Println(p.CheckName(), "驗證通過")
		return
	}
	fmt.Println(p.CheckName(), "驗證失敗")
}

func VolumeNum (p phone) {
	fmt.Println("目前音量：", p.CheckVolume())
}