# ForPage

`ForPage(page int, perPage int) ICollection`

将Collection函数进行分页，按照每页第二个参数的个数，获取第一个参数的页数数据。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
ret := intColl.ForPage(1, 2)
ret.DD()

if ret.Count() != 2 {
    t.Fatal("For page错误")
}

/*
IntCollection(2):{
	0:	3
	1:	4
}
*/
```