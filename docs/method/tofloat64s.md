# ToFloat64s

`ToFloat64s() ([]float64, error)`

将Collection变化为float64数组，如果Collection内的元素类型不符合，或者Collection有错误，则返回错误。

```go
arr := NewFloat64Collection([]float64{1.0 ,2.0,3.0,4.0,5.0})

arr.DD()

max, err := arr.Max().ToFloat64()
if err != nil {
    t.Fatal(err)
}

if max != 5 {
    t.Fatal(errors.New("max error"))
}


arr2 := arr.Filter(func(obj interface{}, index int) bool {
    val := obj.(float64)
    if val > 2.0 {
        return true
    }
    return false
})
if arr2.Count() != 3 {
    t.Fatal(errors.New("filter error"))
}

out, err := arr2.ToFloat64s()
if err != nil || len(out) != 3 {
    t.Fatal(errors.New("to float64s error"))
}

```