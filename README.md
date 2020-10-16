# collection

Collection包目标是用于替换golang原生的Slice，使用场景是在大量不追求极致性能，追求业务开发效能的场景。


Collection的使用手册线上地址：http://collection.funaio.cn

| 版本 | 说明 |
| ------| ------ |
| 1.3.0 |  增加文档说明 |
| 1.2.0 |  增加对象指针数组，增加测试覆盖率, 增加ToInterfaces方法 |
| 1.1.2 |  增加一些空数组的判断，解决一些issue |
| 1.1.1 |  对collection包进行了json解析和反解析的支持，对mix类型支持了SetField和RemoveFields的类型设置 |
| 1.1.0 |  增加了对int32的支持，增加了延迟加载，增加了Copy函数，增加了compare从ICollection传递到IMix，使用快排加速了Sort方法 |
| 1.0.1 |  第一次发布 |

`go get github.com/jianfengye/collection`

创建collection库的说明文章见：[一个让业务开发效率提高10倍的golang库](https://www.cnblogs.com/yjf512/p/10818089.html)

Collection包目前支持的元素类型：int32, int, int64, float32, float64, string, struct, struct_point。

使用下列几个方法进行初始化Collection:

```go
NewIntCollection(objs []int) *IntCollection

NewInt64Collection(objs []int64) *Int64Collection

NewInt32Collection(objs []int32) *Int32Collection

NewFloat64Collection(objs []float64) *Float64Collection

NewFloat32Collection(objs []float32) *Float32Collection

NewStrCollection(objs []string) *StrCollection

NewObjCollection(objs interface{}) *ObjCollection

NewObjPointCollection(objs interface{}) *ObjPointCollection
```

Collection的Error是随着Collection对象走，或者下沉到IMix中，所以可以放心在ICollection和IMix进行链式调用，只需要最后进行一次错误检查即可。

```
ret, err := objColl.Map(func(item interface{}, key int) IMix {
    foo := item.(Foo)
    return NewMix(foo.A)
}).Reduce(func(carry IMix, item IMix) IMix {
    ret, _ := carry.ToString()
    join, _ := item.ToString()
    return NewMix(ret + join)
}).ToString()
if err != nil {
    ...
}
```






License
------------
`collection` is licensed under [Apache License](LICENSE).

## Contributors

This project exists thanks to all the people who contribute. [[Contribute](CONTRIBUTING.md)].
<a href="https://github.com/jianfengye/collection/graphs/contributors"><img src="https://opencollective.com/collection/contributors.svg?width=890&button=false" /></a>


## Backers

Thank you to all our backers! 🙏 [[Become a backer](https://opencollective.com/collection#backer)]

<a href="https://opencollective.com/collection#backers" target="_blank"><img src="https://opencollective.com/collection/backers.svg?width=890"></a>


## Sponsors

Support this project by becoming a sponsor. Your logo will show up here with a link to your website. [[Become a sponsor](https://opencollective.com/collection#sponsor)]

<a href="https://opencollective.com/collection/sponsor/0/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/1/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/2/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/3/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/4/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/5/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/6/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/7/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/8/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/collection/sponsor/9/website" target="_blank"><img src="https://opencollective.com/collection/sponsor/9/avatar.svg"></a>


