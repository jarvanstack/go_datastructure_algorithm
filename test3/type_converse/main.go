package main

import (
	"fmt"
	"reflect"
)
//检查是否是类型匹配，我们这里使用反射做的，不知奥性能如何，但是能做，十分简单.
func main() {
	fmt.Printf("CheckInt(1)=%#v\n", CheckInt(1))
	fmt.Printf("%v\n", 	CheckInt("hello"))
}
func CheckInt(in interface{})bool  {
	return reflect.TypeOf(in).Kind() == reflect.Int
}
