package main

import (
	"fmt"
)

func IsClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("recv success", ret)
}

func main() {
	c := make(chan int, 1)
	//go recv(c)
	c <- 1
	fmt.Println(<-c)
	fmt.Println(IsClosed(c))
	close(c)
	fmt.Println(IsClosed(c))
}
