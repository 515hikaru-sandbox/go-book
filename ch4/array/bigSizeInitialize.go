package main

import (
	"fmt"
)

func main() {
	r := [...]int{100: -1}
	for i, v := range r {
		fmt.Println(i, v)
	}
}
