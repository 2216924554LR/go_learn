package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	// sem := make(chan int, 1)
	for i := 0; i < 3; i++ {
		//time.Sleep(time.Second)
		fmt.Println("current i: ", i)
		go func(id int) {
			fmt.Println("current id: ", id)
			defer wg.Done()
			// sem <- 1
			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
			}
			// <-sem
		}(i)
	}
	wg.Wait()
}
