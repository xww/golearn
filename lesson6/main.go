package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func check(x int) error {
	if x <= 0 {
		return errors.New("x<0")
	}
	return nil
}

func main() {
	/*//x:=3
	x:=6
	if x>5{
		fmt.Println(x)
	}else if x >7 {
		fmt.Println(x)
	}*/
	x := -3
	if err := check(x); err == nil {

		x++
		fmt.Println(x)

	} else {
		fmt.Println(err)
	}

	//fmt.Println(err)
	//while do wile
	for i := 0; i < 4; i++ {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	// map channel slice  array :range
	//int data[3] ={1,2,3}
	var data [3]int = [3]int{10, 20, 30}
	//for i=0;i<lenght(data);i++{}
	for i := range data {
		//fmt.Println(i,data[i])
		//fmt.Println(i, s)
		fmt.Println(data[i])
	}

	data2 := [3]string{"aa", "bb", "cc"}
	var data2Copy [3]string
	for i, s := range data2 {
		//fmt.Println(&i,&s)
		data2Copy[i] = s
	}
	for _, s := range data2Copy {
		fmt.Println(s)
	}

	//goto break continue

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i > 2 {
			goto test

		}
	}
	test:
	fmt.Println("end")

	testswitch()
}

func test()  {

	fmt.Println(" ")
}

func testswitch()  {
	x:=1
	a,b,c :=1,2,3
	switch x {
	default:
		fmt.Println("all not equal")
	case a:
		fmt.Println("x=a")

	case b:

	case c,a:
		fmt.Println("x=c")

	}
}
