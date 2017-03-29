package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"golang.org/x/sys/unix"
)

func toBytes(x int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&x)),
		Len:  int(unsafe.Sizeof(x)),
		Cap:  int(unsafe.Sizeof(x)),
	}))
}

// STARTIMPL OMIT
func freeze(x *int) *int {
	// get underlying bytes of x
	bytes := toBytes(*x)

	// allocate new memory and copy x's data into it
	frozen, _ := unix.Mmap(-1, 0, len(bytes), unix.PROT_READ|unix.PROT_WRITE,
	                                          unix.MAP_ANON|unix.MAP_PRIVATE)
	copy(frozen, bytes)

	// freeze the memory
	unix.Mprotect(frozen, unix.PROT_READ)

	// return the new memory as an *int
	return (*int)(unsafe.Pointer(&frozen[0]))
}

// ENDIMPL OMIT

// START OMIT
func main() {
	x := new(int)
	*x++ // ok

	x = freeze(x)

	fmt.Println(*x) // ok; prints 1
	//*x++            // not ok; panics!
}

// END OMIT
