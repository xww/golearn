package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func test(f func(s string) (int, error)) int {
	a, _ := f("abc")
	return a
}

func test2(s string) (int, error) {
	//fmt.Println(s)
	return 2, errors.New("error")
}

func main() {
	/*f, err := os.Open("filename")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	//defer f.Close()

	var c []byte
	b := bytes.NewBuffer(c)
	content,_ := f.Read(b.Bytes())
	fmt.Println(content)

	//f.Close()*/

	x, y := 1, 2
	defer func() {
		if err := recover();err !=nil{
			fmt.Println(err)
			//^666

		}
	}()

	fmt.Println(x, y+1)

	fmt.Println(testdefer())

	//panic recover try catch

	panic("database is not connencted")
	fmt.Println("test panic")

}

func testdefer()  ( int){
	a :=100
	defer func() {
		fmt.Println("defer",a)
		a+=100
	}()
	//return a
	temp:=a
	return temp
}
