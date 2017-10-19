package main

import (
	"github.com/xww/golearn/common"
	//"github.com/docker/docker/builder/dockerfile"

	"reflect"
)

func main() {
	common.SayHello()
	//dockerfile.BuildFromConfig()
	//govendor
	x := 100
	v := reflect.ValueOf(&x) //传入地址
	v.Elem().SetInt(200) //成功修改x值为200

	v.Interface()
}
