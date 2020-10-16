# DD 

`DD()`

DD方法按照友好的格式展示Collection

```
a1 := Foo{A: "a1"}
a2 := Foo{A: "a2"}

objColl := NewObjCollection([]Foo{a1, a2})
objColl.DD()

/*
ObjCollection(2)(collection.Foo):{
	0:	{A:a1}
	1:	{A:a2}
}
*/

intColl := NewIntCollection([]int{1,2})
intColl.DD()

/*
IntCollection(2):{
	0:	1
	1:	2
}
*/
```