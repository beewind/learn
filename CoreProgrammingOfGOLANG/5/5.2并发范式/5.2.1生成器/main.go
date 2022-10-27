package main

import (
	"fmt"
	"math/rand"
)

//简单的带缓冲的生成器
func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
func GenerateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

//多个goroutine增强性生成器
func GenerateInt() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			select {
			//使用select扇入技术,增加生成的随机源
			case ch <- <-GenerateIntA():
			case ch <- <-GenerateIntB():
			}
		}
	}()
	return ch
}
func main() {
	ch := GenerateInt()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}
