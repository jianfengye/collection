# Contains

`Contains(obj interface{}) bool`

判断一个元素是否在Collection中，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
if intColl.Contains(1) != true {
    t.Fatal("contain 错误1")
}
if intColl.Contains(5) != false {
    t.Fatal("contain 错误2")
}
```