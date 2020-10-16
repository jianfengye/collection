# Sum

`Sum() IMix`

返回Collection中的元素的和

```go
intColl := NewIntCollection([]int{1, 2, 2, 3})
intColl.Sum().DD()
sum, err := intColl.Sum().ToInt()
if err != nil {
    t.Fatal(err)
}

if sum != 8 {
    t.Fatal("sum 错误")
}

/*
IMix(int): 8 
*/
```