# Shuffle

`Shuffle() ICollection`

将Collection中的元素进行乱序排列，随机数种子使用时间戳

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
newColl := intColl.Shuffle()
newColl.DD()
if newColl.Err() != nil {
    t.Fatal(newColl.Err())
}

/*
IntCollection(4):{
	0:	1
	1:	3
	2:	2
	3:	2
}
*/
```