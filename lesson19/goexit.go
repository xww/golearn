package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Goexit defer panic
	exit := make(chan int)
	go func() {
		defer close(exit)
		defer fmt.Println("a")

		func() {
			defer func() {
				fmt.Println("b", recover() == nil)
			}()

			func() {
				fmt.Println("c")
				runtime.Goexit()
				fmt.Println("c exit")
			}()
			fmt.Println("b exit")
		}()
		fmt.Println("a exit")
	}()
	<-exit
	fmt.Println("main exit")
}
