package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"test/code"
)

type person struct {
	name string
	age  int
}

/*
1.新建person{name string,age int}链表l
2.插入2个值
3.遍历链表l
4.新建l2,向l2插入l,遍历l
*/
func list_code() {
	l := list.New()
	l.PushBack(person{"eugene1", 1})
	l.PushBack(person{"eugene2", 2})

	l2 := list.New()
	l2.PushBackList(l)
	l2.PushBack(person{"eugene3", 3})
	for ele := l2.Front(); ele != nil; ele = ele.Next() {
		p := ele.Value.(person)
		fmt.Printf("%#v\n", p)
	}
	for ele := l.Front(); ele != nil; ele = ele.Next() {
		p := ele.Value.(person)
		fmt.Printf("%#v\n", p)
	}
}

/*
0.分别创建len=3,len=9的ring,len=3的链对一个节点赋值
1.link two links
2.link one link
3.use ring.Do
*/
type INT struct {
	x int
}

func ring_code() {
	r := ring.New(3)
	r.Value = &INT{1}
	n := r.Next()
	s := r.Value.(*INT)

	fmt.Printf("r.Value: %v\n", s)
	fmt.Println(r == n)
	r.Do(changeTo2)

	fmt.Println("双链link---------")
	newr := ring.New(9)
	next_newr := newr.Next()
	r.Link(newr)
	r.Do(print)
	fmt.Println("----------")
	next_newr.Do(print)

	fmt.Println("单链link---------")
	next_r := r.Next()
	r.Link(newr.Move(5))

	r.Do(print)
	fmt.Println("---------")
	next_r.Do(print)
}
func print(a any) {
	fmt.Println(a)
}
func changeTo2(a interface{}) {
	if a == nil {
		return
	}
	if a.(*INT).x == 1 {
		a.(*INT).x = 2
		fmt.Println(a)
	}
}

func main() {
	code.Heap_code1()
}
