# Last

`Last(...func(item interface{}, key int) bool) IMix`

获取该Collection中满足过滤的最后一个元素，如果没有填写过滤条件，默认返回最后一个元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 3, 2})
last, err := intColl.Last().ToInt()
if err != nil {
    t.Fatal("last get error")
}
if last != 2 {
    t.Fatal("last 获取错误")
}

last, err = intColl.Last(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 2
}).ToInt()

if err != nil {
    t.Fatal("last get error")
}
if last != 3 {
    t.Fatal("last 获取错误")
}
```