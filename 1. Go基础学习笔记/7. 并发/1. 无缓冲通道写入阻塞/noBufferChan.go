package main

import (
	"time"
	"fmt"
)

func noBufferChanWrite(ch chan int) {
	ch <- 1
	fmt.Println("1 have been wrote in a No-buffer chan")    // 没有<-ch的情况下，会被阻塞，不会打印
}

func OneBufferChanWrite(ch chan<- int) {
	ch <- 1
	fmt.Println("1 have been wrote in a One-buffer chan")    // 没有被阻塞的打印语句
}

func main() {
	ch1 := make(chan int)
	go noBufferChanWrite(ch1)
	ch2 := make(chan int, 1)
	go OneBufferChanWrite(ch2)
	time.Sleep(5 * time.Second)
	fmt.Println("main end")
}