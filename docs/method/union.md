# Union

`Union(arr ICollection) ICollection`

获取两个集合的并集

``` go
func TestUnion(t *testing.T) {
    oldColl := NewStrCollection([]string{"test1", "test2", "test3", "test4"})
    newColl := NewStrCollection([]string{"test3", "test4", "test5", "test6"})
    o := oldColl.Union(newColl)

	o.DD()
	ret, err := o.ToStrings()
	if err != nil {
		t.Fatal(err)
	}

	if len(ret) != 6 {
		t.Fatal("union len error ")
	}

	if ret[5] != "test6" {
		t.Fatal("union val error")
	}
}

/*
StrCollection(6):{
	0:	test1
	1:	test2
	2:	test3
	3:	test4
	4:	test5
	5:	test6
}
*/
```
