package main

import (
	"crypto/sha256"
	"fmt"
)

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func main() {
	c := sha256.Sum256([]byte("x"))
	zero(&c)
	fmt.Printf("zero1: %x\n", c)
	c = sha256.Sum256([]byte("x"))
	zero2(&c)
	fmt.Printf("zero2: %x\n", c)
}
