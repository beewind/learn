package main

import (
	"fmt"
	"sync"
)

//工作任务
type task struct {
	begin  int
	end    int
	result chan<- int
}

//任务执行:计算begin到end的和
//执行结果写入chan result
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}
func main() {
	//创建任务
	taskchan := make(chan task, 10)

	//创建结果通道
	resultchan := make(chan int, 10)

	//wait用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	//初始化task的goroutine,计算100个自然数之和
	go InitTask(taskchan, resultchan, 100)

	//每个任务启动一个goroutine进行处理
	go DistributeTask(taskchan, wait, resultchan)

	//通过结果通道汇总和
	sum := ProcessResult(resultchan)

	fmt.Println(sum)
}
func InitTask(taskchan chan task, r chan int, p int) {
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
		taskchan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tsk
	}
	close(taskchan)
}
func DistributeTask(taskchan chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskchan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(result)
}
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}
func ProcessResult(r chan int) int {
	sum := 0
	for v := range r {
		sum += v
	}
	return sum
}
