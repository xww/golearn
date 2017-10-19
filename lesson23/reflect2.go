package lesson23

import (
	"fmt"
	"reflect"
)

type X1 struct {
}

func (*X1) Test(x, y int) (int, error) {
	return x + y, fmt.Errorf("err:%d", x+y)

}

func (X1) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)

}
func main11() {
	var a X1
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Test")
	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}
	out := m.Call(in)
	for _, v := range out {
		fmt.Println(v)
	}

}
func main222() {
	var a X1
	v := reflect.ValueOf(&a)
	m := v.MethodByName("Format")
	out := m.Call([]reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf("x"),
		reflect.ValueOf(211),
	})
	fmt.Println(out)

	out = m.CallSlice([]reflect.Value{
		reflect.ValueOf("%s = %d"),
		reflect.ValueOf([]interface{}{"x",211}),
	})
	fmt.Println(out)
}
