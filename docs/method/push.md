# Push

`Push(item interface{}) ICollection`

往Collection的右侧推入一个元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
intColl.Push(7)
intColl.DD()
if intColl.Count() != 7 {
    t.Fatal("Push 后本体错误")
}

/*
IntCollection(7):{
	0:	1
	1:	2
	2:	3
	3:	4
	4:	5
	5:	6
	6:	7
}
*/
```