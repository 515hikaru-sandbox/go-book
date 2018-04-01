package main

import (
	"fmt"
)

func main() {
	// 配列は直接比較可能
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [3]int{1, 2}
	fmt.Println(d)
	// fmt.println(a == d)
}
