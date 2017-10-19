package main

import (
	"fmt"
	"unsafe"
)

type user struct {
	name string
	age  int
}

type attr struct {
	owner int
	perm  int
}

type file struct {
	name string
	 attr
}

func main() {
	//fmt.Println(unsafe.Sizeof(node))
	u := user{
		name: "xww",
		age:  11,
	}
	fmt.Println(u)
	u2 := user{"xww2", 12}
	fmt.Println(u2)

	u3 := struct {
		name string
		age  int
	}{
		name: "xww",
		age:  11,
	}
	fmt.Println(u3)

	/*f := file{
		name:"aa.txt",
	}
	f.attr.owner=1
	f.attr.perm=2*/

	var a struct{}
	var b [100]struct{}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))

	f2:=file{
		name:"aa.txt",
		attr:attr{
			owner:1,
			perm:2,

		},
	}
	fmt.Println(f2)

}
