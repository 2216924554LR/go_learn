package defer_demo

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a: ", i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b: ", i)
			}()
		}
	}()
}

func Test2(t *testing.T) {
	var a = 123
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a)
	}(a)
	a = 789
	time.Sleep(2 * time.Second)

}

func Test3(t *testing.T) {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}
