package main

import (
	"fmt"
	"github.com/xww/golearn/lesson7/pkg"
)

func main() {
	sayHello()
	sayHello2()

	a := sayHello
	b := sayHello2
	a = b
	a()
	if a == nil {
		fmt.Println("nil")
	}

	var a1 *int = sayHello3()
	fmt.Println(a1, *a1)

	pkg.Simle()
}

func sayHello() {
	fmt.Println("hello ")

	a := func() { fmt.Println("222") }
	a()

}

func sayHello2() {
	fmt.Println("hello ")

	a := func() { fmt.Println("222") }
	a()
}

func sayHello3() *int {
	a := 100 //栈   堆
	return &a
}
