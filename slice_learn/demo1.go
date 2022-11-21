package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a cap: %d\n", cap(a))
	fmt.Printf("a len: %d\n", len(a))
	fmt.Printf("a : %d\n", a)
	fmt.Printf("a address: %p \n", &a)
	fmt.Printf("a pointer: %d\n", &a[0])

	b := []int{}
	fmt.Printf("b cap: %d\n", cap(b))
	fmt.Printf("b len: %d\n", len(b))
	fmt.Printf("b : %d\n", b)
	fmt.Printf("b address: %p \n", &b)

}
