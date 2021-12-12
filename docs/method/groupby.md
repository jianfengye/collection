# GroupBy 

`GroupBy(func(interface{}, int) interface{}) map[interface{}]ICollection`

GroupBy 类scala groupby 设计, 根据某个函数分组

```

func TestInt32Collection_GroupBy(t *testing.T) {
	objColl := NewInt32Collection([]int32{1, 1, 20, 4})
	groupBy := objColl.GroupBy(func(item interface{}, i2 int) interface{} {
		foo := item.(int32)
		return foo
	})

	for k, collection := range groupBy {
		t.Log(k)
		collection.DD()
	}
}

/*
=== RUN   TestInt32Collection_GroupBy
    /Users/yejianfeng/Documents/workspace/collection/int32_collection_test.go:97: 1
Int32Collection(2):{
	0:	1
	1:	1
}
    /Users/yejianfeng/Documents/workspace/collection/int32_collection_test.go:97: 20
Int32Collection(1):{
	0:	20
}
    /Users/yejianfeng/Documents/workspace/collection/int32_collection_test.go:97: 4
Int32Collection(1):{
	0:	4
}
*/
```