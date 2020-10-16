# Median

`Median() IMix`

获取Collection的中位数，如果Collection个数是单数，返回排序后中间的元素，如果Collection的个数是双数，返回排序后中间两个元素的算数平均数。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
median, err := intColl.Median().ToFloat64()
if err != nil {
    t.Fatal(err)
}

if median != 2.0 {
    t.Fatal("Median 错误" + fmt.Sprintf("%v", median))
}
```