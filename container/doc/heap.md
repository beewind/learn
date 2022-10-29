heap 需要实现Sort的接口,以及Push方法,Pop方法

## 结构体

```go
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}
```
