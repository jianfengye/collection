# Filter

`Filter(func(item interface{}, key int) bool) ICollection`

根据过滤函数获取Collection过滤后的元素。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl.Filter(func(obj interface{}, index int) bool {
    val := obj.(int)
    if val == 2 {
        return true
    }
    return false
}).DD()

/*
IntCollection(2):{
	0:	2
	1:	2
}
*/
```