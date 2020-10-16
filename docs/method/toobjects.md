# ToObjs

`ToObjs(interface{}) error`

这个方法是1.2.2才加入的。

这个方法只能由ObjCollection 或者 ObjPointCollection 调用，否则会返回error

将 collection 中的数组复原成为 Slice。参数需要传递指向目标 Slice 的指针。

``` golang
func TestObjCollection_ToObjs(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 3}

	bArr := []Foo{}
	objColl := NewObjCollection([]Foo{a1, a2, a3})
	err := objColl.ToObjs(&bArr)
	if err != nil {
		t.Fatal(err)
	}
	if len(bArr) != 3 {
		t.Fatal("toObjs error len")
	}
	if bArr[1].A != "a2" {
		t.Fatal("toObjs error copy")
	}
}

```