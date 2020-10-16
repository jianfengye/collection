# Mode

`Mode() IMix`

获取Collection中的众数，如果有大于两个的众数，返回第一次出现的那个。

```go
intColl := NewIntCollection([]int{1, 2, 2, 3, 4, 5, 6})
mode, err := intColl.Mode().ToInt()
 if err != nil {
     t.Fatal(err.Error())
 }
 if mode != 2 {
     t.Fatal("Mode error")
 }
 
 intColl = NewIntCollection([]int{1, 2, 2, 3, 4, 4, 5, 6})
 
 mode, err = intColl.Mode().ToInt()
 if err != nil {
     t.Fatal(err.Error())
 }
 if mode != 2 {
     t.Fatal("Mode error")
 }
```