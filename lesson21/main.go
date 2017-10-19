package main

import (
	"sync"
	"fmt"
	"time"
)

func main1() {

	//mutex
	var wg sync.WaitGroup
	wg.Add(2)
	var m sync.Mutex
	a := 0
	go func() {
		m.Lock()
		for i:=0;i<5000;i++{
			a=a+1
		}
		wg.Done()
		m.Unlock()
	}()
	go func() {
		m.Lock()
		for i:=0;i<5000;i++{
			a=a+1
		}
		wg.Done()
		m.Unlock()
	}()

	wg.Wait()
	fmt.Println(a)
}

type data struct {
	sync.Mutex

}

func (d *data)test(s string)  {
	d.Lock()
	for i := 0; i<5;i++{
		fmt.Println(s,i)
		time.Sleep(time.Second)
	}
	d.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data
	go func() {
		//d.Lock()
		defer wg.Done()
		d.test("read")
	}()
	go func() {
		defer wg.Done()
		d.test("write")
	}()
	wg.Wait()


}
