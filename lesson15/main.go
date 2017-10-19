/*
package main

import "fmt"

type Stu struct {
	name string
	age int
}

type N int

func main() {
	var s Stu
	fmt.Println(s.setAge(11))

	var a N = 23
	fmt.Println(a.toString())

}

func (s *Stu)setAge(a int )int  {
	s.age = a
	return s.age

}

func (n *N)toString()string  {
	return  fmt.Sprintf("%s", n)
}
*/

package main

import (
	"sync"
)

type Stu struct {
	//name string
	age int
	people
}

func (s Stu) value() {
	s.age = 11

}
func (s *Stu) pointer() {
	s.age = 11
}

type people struct {
	name string
}

func (p *people) getname() string {
	return p.name
}

type data struct {
	sync.Mutex
	buf [1024]byte
}

func main() {
	/*var s1 Stu
	s1.age=22
	s1.value()
	fmt.Println(s1.age)
	s1.pointer()

	var s2 *Stu
	s2 = &Stu{
		age:22,
		people:people{name:"xww"},
	}
	s2.pointer()
	fmt.Println(s2.age)
	s2.value()

	fmt.Println(s2.getname())

	d := data{}
	d.Lock()

	//ye wu chu li

	d.Unlock()*/

	var n N
	var t Ner = &n
	t.a()

}

type Ner interface {
	a()
	b(int)
	c(string) string
}
type N int

func (N) a()               {}
func (*N) b(int)           {}
func (*N) c(string) string { return "" }
