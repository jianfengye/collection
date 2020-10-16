# Nth

`Nth(n int, offset int) ICollection`

Nth(n int, offset int) 获取从offset偏移量开始的每第n个，偏移量offset的设置为第一个。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
ret := intColl.Nth(4, 1)
ret.DD()

if ret.Count() != 2 {
    t.Fatal("Nth 错误")
}

/*
IntCollection(2):{
	0:	2
	1:	6
}
*/
```