package main

import "fmt"

type data struct {
	x int
	s string
}

func main() {
	/*
		& += &= != () >>= ... . :

		1. * / % >> << &
		2. + - | ^
		3.&&
		4.||
		5;
	*/

	const (
		read byte = 1 << iota
		write
		exec
		freeze
	)

	a := read | write | freeze
	b := read | write | exec
	c := a & b
	fmt.Printf("%04b & %04b = %04b\n", a, b, c)

	/*d :=1
	e := (++d) > 1
	if e{
	fmt.Println(d)
	}*/
	d := 1
	p := &d
	*p++
	fmt.Println(*p)

	x := 10
	var p1 *int = &x
	*p += 20
	fmt.Println(p1, *p1)

	p2 := &x

	fmt.Println(p2)
	p2 = nil
	fmt.Println(p2)

	var a1 data = data{1, "abc"}
	fmt.Println(a1)
	b1 := data{
		1,
		"abc",
	}
	c1 := &data{
		x: 1,
		s: "abc",
	}
	fmt.Println(b1, c1)

}
