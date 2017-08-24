package main

import "fmt"

const x, y int = 1, 0x23
const s = "hello"
const c = "aa"

func main() {
	const d = 123
	//fmt.Println(d)

	const y = 123
	{
		const x = "aavc"
		//fmt.Println(x)
	}

	const (
		x1, y1      = 1, 2
		b      byte = byte(x)
		n           = uint8(12)
	)

	const (
		x2 = iota
		y2
		z2 = 100
		d1
		e  = iota
		f
	)
	fmt.Println(x2, y2, z2, d1, e, f)

	j := 2
	j = 3
	fmt.Println(&j)

	const k = 2
	b1 := k + 2

	fmt.Println(b1)

}
