# Prepend

`Prepend(item interface{}) ICollection`

往Collection左侧加入元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
intColl.Prepend(0)
if intColl.Err() != nil {
    t.Fatal(intColl.Err().Error())
}

intColl.DD()
if intColl.Count() != 7 {
    t.Fatal("Prepend错误")
}

/*
IntCollection(7):{
	0:	0
	1:	1
	2:	2
	3:	3
	4:	4
	5:	5
	6:	6
}
*/
```