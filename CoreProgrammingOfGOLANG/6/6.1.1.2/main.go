package main

import (
	"reflect"
)

/*
	对于reflect.TypeOf(a):
	传入的实参两种类型:
		- 接口变量:
			- 绑定了具体类型实例 ==> 接口的动态类型
			- 未绑定具体类型实例 ==> 接口的静态类型
		- 具体类型变量 => 返回具体类型信息
*/
type INT int
type A struct {
	a int
}
type B struct {
	b string
}
type Ita interface {
	String() string
}

func (b B) String() string {
	return b.b
}

func main() {
	var a INT = 12
	var b int = 14

	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta == tb {
		println("ta==tb")
	} else {
		println("ta!=tb")
	}
	println(ta.Name()) //INT
	println(tb.Name()) //int

	println(ta.Kind().String()) // int
	println(tb.Kind().String()) // int

	s1 := A{1}
	s2 := B{"tata"}

	println(reflect.TypeOf(s1).Name()) //A
	println(reflect.TypeOf(s2).Name()) //B

	println(reflect.TypeOf(s1).Kind().String()) //int   (wrong)==> struct
	println(reflect.TypeOf(s2).Kind().String()) //string(wrong)==> struct

	ita := new(Ita)
	var itb Ita = s2

	println(reflect.TypeOf(ita).Name())                 //Ita	(wrong) => 空
	println(reflect.TypeOf(ita).Kind().String())        //interface(wrong)=>ptr
	println(reflect.TypeOf(ita).Elem().Name())          //Ita
	println(reflect.TypeOf(ita).Elem().Kind().String()) //interface

	println(reflect.TypeOf(itb).Name())          //B
	println(reflect.TypeOf(itb).Kind().String()) //struct

}
