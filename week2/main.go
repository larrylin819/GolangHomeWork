package main

import "fmt"

func main(){
	cnt := 0
	total := 0
	ax := []int {60, 61, 62, 63, 65}
	for _,value := range ax {
		total += value
		cnt += 1
	}
	total2 := float64(total)
	cnt2 := float64(cnt)
	fmt.Println(total2/cnt2)

	x := [ ] int {
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	tmp := 999999999999
	for _,value := range x {
		if tmp > value {
			tmp = value
		}
	}
	fmt.Println(tmp)

}