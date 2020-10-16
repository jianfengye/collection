# Reduce

`Reduce(func(carry IMix, item IMix) IMix) IMix`

对Collection中的所有元素进行聚合计算。

如果希望在某次调用的时候中止，在此次调用的时候设置Collection的Error，就可以中止调用。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4})
sumMix := intColl.Reduce(func(carry IMix, item IMix) IMix {
    carryInt, _ := carry.ToInt()
    itemInt, _ := item.ToInt()
    return NewMix(carryInt + itemInt)
})

sumMix.DD()

sum, err := sumMix.ToInt()
if err != nil {
    t.Fatal(err.Error())
}
if sum != 10 {
    t.Fatal("Reduce计算错误")
}

/*
IMix(int): 10 
*/
```