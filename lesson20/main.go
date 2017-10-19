package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	//channel

	event := make(chan struct{})
	c := make(chan string)

	go func() {
		s := <-c
		fmt.Println(s)
		close(event)
	}()

	time.Sleep(2 * time.Second)
	c <- "hi"
	<-event

}

func main1() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	c <- 3
	c <- 4
	fmt.Println(<-c)
	//fmt.Println(<-c)
	//cap len

	a := make(chan int, 1)
	b := make(chan int, 3)
	a <- 1
	b <- 11
	b <- 2
	fmt.Println("a:", len(a), cap(a))
	fmt.Println("b:", len(b), cap(b))

	value, ok := <-b
	fmt.Println(value, ok)
}

func main3() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for {
			x, ok := <-c
			if !ok {
				return
			}
			fmt.Println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}

func main4() {
	done := make(chan struct{})
	c := make(chan int)
	go func() {
		defer close(done)
		for x := range c {
			fmt.Println(x)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)

	<-done
}

func main5() {
	a := make(chan int, 3)
	a <- 1
	//a<-2
	close(a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	fmt.Println(<-a)
	//a<-2
}

func main6() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		defer wg.Done()
		for x := range recv {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 3; i++ {
			send <- i
		}
	}()
	wg.Wait()

	/*d := make(chan int,2)
	var send2 chan<- int = d
	var recv2 <-chan int = d

	<-send2*/


}
func main7() {
	//select
	var wg sync.WaitGroup
	a,b := make(chan int),make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for   {
			var name string
			var x int
			var ok bool
			select {
			case x,ok = <-a:
				name = "a"
			case x,ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			fmt.Println(name,x)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)

		for i:=0;i<10 ;i++  {
			select {
			case a<-i:
			case b<-i*10:

			}

		}
	}()
	wg.Wait()
}

func main() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("aa")
		}
	}
}
