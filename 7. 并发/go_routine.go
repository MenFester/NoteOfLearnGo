package main

import (
	"time"
	"fmt"
)

func goRoutine1(ch chan int) {
	<- ch
	fmt.Println("hello, go routine1")
}

func goRoutine2(ch chan int) {
	fmt.Println("hello, go routine2")
	ch <- 1
	time.Sleep(2 * time.Second)
	fmt.Println("end go routine2")
}

func sendData(ch chan<- int) {
	ch <- 10
}

func main() {
	ch := make(chan int)
	go goRoutine1(ch)
	time.Sleep(3 * time.Second)
	go goRoutine2(ch)
	fmt.Println("main routine")

	// chSend := make(chan<- int)    // 执行fmt.Println(<-chSend)时报错
	chSend := make(chan int)
	go sendData(chSend)
	fmt.Println(<-chSend)
	
	time.Sleep(3 * time.Second)    // 没有主协程的等待，其他协程的打印输出行为不确定
}