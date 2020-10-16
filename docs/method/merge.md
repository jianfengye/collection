# Merge

`Merge(arr ICollection) ICollection`

将两个Collection的元素进行合并，这个函数会修改原Collection。

```go
intColl := NewIntCollection([]int{1, 2 })

intColl2 := NewIntCollection([]int{3, 4})

intColl.Merge(intColl2)

if intColl.Err() != nil {
    t.Fatal(intColl.Err())
}

if intColl.Count() != 4 {
    t.Fatal("Merge 错误")
}

intColl.DD()

/*
IntCollection(4):{
	0:	1
	1:	2
	2:	3
	3:	4
}
*/
```