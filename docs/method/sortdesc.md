# SortDesc

`SortDesc() ICollection`

将Collection中的元素按照降序排列输出，必须设置compare函数

```go
intColl := NewIntCollection([]int{2, 4, 3})
intColl2 := intColl.SortDesc()
if intColl2.Err() != nil {
    t.Fatal(intColl2.Err())
}
intColl2.DD()

/*
IntCollection(3):{
	0:	4
	1:	3
	2:	2
}
*/
```