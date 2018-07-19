package main

import (
	"time"
	"fmt"
)

type CellPhone interface {
	Name() string
	Size() int
	TalkTime() time.Duration
}

type Iphone struct {
	version       string
	width, height int
	battery       time.Duration
}

func (i Iphone) Name() string {
	return i.version
}

func (i Iphone) Size() int {
	return i.height * i.width
}

func (i Iphone) TalkTime() time.Duration {
	return i.battery * time.Hour
}

type Pixel struct {
	version       string
	width, height int
	battery       time.Duration
}

func (p Pixel) Name() string {
	return p.version
}

func (p Pixel) Size() int {
	return p.height * p.width
}

func (p Pixel) TalkTime() time.Duration {
	return p.battery * time.Minute
}

//他是圓形的
type IWatch struct {
	version string
	radius  int
	battery time.Duration
}


func (w IWatch) Name() string {
	return w.version
}

func (w IWatch) Size() int {
	return 2 * w.radius * 3
}

func (w IWatch) TalkTime() time.Duration {
	return w.battery * time.Minute
}

func main() {
	iphone := Iphone{width: 30, height: 90, battery: 24, version: "iphone-XX"}
	pixel := Pixel{width: 40, height: 120, battery: 1300, version: "pixel-3"}
	iwatch := IWatch{radius: 4, battery: 50, version: "iwatch"}

	ShowPhone(iphone)
	ShowPhone(pixel)
	ShowPhone(iwatch)
}

func ShowPhone(c CellPhone) {
	fmt.Printf("Product %v \n", c.Name())
	fmt.Printf("Size %v \n", c.Size())
	fmt.Printf("Talk time %v \n", c.TalkTime())
	fmt.Println()
}
