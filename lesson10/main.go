package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "ab c\x62\u0041"
	fmt.Println(s)
	//nil ""
	var s2 string
	//fmt.Println(s2==nil)
	fmt.Println(s2 == "")

	s3 := `abc\ndef
	dfr` //yuuanshi esc
	fmt.Println(s3)

	s4 := "ab" + "cd"
	fmt.Println(s4 == "abcd")
	fmt.Println(s4 > "bbc")

	s5 := "abcdefg"
	//fmt.Printf("%c",s5[1])

	s6 := s5[1:3]
	fmt.Println(s6)
	s7 := s5[:4]
	fmt.Println(s7)

	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s5)))
	//fmt.Printf("%#v\n",(*reflect.StringHeader)(unsafe.Pointer(&s6)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s7)))

	for i := 0; i < len(s5); i++ {
		fmt.Printf("%c", s5[i])
	}

	for i, c := range s5 {
		fmt.Printf("%d, %c   ", i, c)
	}

	b := []byte(s5)
	fmt.Println(b)
	b[0] = 98
	s8 := string(b)
	fmt.Println(s8)

	//s9 := make([]byte,100)
	//bytes.NewBuffer()
}
