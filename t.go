package main

import (
	"fmt"
	"unsafe"
)

type P struct {
	name string
	My   string
}

type B struct {
	name string
	My   string
}

func main() {
	p := &P{
		"1",
		"my",
	}

	b := (*B)(unsafe.Pointer(p))
	b.name = "32"
	fmt.Println(p, b)
}
