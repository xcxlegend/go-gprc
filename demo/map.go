package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m := make(map[string]string, 1)
	fmt.Println(len(m))
	var change = func(m map[string]string) {
		m = make(map[string]string, 2)
	}

	change(m)
	fmt.Println(len(m))

	s := make([]int, 0, 1)
	s = append(s, 1)
	fmt.Println(s)
	fmt.Println(cap(s))
	var changeSlice = func(s []int) []int {
		fmt.Println(unsafe.Pointer(&s))
		s = append(s, 2)
		return s
	}
	fmt.Println(unsafe.Pointer(&s))
	s = changeSlice(s)
	fmt.Println(unsafe.Pointer(&s))
	fmt.Println(s)
	fmt.Println(cap(s))

}
