ring是一个环形链表

## 结构体

```go
type Ring struct {
  next, prev *Ring
  Value    any // for use by client; untouched by this library
}
```

## 方法

Init方法创建一个空环,指向自身,所以空环只有一个nil元素?

```go
func (r *Ring) init() *Ring {
  r.next = r
  r.prev = r
  return r
}
```

New方法创建n个节点的环

```go
func New(n int) *Ring {
  if n <= 0 {
    return nil
  }
  r := new(Ring)
  p := r
  for i := 1; i < n; i++ {
    p.next = &Ring{prev: p}
    p = p.next
  }
  p.next = r
  r.prev = p
  return r
}
```



Next,Prev返回下一个和前一个节点

```go
// Next returns the next ring element. r must not be empty.

func (r *Ring) Next() *Ring {
  if r.next == nil {
    return r.init()
  }
  return r.next
}
```

```go
// Prev returns the previous ring element. r must not be empty.

func (r *Ring) Prev() *Ring {
  if r.next == nil {
    return r.init()
  }
  return r.prev
}
```



Move向next方向为正,prev为负,移动n步返回终点的节点

```go
func (r *Ring) Move(n int) *Ring {
  if r.next == nil {
    return r.init()
  }
  switch {
  case n < 0:
    for ; n < 0; n++ {
      r = r.prev
    }
  case n > 0:
    for ; n > 0; n-- {
      r = r.next
    }
  }
  return r
}
```

link连接两个ring,当一个ring自link时会产生两个ring

```go
func (r *Ring) Link(s *Ring) *Ring {
  n := r.Next()
  if s != nil {
    p := s.Prev()
    // Note: Cannot use multiple assignment because
    // evaluation order of LHS is not specified.
    r.next = s
    s.prev = r
    n.prev = p
    p.next = n
  }
  return n
}
```

移除r next的n%r.Len()个节点

```go
func (r *Ring) Unlink(n int) *Ring {
  if n <= 0 {
    return nil
  }
  return r.Link(r.Move(n + 1))

}
```

返回ring内元素的个数

```go
func (r *Ring) Len() int {
  n := 0
  if r != nil {
     n = 1
     for p := r.Next(); p != r; p = p.next {
     n++
    }
  }
  return n
}
```

Do使用f对每个元素处理,Do是值拷贝,注意作用范围

```go
func (r *Ring) Do(f func(any)) {
  if r != nil {
    f(r.Value)
    for p := r.Next(); p != r; p = p.next {
      f(p.Value)
    }
  }
}
```



