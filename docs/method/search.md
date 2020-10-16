# Search

`Search(item interface{}) int`

查找Collection中第一个匹配查询元素的下标，如果存在，返回下标；如果不存在，返回-1

*注意* 此函数要求设置compare方法，基础元素数组（int, int64, float32, float64, string）可直接调用！

```go
intColl := NewIntCollection([]int{1,2})
if intColl.Search(2) != 1 {
    t.Fatal("Search 错误")
}

intColl = NewIntCollection([]int{1,2, 3, 3, 2})
if intColl.Search(3) != 2 {
    t.Fatal("Search 重复错误")
}
```