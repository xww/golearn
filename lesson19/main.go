package main

import (
	"runtime"
	"math"
	"fmt"
	"sync"
	"time"
)

func count()  {
	x := 0
	for i:=0;i<math.MaxInt32;i++{
		x += 1
	}
	fmt.Println(x)
}

func test1(n int)  {
	for i:=0;i<n;i++{
		count()
	}
}

func test2(n int)  {
	var wg sync.WaitGroup
	wg.Add(n)
	for i:=0;i<n;i++{
		//wg.Add(1)
		go func() {
			count()
			wg.Done()
		}()
	}
	wg.Wait()

}

func main() {
	//GOMAXPROCS
	//runtime.GOMAXPROCS(runtime.NumCPU())
	n := runtime.GOMAXPROCS(0)
	start := time.Now()
	fmt.Println(start)
	test1(n)
	test2(n)
	end := time.Now()
	fmt.Println(end)

	//Gosched


}
