# Intersect

`Intersect(arr ICollection) ICollection`

获取两个集合的交集

```
func TestIntersect(t *testing.T) {
	oldColl := NewStrCollection([]string{"test1", "test2", "test3", "test4"})
	newColl := NewStrCollection([]string{"test3", "test4", "test5", "test6"})
	o := oldColl.Intersect(newColl)
	o.DD()

	ret, err := o.ToStrings()
	if err != nil {
		t.Fatal(err)
	}

	if len(ret) != 2 {
		t.Fatal("intersect len error ")
	}

	if ret[1] != "test4" {
		t.Fatal("intersect val error")
	}
}


/*
StrCollection(2):{
	0:	test3
	1:	test4
}
*/
```
