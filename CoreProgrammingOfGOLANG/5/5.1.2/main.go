// package main

// import (
// 	"runtime"
// 	"time"
// )

// func main() {
// 	go func() {
// 		sum := 0
// 		for i := 0; i < 10000; i++ {
// 			sum += i
// 		}
// 		println(sum)
// 		time.Sleep(1 * time.Second)
// 	}()
// 	//NumGoroutine 可以返回当前程序的 goroutine 数目
// 	println("NumGoroutine:", runtime.NumGoroutine())
// 	getGOMAXPROCS()
// 	time.Sleep(5 * time.Second)
// }
// func getGOMAXPROCS() {
// 	//输入<=1时，返回当前GOMAXPROCS的值；反之设置GOMAXPROCS的值
// 	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
// }
package main

import "fmt"

type A struct {
	num int
}

func (a A) Test1() {
	fmt.Println(1)
}
func (a *A) Test2() {
	fmt.Println(2)
	a.num++
}
func main() {
	a := A{0}
	ap := &A{0}
	a.Test1()
	a.Test2()
	ap.Test1()
	ap.Test2()
}
