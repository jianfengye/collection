# Map

`Map(func(item interface{}, key int) interface{}) ICollection`

对Collection中的每个函数都进行一次函数调用，并将返回值组装成ICollection

这个回调函数形如： `func(item interface{}, key int) interface{}`

如果希望在某此调用的时候中止，就在此次调用的时候设置Collection的Error，就可以中止，且此次回调函数生成的结构不合并到最终生成的ICollection。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
newIntColl := intColl.Map(func(item interface{}, key int) interface{} {
    v := item.(int)
    return v * 2
})
newIntColl.DD()

if newIntColl.Count() != 4 {
    t.Fatal("Map错误")
}

newIntColl2 := intColl.Map(func(item interface{}, key int) interface{} {
    v := item.(int)

    if key > 2 {
        intColl.SetErr(errors.New("break"))
        return nil
    }

    return v * 2
})
newIntColl2.DD()

/*
IntCollection(4):{
	0:	2
	1:	4
	2:	6
	3:	8
}
IntCollection(3):{
	0:	2
	1:	4
	2:	6
}
*/
```