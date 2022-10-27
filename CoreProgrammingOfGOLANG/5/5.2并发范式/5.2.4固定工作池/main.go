package main

import (
	"fmt"
)

const (
	NUMBER = 10
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}
func main() {
	workers := NUMBER

	taskchan := make(chan task, 10)
	resultchan := make(chan int, 10)
	done := make(chan struct{}, 10)

	go InitTask(taskchan, resultchan, 100)
	DistributeTask(taskchan, workers, done)
	go CloseResult(done, resultchan, workers)
	sum := ProcessResult(resultchan)
	fmt.Println(sum)
}
func ProcessResult(r chan int) int {
	sum := 0
	for v := range r {
		sum += v
	}
	return sum
}
func InitTask(t chan task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for i := 0; i < qu; i++ {
		b := i*10 + 1
		e := (i + 1) * 10
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		t <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		t <- tsk
	}
	close(t)
}
func DistributeTask(taskchan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskchan, done)
	}
}
func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for t := range taskchan {
		t.do()
	}
	done <- struct{}{}
}
func CloseResult(done chan struct{}, resultchan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
	close(resultchan)
}
