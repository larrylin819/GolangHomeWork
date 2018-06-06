package main

import (
	"fmt"
	"strconv"
	srcc "./src"
)

func main() {

	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x string = "I Love Golang_"

	fmt.Println(a+int(b))
	fmt.Println(a+int(b)+int(c))
	fmt.Println(f/float64(e))
	fmt.Println(a+strToInt(d))
	fmt.Println(x+srcc.IntToStr(a))
}

func strToInt(str string) int {
	int, _ := strconv.Atoi(str)
	return int
}