# Copy

Copy方法根据当前的数组，创造出一个同类型的数组，有相同的元素

```
func TestAbsCollection_Copy(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2})
	intColl2 := intColl.Copy()
	intColl2.DD()
	if intColl2.Count() != 2 {
		t.Fatal("Copy失败")
	}
	if reflect.TypeOf(intColl2) != reflect.TypeOf(intColl) {
		t.Fatal("Copy类型失败")
	}
}
```