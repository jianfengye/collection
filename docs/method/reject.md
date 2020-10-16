# Reject

`Reject(func(item interface{}, key int) bool) ICollection`

将满足过滤条件的元素删除

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
retColl := intColl.Reject(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 3
})
if retColl.Count() != 3 {
    t.Fatal("Reject 重复错误")
}

retColl.DD()

/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```