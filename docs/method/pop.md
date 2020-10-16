# Pop

`Pop() IMix`

从Collection右侧弹出一个元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
pop := intColl.Pop()
in, err := pop.ToInt()
if err != nil {
    t.Fatal(err.Error())
}
if in != 6 {
    t.Fatal("Pop 错误")
}
intColl.DD()
if intColl.Count() != 5 {
    t.Fatal("Pop 后本体错误")
}

/*
IntCollection(5):{
	0:	1
	1:	2
	2:	3
	3:	4
	4:	5
}
*/
```