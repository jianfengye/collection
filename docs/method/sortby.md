# SortBy

`SortBy(key string) ICollection`

根据对象数组中的某个元素进行Collection升序排列。这个元素必须是Public元素

注：这个函数只对ObjCollection生效。这个对象数组的某个元素必须是基础类型。

```go
type Foo struct {
	A string
	B int
}

func TestObjCollection_SortBy(t *testing.T) {
	a1 := Foo{A: "a1", B: 3}
	a2 := Foo{A: "a2", B: 2}

	objColl := NewObjCollection([]Foo{a1, a2})

	newObjColl := objColl.SortBy("B")

	newObjColl.DD()

	obj, err := newObjColl.Index(0).ToInterface()
	if err != nil {
		t.Fatal(err)
	}

	foo := obj.(Foo)
	if foo.B != 2 {
		t.Fatal("SortBy error")
	}
}

/*
ObjCollection(2)(collection.Foo):{
	0:	{A:a2 B:2}
	1:	{A:a1 B:3}
}
*/
```