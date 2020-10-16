# Index

`Index(i int) IMix`

Index获取元素中的第几个元素，下标从0开始，如果i超出了长度，则Collection记录错误。

```
intColl := NewIntCollection([]int{1,2})
foo := intColl.Index(1)
foo.DD()

/*
IMix(int): 2 
*/
```