package main

import (
	"fmt"
	"runtime"
)

func main() {
	//fmt.Println("aaa")
	runtime.GOMAXPROCS(1)
	exit := make(chan int)

	go func() {
		defer func() {
			//exit<-1
			close(exit)
		}()

		go func() {
			fmt.Println("b") //10 ms
		}()

		fmt.Println("a1")
		runtime.Gosched()
		fmt.Println("a2")

	}()
	<-exit

}
