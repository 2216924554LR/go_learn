package ch01

import (
	"fmt"
	"testing"
)

var x, y, z int

var (
	a int
	b float32
)

func TestDemo1(t *testing.T) {
	n, s := 0x1234, "Hello, World!"
	println(x, y, z)
	println(s, n)
	println(a, b)

}

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {

}

func TestDemo1_2(t *testing.T) {
	c := Black
	test(c)

	//x := 1
	//test(x)
	test(1)
}

func TestDemo1_3(t *testing.T) {
	s := "abc汉子"
	for n, r := range s {
		fmt.Println(n)
		fmt.Printf("%c\n", r)
	}
}
