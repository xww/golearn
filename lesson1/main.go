package main

import (
	"fmt"
	"reflect"
)

var x = 100

func main() {

	var a, s = 1000, "sadsd"
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(s))

	r := 1
	fmt.Println(r)

	x := 3
	fmt.Println(x)

	x, y := 1, 2
	x, y = y, x

}
