package main

import (
	"fmt"
	"time"
	"sync"
)

func task1(s string){
	fmt.Println("task1 start")
	time.Sleep(1*time.Second)
	fmt.Println("task1 end",s)
}

func task2(){
	fmt.Println("task2 start")
	time.Sleep(2*time.Second)
	fmt.Println("task2 end")
}

func main() {
	//goroutine =协程 或者 线程 go runtime
	fmt.Println(time.Now())
	//exit := make(chan int)
	/*go task1("task111")
	go task2()*/
	var wg sync.WaitGroup
	wg.Add(1)
	go func(s string) {
		fmt.Println(" start task3")
		time.Sleep(2*time.Second)
		fmt.Println("end task3",s)
		//exit<-1
		wg.Done()
	}("task333")
	//time.Sleep(100*time.Second)
	fmt.Println(time.Now())
	//<-exit
	wg.Wait()

}
