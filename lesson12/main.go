package main

import (
	"unsafe"
	"fmt"
)

type slice struct {
	array unsafe.Pointer
	len int
	cap int
}

func main() {
	x := [...]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	s := x[2:5]
	fmt.Println(len(s))
	fmt.Println(cap(s))
	//fmt.Println(s[4])
	s1 := make([]int, 3, 5)
	fmt.Println(s1)

	var a []int
	b := []int{}

	fmt.Println(a==nil,b==nil)

	//s2 := make()
	s2 := append(s1,10)
	s2 = append(s2,10)
	s2 = append(s2,10)
	fmt.Println(s2)
	fmt.Println(s1)

	s3 := x[2:4]
	n:=copy(x[3:],s3)
	fmt.Println(n)
	fmt.Println(x)





}
