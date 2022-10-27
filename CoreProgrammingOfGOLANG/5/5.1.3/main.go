package main

import "runtime"

func main() {
	example2()
}
func example1() {
	c := make(chan struct{})
	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		i <- struct{}{}
	}(c)
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	<-c
}
func example2() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(a chan struct{}, b chan int) {
		for i := 0; i < 10; i++ {
			b <- i
		}
		close(b)
		a <- struct{}{}
	}(c, ci)
	println("NumGoroutine=", runtime.NumGoroutine())
	<-c
	println("NumGoroutine=", runtime.NumGoroutine())
	for v := range ci {
		println(v)
	}
}
