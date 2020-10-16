# Unique

`Unique() ICollection`

将Collection中重复的元素进行合并，返回唯一的一个数组。

*注意* 此函数要求设置compare方法，基础元素数组（int, int64, float32, float64, string）可直接调用！

```go
intColl := NewIntCollection([]int{1,2, 3, 3, 2})
uniqColl := intColl.Unique()
if uniqColl.Count() != 3 {
    t.Fatal("Unique 重复错误")
}

uniqColl.DD()
/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```