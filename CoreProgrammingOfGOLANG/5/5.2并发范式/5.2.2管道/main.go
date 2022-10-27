package main

import (
	"fmt"
	"time"
)

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}
func main() {
	// in := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		in <- i
	// 	}
	// 	close(in)
	// }()
	// out := chain(chain(chain(in)))
	// for v := range out {
	// 	fmt.Println(v)
	// }
	c := make(chan int, 5)
	cs := make(chan struct{})
	go roundChan(c, c, cs)
	time.Sleep(5 * time.Second)
	close(cs)
}
func roundChan(in, out chan int, done chan struct{}) {
Lable:
	for {
		select {
		case v := <-in:
			fmt.Println(v)
			out <- 1 + v
		case <-done:
			break Lable
		default:
			in <- 0
		}
		time.Sleep(1 * time.Second)
	}
	close(in)
}
