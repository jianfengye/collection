# Reverse

`Reverse() ICollection`

将Collection数组进行转置

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
vs := intColl.Reverse()
vs.DD()

/*
IntCollection(6):{
	0:	6
	1:	5
	2:	4
	3:	3
	4:	2
	5:	1
}
*/
```