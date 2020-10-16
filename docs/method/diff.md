# Diff

`Diff(arr ICollection) ICollection`

获取前一个Collection不在后一个Collection中的元素，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl2 := NewIntCollection([]int{2, 3, 4})

diff := intColl.Diff(intColl2)
diff.DD()
if diff.Count() != 1 {
    t.Fatal("diff 错误")
}

/*
IntCollection(1):{
	0:	1
}
*/
```