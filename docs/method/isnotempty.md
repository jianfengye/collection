# IsNotEmpty

`IsNotEmpty() bool`

判断一个Collection是否为空，为空返回false，否则返回true
```go
intColl := NewIntCollection([]int{1,2})
println(intColl.IsNotEmpty()) // true