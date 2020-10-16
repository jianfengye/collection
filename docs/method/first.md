# First

`First(...func(item interface{}, key int) bool) IMix`

获取符合过滤条件的第一个元素，如果没有填写过滤函数，返回第一个元素。

注：只能传递0个或者1个过滤函数，如果传递超过1个过滤函数，只有第一个过滤函数起作用

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl.First(func(obj interface{}, index int) bool {
    val := obj.(int)
    if val > 2 {
        return true
    }
    return false
}).DD()

/*
IMix(int): 3 
*/

```

```
func TestIntCollection_Filter(t *testing.T) {
	intColl := NewIntCollection([]int{1,2,3})
	a, err := intColl.First().ToInt()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a, 1) {
		t.Fatal("filter error")
	}
}
```