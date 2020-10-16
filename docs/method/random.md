# Random

`Random() IMix`

随机获取Collection中的元素，随机数种子使用时间戳

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6})
out := intColl.Random()
out.DD()

_, err := out.ToInt()
if err != nil {
    t.Fatal(err.Error())
}

/*
IMix(int): 5 
*/
```