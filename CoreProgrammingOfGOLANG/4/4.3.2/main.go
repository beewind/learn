package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}
type St struct{}

func (St) Ping() {
	println("ping")
}
func (*St) Pang() {
	println("pang")
}
func main() {
	var st *St = nil
	var it Inter = st
	var itp *Inter = &it
	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)
	fmt.Printf("%p\n", itp)

	if it != nil {
		it.Pang()
		it.Ping() //这里出现了
	}
}
