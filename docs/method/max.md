# Max

`Max() IMix`

获取Collection中的最大元素，必须设置compare函数

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
max, err := intColl.Max().ToInt()
if err != nil {
    t.Fatal(err)
}

if max != 3 {
    t.Fatal("max错误")
}

```