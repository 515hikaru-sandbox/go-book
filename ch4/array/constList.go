package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "gbp", RMB: "&yen;"}
	fmt.Printf("%d, %d, %d, %d\n", USD, EUR, GBP, RMB)
	fmt.Println(symbol)
}
