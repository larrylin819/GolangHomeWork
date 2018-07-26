package main

import (
	"phone"
)

type iphone1 struct {
	user string
	passwd string
	volume int
}

type iphone2 struct {
	user string
	passwd string
	volume int
}

func (i1 iphone1) CheckName() string {
	return i1.user
}

func (i1 iphone1) CheckPass() string {
	return i1.passwd
}

func (i1 iphone1) CheckVolume() int {
	return i1.volume
}

func (i2 iphone2) CheckName() string {
	return i2.user
}

func (i2 iphone2) CheckPass() string {
	return i2.passwd
}

func (i2 iphone2) CheckVolume() int {
	return i2.volume
}

func main() {
	i1 := iphone1{user:"Larry", passwd:"0000", volume:10}
	i2 := iphone2{user:"Larry2", passwd:"1234", volume:20}
	phone.CheckUser(i1)
	phone.CheckUser(i2)
	phone.VolumeNum(i1)
	phone.VolumeNum(i2)
}

