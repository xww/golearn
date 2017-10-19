package main

import "fmt"

func main() {
	//数组类型：元素类型，长度
	var a1 [3]*int
	var a2 [3]*int
	a1 = a2
	fmt.Println(a1==a2)

	b := [4]int{1, 2, 3}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 5: 100}
	fmt.Println(c)

	type stu struct {
		name string
		age  int
	}

	/*s1 := stu{
		name:"xww",
		age:11,
	}
	s2 := stu{
		name:"xww2",
		age:12,
	}*/
	d := [...]stu{
		{name: "xww1", age: 11},
		{name: "xww2", age: 120},
	}
	fmt.Println(d)

	e := [2][4]int{
		{1, 2, 3},
	}
	fmt.Println(e)

	fmt.Println(len(d))
	fmt.Println(len(e))
	fmt.Println(cap(e))

	// 指针数组   数组指针 字符串指针
	x,y := 10,20
	z := [...]*int{&x,&y}
	p:=&z
	fmt.Println(*p)
	fmt.Println(&x,&y)

	f :=[2]int{11,22}
	var g [2]int
	g=f
	fmt.Printf("x: %p, %v \n",&f,f)
	fmt.Printf("x: %p, %v \n",&g,g)

	test(&f)
	fmt.Println(f)




}

func test(x *[2]int)  {
	fmt.Printf("x: %p, %v \n",&x,x)
	x[0]=1111
}
