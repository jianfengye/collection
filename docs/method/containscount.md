# ContainsCount
```go
ContainsCount(obj interface{}) int
```

判断包含某个元素的个数，返回0代表没有找到，返回正整数代表个数。必须设置compare函数

```go
func TestAbsCollection_ContainsCount(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 2, 3})
	count := intColl.ContainsCount(2)
	if count != 2 {
		t.Fatal(errors.New("contains count error"))
	}
}
```