package ch01

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestDemo2(t *testing.T) {
	var x struct {
		a int
		b int
		c []int
	}
	pb := &x.b
	*pb = 42
	fmt.Println(x.b)
}

func TestDemo2_2(t *testing.T) {
	v1 := uint(12)
	v2 := int(13)

	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(reflect.TypeOf(&v1))

	p := &v1
	p = (*uint)(unsafe.Pointer(&v2))
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(*p)

}
