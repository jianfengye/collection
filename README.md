# collection

Collection包目标是用于替换golang原生的Slice，使用场景是在大量不追求性能，追求业务开发效能的场景。

Collection包目前支持的元素类型：int, int64, float32, float64, string, struct。

Collection的Error是随着Collection对象走，或者下沉到IMix中，所以可以放心在IArray和IMix进行链式调用，只需要最后进行一次错误检查即可。

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

具体方法示例：

```

===================
intColl := NewIntArray([]int{1,2})
intColl.DD()

===================
intColl := NewIntArray([]int{1,2})
intColl.Append(3)

===================
intColl := NewIntArray([]int{1,2})
foo := intColl.Index(1)

===================
intColl := NewIntArray([]int{1,2})
if intColl.IsEmpty() != false {
    t.Error("IsEmpty 错误")
}

===================
intColl := NewIntArray([]int{1,2})
if intColl.IsNotEmpty() != true {
    t.Error("IsNotEmpty 错误")
}

===================
intColl := NewIntArray([]int{1,2})
if intColl.Search(2) != 1 {
    t.Error("Search 错误")
}

===================
intColl = NewIntArray([]int{1,2, 3, 3, 2})
if intColl.Search(3) != 2 {
    t.Error("Search 重复错误")
}

===================
intColl := NewIntArray([]int{1,2, 3, 3, 2})
uniqColl := intColl.Unique()
if uniqColl.Count() != 3 {
    t.Error("Unique 重复错误")
}


===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5})
retColl := intColl.Reject(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 3
})
if retColl.Count() != 3 {
    t.Error("Reject 重复错误")
}


===================
intColl := NewIntArray([]int{1, 2, 3, 4, 3, 2})
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

===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5})
retColl := intColl.Slice(2)

retColl = intColl.Slice(2,2)

retColl = intColl.Slice(2, -1)

===================
intColl := NewIntArray([]int{1, 2 })

intColl2 := NewIntArray([]int{3, 4})

intColl.Merge(intColl2)

===================
intColl := NewIntArray([]int{1, 2 })

intColl2 := NewIntArray([]int{3, 4})

m, err := intColl.Combine(intColl2)

===================
intColl := NewIntArray([]int{1, 2 })

intColl2 := NewIntArray([]int{3, 4})

m, err := intColl.CrossJoin(intColl2)

===================
intColl := NewIntArray([]int{1, 2, 3, 4})
sum := 0
intColl.Each(func(item interface{}, key int) {
    v := item.(int)
    sum = sum + v
})
if sum != 10 {
    t.Error("Each 错误")
}
===================
intColl := NewIntArray([]int{1, 2, 3, 4})
newIntColl := intColl.Map(func(item interface{}, key int) IMix {
    v := item.(int)
    return NewMix(v * 2)
})

===================
intColl := NewIntArray([]int{1, 2, 3, 4})
sumMix := intColl.Reduce(func(carry IMix, item IMix) IMix {
    carryInt, _ := carry.ToInt()
    itemInt, _ := item.ToInt()
    return NewMix(carryInt + itemInt)
})


===================

intColl := NewIntArray([]int{1, 2, 3, 4})
if intColl.Every(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 1
}) != false {
    t.Error("Every错误")
}

if intColl.Every(func(item interface{}, key int) bool {
    i := item.(int)
    return i > 0
}) != true {
    t.Error("Every错误")
}

===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
ret := intColl.ForPage(1, 2)

===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
ret := intColl.Nth(4, 1)
===================
intColl := NewIntArray([]int{1, 2, 3})
ret := intColl.Pad(5, 0)
if ret.Err() != nil {
    t.Error(ret.Err().Error())
}

ret.DD()
if ret.Count() != 5 {
    t.Error("Pad 错误")
}

ret = intColl.Pad(-5, 0)
if ret.Err() != nil {
    t.Error(ret.Err().Error())
}
ret.DD()
if ret.Count() != 5 {
    t.Error("Pad 错误")
}
===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
pop := intColl.Pop()
in, err :=  pop.ToInt()
if err != nil {
    t.Error(err.Error())
}
if in != 6 {
    t.Error("Pop 错误")
}
intColl.DD()
if intColl.Count() != 5 {
    t.Error("Pop 后本体错误")
}
===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
intColl.Push(7)
===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
intColl.Prepend(0)
===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
out := intColl.Random()
===================
intColl := NewIntArray([]int{1, 2, 3, 4, 5, 6})
vs := intColl.Reverse()
===================
```
接口说明：
```

// IArray表示数组结构，有几种类型
type IArray interface {
	// IArray错误信息，链式调用的时候需要检查下这个error是否存在，每次调用之后都检查一下
	Err() error
	// 设置IArray的错误信息
	SetErr(error) IArray

	/*
	下面的方法对所有Array都生效
	 */
	// 复制一份当前相同类型的IArray结构，但是数据是空的
	NewEmpty(err ...error) IArray
	// 判断是否是空数组
	IsEmpty() bool
	// 判断是否是空数组
	IsNotEmpty() bool
	// 放入一个元素到数组中，对所有Array生效, 仅当item和array结构不一致的时候返回错误
	Append(item interface{}) IArray
	// 删除一个元素, 需要自类实现
	Remove(index int) IArray
	// 增加一个元素。
	Insert(index int, item interface{}) IArray
	// 查找数据中是否包含，-1不包含，>=0 返回数组中元素下标，对所有Array生效
	Search(item interface{}) int
	// 过滤数组中重复的元素，仅对基础Array生效
	Unique() IArray
	// 按照某个方法进行过滤, 保留符合的
	Filter(func(item interface{}, key int) bool) IArray
	// 按照某个方法进行过滤，去掉符合的
	Reject(func(item interface{}, key int) bool) IArray
	// 获取满足条件的第一个, 如果没有填写过滤条件，就获取所有的第一个
	First(...func(item interface{}, key int) bool) IMix
	// 获取满足条件的最后一个，如果没有填写过滤条件，就获取所有的最后一个
	Last(...func(item interface{}, key int) bool) IMix
	// 获取数组片段，对所有Array生效
	Slice(...int) IArray
	// 获取某个下标，对所有Array生效
	Index(i int) IMix
	// 获取数组长度，对所有Array生效
	Count() int
	// 将两个数组进行合并，参数的数据挂在当前数组中，返回当前数组，对所有Array生效
	Merge(arr IArray) IArray

	// 每个元素都调用一次的方法
	Each(func(item interface{}, key int))
	// 每个元素都调用一次的方法, 并组成一个新的元素
	Map(func(item interface{}, key int) IMix) IArray
	// 合并一些元素，并组成一个新的元素
	Reduce(func(carry IMix, item IMix) IMix) IMix
	// 判断每个对象是否都满足, 如果collection是空，返回true
	Every(func(item interface{}, key int) bool) bool
	// 按照分页进行返回
	ForPage(page int, perPage int) IArray
	// 获取第n位值组成数组
	Nth(n int, offset int) IArray
	// 组成的个数
	Pad(start int, def interface{}) IArray
	// 从队列右侧弹出结构
	Pop() IMix
	// 推入元素
	Push(item interface{}) IArray
	// 前面插入一个元素
	Prepend(item interface{}) IArray
	// 随机获取一个元素
	Random() IMix
	// 倒置
	Reverse() IArray
	// 随机乱置
	Shuffle() IArray
	// 打印出当前数组结构
	DD()
	// 打印出json
	ToJson() []byte
	/*
	下面的方法对ObjArray生效
	 */
	// 返回数组中对象的某个key组成的数组，仅对ObjectArray生效, key为对象属性名称，必须为public的属性
	Pluck(key string) IArray
	// 按照某个字段进行排序
	SortBy(key string) IArray
	// 按照某个字段进行排序,倒序
	SortByDesc(key string) IArray


	/*
	下面的方法对基础Array生效，但是ObjArray一旦设置了Compare函数也生效
	 */
	// 比较a和b，如果a>b, 返回1，如果a<b, 返回-1，如果a=b, 返回0
	// 设置比较函数，理论上所有Array都能设置比较函数，但是强烈不建议基础Array设置。
	SetCompare(func(a interface{}, b interface{}) int) IArray
	// 数组中最大的元素，仅对基础Array生效, 可以传递一个比较函数
	Max() IMix
	// 数组中最小的元素，仅对基础Array生效
	Min() IMix
	// 判断是否包含某个元素，（并不进行定位），对基础Array生效
	Contains(obj interface{}) bool
	// 根据key对象计数
	CountBy() IMap
	// 比较两个数组，获取第一个数组不在第二个数组中的元素，组成新数组
	Diff(arr IArray) IArray
	// 进行排序, 升序
	Sort() IArray
	// 进行排序，倒序
	SortDesc() IArray
	// 进行拼接
	Join(split string, format ...func(item interface{}) string) string

	/*
	下面的方法对基础Array生效
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
	// 转化为golang原生的字符数组，仅对StrArray生效
	ToStrings() ([]string, error)
	// 转化为golang原生的Int64数组，仅对Int64Array生效
	ToInt64s() ([]int64, error)
	// 转化为golang原生的Int数组，仅对IntArray生效
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
