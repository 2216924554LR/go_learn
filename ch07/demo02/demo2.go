package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("recv over")
		exit <- true
	}()

	data <- 1
	data <- 2
	data <- 3
	time.Sleep(2 * time.Second)
	close(data)
	fmt.Println("send over")
	<-exit
}
