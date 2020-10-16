# Every

`Every(func(item interface{}, key int) bool) bool`

判断Collection中的每个元素是否都符合某个条件，只有当每个元素都符合条件，才整体返回true，否则返回false。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
if intColl.Every(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 1
}) != false {
    t.Fatal("Every错误")
}

if intColl.Every(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 0
}) != true {
    t.Fatal("Every错误")
}
```