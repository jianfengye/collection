# Slice

`Slice(...int) ICollection`

获取Collection中的片段，可以有两个参数或者一个参数。

如果是两个参数，第一个参数代表开始下标，第二个参数代表结束下标，当第二个参数为-1时候，就代表到Collection结束。

如果是一个参数，则代表从这个开始下标一直获取到Collection结束的片段。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
retColl := intColl.Slice(2)
if retColl.Count() != 3 {
    t.Fatal("Slice 错误")
}

retColl.DD()

retColl = intColl.Slice(2,2)
if retColl.Count() != 2 {
    t.Fatal("Slice 两个参数错误")
}

retColl.DD()

retColl = intColl.Slice(2, -1)
if retColl.Count() != 3 {
    t.Fatal("Slice第二个参数为-1错误")
}

retColl.DD()

/*
IntCollection(3):{
	0:	3
	1:	4
	2:	5
}
IntCollection(2):{
	0:	3
	1:	4
}
IntCollection(3):{
	0:	3
	1:	4
	2:	5
}
*/

```