# collection

Collection包目标是用于替换golang原生的Slice，使用场景是在大量不追求性能，追求业务开发效能的场景。

Collection包目前支持的元素类型：int, int64, float32, float64, string, struct。

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

支持的方法有:

[DD](#DD)

[NewEmpty](#NewEmpty)

[Append](#Append)

[Index](#Index) 

[IsEmpty](#IsEmpty)

[IsNotEmpty](#IsNotEmpty)

[Search](#Search)

[Unique](#Unique)

[Reject](#Reject)

[Last](#Last)



### DD 

DD方法按照友好的格式展示Collection

```
a1 := Foo{A: "a1"}
a2 := Foo{A: "a2"}

objColl := NewObjCollection([]Foo{a1, a2})
objColl.DD()

/*
ObjCollection(2)(collection.Foo):{
	0:	{A:a1}
	1:	{A:a2}
}
*/

intColl := NewIntCollection([]int{1,2})
intColl.DD()

/*
IntCollection(2):{
	0:	1
	1:	2
}
*/
```

### NewEmpty

NewEmpty方法根据当前的数组，创造出一个同类型的数组，但长度为0

```
intColl := NewIntCollection([]int{1,2})
intColl2 := intColl.NewEmpty()
intColl2.DD()

/*
IntCollection(0):{
}
*/
```

### Append

Append挂载一个元素到当前Collection，如果挂载的元素类型不一致，则会在Collection中产生Error

```
intColl := NewIntCollection([]int{1,2})
intColl.Append(3)
if intColl.Err() == nil {
    intColl.DD()
}

/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```

### Index

Index获取元素中的第几个元素，下标从0开始，如果i超出了长度，则Collection记录错误。

```
intColl := NewIntCollection([]int{1,2})
foo := intColl.Index(1)
foo.DD()

/*
IMix(int): 2 
*/
```

### IsEmpty

判断一个Collection是否为空，为空返回true, 否则返回false

```go
intColl := NewIntCollection([]int{1,2})
println(intColl.IsEmpty())  // false
```

### IsNotEmpty

判断一个Collection是否为空，为空返回false，否则返回true
```go
intColl := NewIntCollection([]int{1,2})
println(intColl.IsNotEmpty()) // true
```

### Search

查找Collection中第一个匹配查询元素的下标，如果存在，返回下标；如果不存在，返回-1

*注意* 此函数要求设置compare方法，基础元素数组（int, int64, float32, float64, string）可直接调用！

```go
intColl := NewIntCollection([]int{1,2})
if intColl.Search(2) != 1 {
    t.Error("Search 错误")
}

intColl = NewIntCollection([]int{1,2, 3, 3, 2})
if intColl.Search(3) != 2 {
    t.Error("Search 重复错误")
}
```

### Unique

将Collection中重复的元素进行合并，返回唯一的一个数组。

*注意* 此函数要求设置compare方法，基础元素数组（int, int64, float32, float64, string）可直接调用！

```go
intColl := NewIntCollection([]int{1,2, 3, 3, 2})
uniqColl := intColl.Unique()
if uniqColl.Count() != 3 {
    t.Error("Unique 重复错误")
}

uniqColl.DD()
/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```

### Reject

将满足过滤条件的元素删除

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
retColl := intColl.Reject(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 3
})
if retColl.Count() != 3 {
    t.Error("Reject 重复错误")
}

retColl.DD()

/*
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
*/
```

### Last

获取该Collection中满足过滤的最后一个元素，如果没有填写过滤条件，默认返回最后一个元素

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 3, 2})
last, err := intColl.Last().ToInt()
if err != nil {
    t.Error("last get error")
}
if last != 2 {
    t.Error("last 获取错误")
}

last, err = intColl.Last(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 2
}).ToInt()

if err != nil {
    t.Error("last get error")
}
if last != 3 {
    t.Error("last 获取错误")
}
```

### Slice

获取Collection中的片段，可以有两个参数或者一个参数。

如果是两个参数，第一个参数代表开始下标，第二个参数代表结束下标，当第二个参数为-1时候，就代表到Collection结束。

如果是一个参数，则代表从这个开始下标一直获取到Collection结束的片段。

```go
intColl := NewIntCollection([]int{1, 2, 3, 4, 5})
retColl := intColl.Slice(2)
if retColl.Count() != 3 {
    t.Error("Slice 错误")
}

retColl.DD()

retColl = intColl.Slice(2,2)
if retColl.Count() != 2 {
    t.Error("Slice 两个参数错误")
}

retColl.DD()

retColl = intColl.Slice(2, -1)
if retColl.Count() != 3 {
    t.Error("Slice第二个参数为-1错误")
}

retColl.DD()

/*
IntCollection(3):{
	0:	3
	1:	4
	2:	5
}
IntCollection(2):{
	0:	3
	1:	4
}
IntCollection(3):{
	0:	3
	1:	4
	2:	5
}
*/

```


接口说明：
```

// ICollection表示数组结构，有几种类型
type ICollection interface {
	// ICollection错误信息，链式调用的时候需要检查下这个error是否存在，每次调用之后都检查一下
	Err() error
	// 设置ICollection的错误信息
	SetErr(error) ICollection

	/*
	下面的方法对所有Collection都生效
	 */
	// 复制一份当前相同类型的ICollection结构，但是数据是空的
	NewEmpty(err ...error) ICollection
	// 判断是否是空数组
	IsEmpty() bool
	// 判断是否是空数组
	IsNotEmpty() bool
	// 放入一个元素到数组中，对所有Collection生效, 仅当item和Collection结构不一致的时候返回错误
	Append(item interface{}) ICollection
	// 删除一个元素, 需要自类实现
	Remove(index int) ICollection
	// 增加一个元素。
	Insert(index int, item interface{}) ICollection
	// 查找数据中是否包含，-1不包含，>=0 返回数组中元素下标，对所有Collection生效
	Search(item interface{}) int
	// 过滤数组中重复的元素，仅对基础Collection生效
	Unique() ICollection
	// 按照某个方法进行过滤, 保留符合的
	Filter(func(item interface{}, key int) bool) ICollection
	// 按照某个方法进行过滤，去掉符合的
	Reject(func(item interface{}, key int) bool) ICollection
	// 获取满足条件的第一个, 如果没有填写过滤条件，就获取所有的第一个
	First(...func(item interface{}, key int) bool) IMix
	// 获取满足条件的最后一个，如果没有填写过滤条件，就获取所有的最后一个
	Last(...func(item interface{}, key int) bool) IMix
	// 获取数组片段，对所有Collection生效
	Slice(...int) ICollection
	// 获取某个下标，对所有Collection生效
	Index(i int) IMix
	// 获取数组长度，对所有Collection生效
	Count() int
	// 将两个数组进行合并，参数的数据挂在当前数组中，返回当前数组，对所有Collection生效
	Merge(arr ICollection) ICollection

	// 每个元素都调用一次的方法
	Each(func(item interface{}, key int))
	// 每个元素都调用一次的方法, 并组成一个新的元素
	Map(func(item interface{}, key int) IMix) ICollection
	// 合并一些元素，并组成一个新的元素
	Reduce(func(carry IMix, item IMix) IMix) IMix
	// 判断每个对象是否都满足, 如果collection是空，返回true
	Every(func(item interface{}, key int) bool) bool
	// 按照分页进行返回
	ForPage(page int, perPage int) ICollection
	// 获取第n位值组成数组
	Nth(n int, offset int) ICollection
	// 组成的个数
	Pad(start int, def interface{}) ICollection
	// 从队列右侧弹出结构
	Pop() IMix
	// 推入元素
	Push(item interface{}) ICollection
	// 前面插入一个元素
	Prepend(item interface{}) ICollection
	// 随机获取一个元素
	Random() IMix
	// 倒置
	Reverse() ICollection
	// 随机乱置
	Shuffle() ICollection
	// 打印出当前数组结构
	DD()
	// 打印出json
	ToJson() []byte
	/*
	下面的方法对ObjCollection生效
	 */
	// 返回数组中对象的某个key组成的数组，仅对ObjectCollection生效, key为对象属性名称，必须为public的属性
	Pluck(key string) ICollection
	// 按照某个字段进行排序
	SortBy(key string) ICollection
	// 按照某个字段进行排序,倒序
	SortByDesc(key string) ICollection


	/*
	下面的方法对基础Collection生效，但是ObjCollection一旦设置了Compare函数也生效
	 */
	// 比较a和b，如果a>b, 返回1，如果a<b, 返回-1，如果a=b, 返回0
	// 设置比较函数，理论上所有Collection都能设置比较函数，但是强烈不建议基础Collection设置。
	SetCompare(func(a interface{}, b interface{}) int) ICollection
	// 数组中最大的元素，仅对基础Collection生效, 可以传递一个比较函数
	Max() IMix
	// 数组中最小的元素，仅对基础Collection生效
	Min() IMix
	// 判断是否包含某个元素，（并不进行定位），对基础Collection生效
	Contains(obj interface{}) bool
	// 根据key对象计数
	CountBy() IMap
	// 比较两个数组，获取第一个数组不在第二个数组中的元素，组成新数组
	Diff(arr ICollection) ICollection
	// 进行排序, 升序
	Sort() ICollection
	// 进行排序，倒序
	SortDesc() ICollection
	// 进行拼接
	Join(split string, format ...func(item interface{}) string) string

	/*
	下面的方法对基础Collection生效
	 */
	// 获取平均值
	Avg() IMix
	// 获取中位值
	Median() IMix
	// 获取Mode值
	Mode() IMix
	// 获取sum值
	Sum() IMix


	/*
	下面的方法对根据不同的对象，进行不同的调用转换
	 */
	// 转化为golang原生的字符数组，仅对StrCollection生效
	ToStrings() ([]string, error)
	// 转化为golang原生的Int64数组，仅对Int64Collection生效
	ToInt64s() ([]int64, error)
	// 转化为golang原生的Int数组，仅对IntCollection生效
	ToInts() ([]int, error)
	// 转化为obj数组
	ToMixs() ([]IMix, error)
	// 转化为float64数组
	ToFloat64s() ([]float64, error)
	// 转化为float32数组
	ToFloat32s() ([]float32, error)
}
```

```
type IMix interface {
	Err() error
	SetErr(err error) IMix

	Equal(n IMix) bool // 两个IMix结构是否相同
	Type() reflect.Type // 获取类型

	Add(mix IMix) (IMix, error) // 加法操作
	Sub(mix IMix) (IMix, error) // 减法操作
	Div(n int) (IMix, error) // 除法操作
	Multi(n int) (IMix, error) // 乘法操作

	ToString() (string, error)
	ToInt64() (int64, error)
	ToInt() (int, error)
	ToFloat64() (float64, error)
	ToFloat32() (float32, error)
	ToInterface() interface{} // 所有函数可用

	Format() string // 打印成string
	DD()
}
```

### Index

获取Collection中某个位置的元素，位置下标从0开始

```
intColl := NewIntCollection([]int{1,2})
foo := intColl.Index(1)
foo.DD()

/*
IMix(int): 2 
*/
```

