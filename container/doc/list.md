# container包

## list.go

list是一个双向链表

![](/home/zhuyao/workspace/golang/src/github.com/eugene/learn/container/doc/list.assets/image-20221026100818225.png)

### Element

```go
type Element struct{
    next,prev *Element//指向前,后的元素的指针
    list *List	// 元素所属的链表
    Value any	//实际存储的元素
}
```

##### func(*Element)Next

返回下一个element或nil

```go
func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
        // 当e被remove,e.list=nil	,e.list.root为哨兵节点
		return p
	}
	return nil
}
```

##### func (*Element)Prev

返回前一个element或nil

```go
func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}
```

### List

```go
type List struct{
    root Element//哨兵节点, only &root, root.prev, and root.next are used
    len int//链表长度
}
```

##### func (l *List) Init

```go
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}
```

##### 常用函数

```go
// 分配内存,初始化List
func New() *List { return new(List).Init() }

// 返回元素个数
func (l *List) Len() int { return l.len }

// 返回第一个元素
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// 返回最后一个元素
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}
```

##### func (l *List) insert(e, at *Element) *Element

其他PushXxx,InserXxx基于insert变形

```go
//在at后插入e
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}
```

##### func (l *List) insertValue(v any, at *Element) *Element 

##### func (l *List) remove(e *Element) 

```go
//移除e
func (l *List) remove(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.list = nil
	l.len--
}
```

##### func (l *List) move(e, at *Element)

其他MoveXxx基于此变形

```go
//把e move到at.next
func (l *List) move(e, at *Element) {
	if e == at {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}
```

