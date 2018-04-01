package main

import (
	"fmt"
)

func print_array(a [3]int) {
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func main() {
	var q [3]int = [3]int{1, 2, 3}
	fmt.Println("print: ", q)
	print_array(q)
	var r [3]int = [3]int{1, 2}
	fmt.Println("print: ", r)
	print_array(r)
	q = [3]int{1, 2, 3}
	fmt.Printf("%T\n", q)
}
