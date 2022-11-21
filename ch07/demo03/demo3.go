package main

import "fmt"

func main() {
	data := make(chan int, 3)
	exit := make(chan bool)
	fmt.Println("Len of chan: ", len(data))
	data <- 1
	data <- 2
	data <- 3
	fmt.Println("Len of chan: ", len(data))

	go func() {
		for d := range data {
			fmt.Println(d)
		}
		exit <- true
	}()

	data <- 4
	data <- 5
	close(data)
	<-exit
}
