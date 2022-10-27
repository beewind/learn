package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label

			}
		}
		close(ch)
	}()
	return ch
}
func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	println("NumGoroutine=", runtime.NumGoroutine())

}
