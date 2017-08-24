package main

import (
	"fmt"
	"math"
)

type Student struct {
	Age  int
	Name string
}

func main() {

	/*
		bool
		byte
		int 8 16 32 64 int8 int16
		uint
		float8 16 32 64
		complex64 128 1+2i
		uintptr
		string
		array

		struct
		function
		interface

		map
		slice
		channel
	*/

	s := "0Abc"
	fmt.Println([]byte(s))

	//b := false
	stu := new(Student)
	fmt.Println(*stu)

	fmt.Println(math.MinInt8, math.MaxUint16)

	m := make(map[int]string)
	m[1] = "1"
	m[2] = "2"
	fmt.Println(m)

	//s1 := make([]int,0)
	//s1 = append(s1,12)
	//fmt.Println(s1)

	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)

	/*p:=new(map[int]string)
	m1:=*p
	m1[1]="1"
	fmt.Println(m1)*/

	a := 1
	b := uint8(a)
	fmt.Println(b)

	type color int

	var red color
	red = 1
	var green color
	green = 2
	fmt.Println(red, green)

	/*red1 := int(1)
	red1 = red
	fmt.Println(red1)*/

	//reflect.Kind

}
