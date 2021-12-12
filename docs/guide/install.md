# 安装和版本


Collection包目标是用于替换golang原生的Slice，使用场景是在大量不追求极致性能，追求业务开发效能的场景。

| 版本 | 说明 |
| ------| ------ |
| v1.4.0 |  增加三种新类型 uint32, uint, uint64, 增加GroupBy 和 Split 方法 |
| v1.3.0 |  增加文档说明 |
| 1.2.0 |  增加对象指针数组，增加测试覆盖率, 增加ToInterfaces方法 |
| 1.1.2 |  增加一些空数组的判断，解决一些issue |
| 1.1.1 |  对collection包进行了json解析和反解析的支持，对mix类型支持了SetField和RemoveFields的类型设置 |
| 1.1.0 |  增加了对int32的支持，增加了延迟加载，增加了Copy函数，增加了compare从ICollection传递到IMix，使用快排加速了Sort方法 |
| 1.0.1 |  第一次发布 |

`go get github.com/jianfengye/collection@v1.4.0`

Collection包目前支持的元素类型：int32, int, int64, uint32, uint, uint64, float32, float64, string, struct, struct_point。

使用下列几个方法进行初始化Collection:

```go
NewIntCollection(objs []int) *IntCollection

NewInt64Collection(objs []int64) *Int64Collection

NewInt32Collection(objs []int32) *Int32Collection

NewUIntCollection(objs []uint) *UIntCollection

NewUInt64Collection(objs []uint64) *UInt64Collection

NewUInt32Collection(objs []uint32) *UInt32Collection

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