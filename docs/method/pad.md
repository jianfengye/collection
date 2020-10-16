# Pad

`Pad(start int, def interface{}) ICollection` 

填充Collection数组，如果第一个参数大于0，则代表往Collection的右边进行填充，如果第一个参数小于零，则代表往Collection的左边进行填充。

```go
intColl := NewIntCollection([]int{1, 2, 3})
ret := intColl.Pad(5, 0)
if ret.Err() != nil {
    t.Fatal(ret.Err().Error())
}

ret.DD()
if ret.Count() != 5 {
    t.Fatal("Pad 错误")
}

ret = intColl.Pad(-5, 0)
if ret.Err() != nil {
    t.Fatal(ret.Err().Error())
}
ret.DD()
if ret.Count() != 5 {
    t.Fatal("Pad 错误")
}

/*
IntCollection(5):{
	0:	1
	1:	2
	2:	3
	3:	0
	4:	0
}
IntCollection(5):{
	0:	0
	1:	0
	2:	1
	3:	2
	4:	3
}
*/
```