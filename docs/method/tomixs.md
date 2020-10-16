# ToMixs

`ToMixs() ([]IMix, error)`

将Collection变化为Mix数组，如果Collection内的元素类型不符合，或者Collection有错误，则返回错误

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
arr, err := intColl.ToMixs()
if err != nil {
    t.Fatal(err)
}
if len(arr) != 4 {
    t.Fatal(errors.New("ToInts error"))
}
```