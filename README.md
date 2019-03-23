# collection

这个包目标是用于替换golang原生的Slice和Map，使用场景是在大量不追求性能，追求业务开发效能的场景。

```

type IArray interface {
	// 放入一个元素到数组中，对所有Array生效
	Append(obj interface{})

	// 查找数据中是否包含，-1不包含，>=0 返回数组中元素下标，对所有Array生效
	Search(obj interface{}) int
	// 返回数组中对象的某个key组成的数组，仅对ObjectArray生效
	Column(key string) IArray
	// 过滤数组中重复的元素，仅对基础Array生效
	Unique() IArray

	// 将数组中对象某个key作为map的key，整个对象作为value，作为map返回，如果key有重复会进行覆盖，仅对ObjectArray生效
	KeyBy(key string) *Map

	// 数组中最大的元素，仅对基础Array生效
	Max() *Mix
	// 数组中最小的元素，仅对基础Array生效
	Min() *Mix

	// 获取数组片段，对所有Array生效
	Slice(start, end int) IArray
	// 获取某个下标，对所有Array生效
	Index(i int) *Mix
	// 获取数组长度，对所有Array生效
	Len() int
	// 判断是否包含某个元素，（并不进行定位），对基础Array生效
	Has(obj interface{}) bool
	// 将两个数组进行合并，参数的数据挂在当前数组中，返回当前数组，对所有Array生效
	Merge(arr IArray) IArray

	// 转化为golang原生的字符数组，仅对StrArray生效
	ToString() []string
	// 转化为golang原生的Int64数组，仅对Int64Array生效
	ToInt64() []int64
	// 转化为golang原生的Int数组，仅对IntArray生效
	ToInt() []int
}
```
```
type IMap interface {

	// 设置一个Map的key和value，如果key存在，则覆盖
	Set(key interface{}, value interface{})
	// 删除一个Map的key
	Remove(key interface{})
	// 根据key获取一个Map的value
	Get(key interface{}) *Mix
	// 获取一个Map的长度
	Len() int

	// 获取Map的所有key组成的集合
	Keys() IArray
	// 获取Map的所有value组成的集合
	Values() IArray
}
```
