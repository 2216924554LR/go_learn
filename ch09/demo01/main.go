package main

import (
	"time"
	"unsafe"
)

type data struct {
	x [1024 * 100]byte
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

func main() {
	const N = 10000
	cache := new([N]uintptr)
	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}
}
