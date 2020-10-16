# Each

`Each(func(item interface{}, key int))`

对Collection中的每个函数都进行一次函数调用。传入的参数是回调函数。

如果希望在某次调用的时候中止，在此次调用的时候设置Collection的Error，就可以中止调用。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
sum := 0
intColl.Each(func(item interface{}, key int) {
    v := item.(int)
    sum = sum + v
})

if intColl.Err() != nil {
    t.Fatal(intColl.Err())
}

if sum != 10 {
    t.Fatal("Each 错误")
}

sum = 0
intColl.Each(func(item interface{}, key int) {
    v := item.(int)
    sum = sum + v
    if sum > 4 {
        intColl.SetErr(errors.New("stop the cycle"))
        return
    }
})

if sum != 6 {
    t.Fatal("Each 错误")
}

/*
PASS
*/
```