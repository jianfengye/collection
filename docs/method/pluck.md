# Pluck

`Pluck(key string) ICollection`

将对象数组中的某个元素提取出来组成一个新的Collection。这个元素必须是Public元素

注：这个函数只对ObjCollection生效。

```go
type Foo struct {
	A string
}

func TestObjCollection_Pluck(t *testing.T) {
	a1 := Foo{A: "a1"}
	a2 := Foo{A: "a2"}

	objColl := NewObjCollection([]Foo{a1, a2})

	objColl.Pluck("A").DD()
}

/*
StrCollection(2):{
	0:	a1
	1:	a2
}
*/
```