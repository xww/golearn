package main

import (
	"fmt"
	"reflect"
	"github.com/pkg/errors"
)

func main() {
	a := 100
	p := &a
	fmt.Println(&a)
	fmt.Printf("%p, %v\n", &p, p)

	test(p)

	var a1 []int = []int{1, 2, 37, 40}

	test2("aaa", a1[:]...)

	fmt.Println(reflect.TypeOf(a1))
	b := a1[:]
	fmt.Println(reflect.TypeOf(b))

	test3(1)

}

func test(x *int) {
	fmt.Printf("%p, %v\n", &x, x)

}

func test2(s string, a ...int) { //slice
	fmt.Printf("%T,%v\n", a, a)

}

func test3(a int)( int,  string){
	age := a/3
	return age,"xww"
}
