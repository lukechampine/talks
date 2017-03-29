package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func unsafeEncode(x []uint64) []byte {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&x))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len * 8,
		Cap:  sh.Cap * 8,
	}))
}

func main() {
	// START OMIT
	x := make([]uint64, 100)
	res1 := testing.Benchmark(func(b *testing.B) {
		for i := b.N; i > 0; i-- {
			_ = unsafeEncode(x)
		}
	})
	res2 := testing.Benchmark(func(b *testing.B) {
		buf := make([]byte, 8)
		for i := b.N; i > 0; i-- {
			for j := range x {
				binary.LittleEndian.PutUint64(buf, x[j])
			}
		}
	})
	// END OMIT
	fmt.Printf("unsafe: %v/op\n", time.Duration(int64(res1.T)/int64(res1.N)))
	fmt.Printf("binary: %v/op\n", time.Duration(int64(res2.T)/int64(res2.N)))
}
