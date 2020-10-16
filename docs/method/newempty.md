# NewEmpty

`NewEmpty(err ...error) ICollection`

NewEmpty方法根据当前的数组，创造出一个同类型的数组，但长度为0

```
intColl := NewIntCollection([]int{1,2})
intColl2 := intColl.NewEmpty()
intColl2.DD()

/*
IntCollection(0):{
}
*/
```