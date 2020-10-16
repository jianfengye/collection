# Join

`Join(split string, format ...func(item interface{}) string) string`

将Collection中的元素按照某种方式聚合成字符串。该函数接受一个或者两个参数，第一个参数是聚合字符串的分隔符号，第二个参数是聚合时候每个元素的格式化函数，如果没有设置第二个参数，则使用`fmt.Sprintf("%v")`来该格式化

```go
intColl := NewIntCollection([]int{2, 4, 3})
out := intColl.Join(",")
if out != "2,4,3" {
    t.Fatal("join错误")
}
out = intColl.Join(",", func(item interface{}) string {
    return fmt.Sprintf("'%d'", item.(int))
})
if out != "'2','4','3'" {
    t.Fatal("join 错误")
}
```