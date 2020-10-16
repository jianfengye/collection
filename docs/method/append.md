# Append

`Append(item interface{}) ICollection`

Append挂载一个元素到当前Collection，如果挂载的元素类型不一致，则会在Collection中产生Error

```
intColl := NewIntCollection([]int{1,2})
intColl.Append(3)
if intColl.Err() == nil {
    intColl.DD()
}

/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```