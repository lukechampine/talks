package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//START OMIT
	b := []byte("foo")
	s := *(*string)(unsafe.Pointer(&b))

	copy(b, "bar")
	fmt.Println(string(b), s)

	b = append(b, 'n')
	fmt.Println(string(b), s)
	//END OMIT
}
