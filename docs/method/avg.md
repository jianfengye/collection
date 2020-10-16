# Avg

`Avg() IMix`

返回Collection的数值平均数，这里会进行类型降级，int,int64,float64的数值平均数都是返回float64类型。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
mode, err := intColl.Avg().ToFloat64()
if err != nil {
    t.Fatal(err.Error())
}
if mode != 2.0 {
    t.Fatal("Avg error")
}
```