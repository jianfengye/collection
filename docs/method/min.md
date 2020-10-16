# Min

`Min() IMix`

获取Collection中的最小元素，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
min, err := intColl.Min().ToInt()
if err != nil {
    t.Fatal(err)
}

if min != 1 {
    t.Fatal("min错误")
}

```