package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)
	m["a"]=0
	fmt.Println(m)

	v,ok := m["a"]
	fmt.Println(ok,v)

	v,ok = m["c"]
	if ok{
		fmt.Println(ok,v)
	}

	m["abc"] = 123

	fmt.Println(len(m))

	for k,v:= range m{
		fmt.Println(k,v)
	}

	m1 := make(map[string]int)

	go func() {
		for   {
			m1["a"] = 1
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for   {
			_ = m1["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}



}

