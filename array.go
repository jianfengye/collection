package collection

// IArray表示数组结构，有几种类型
type IArray interface {

	/*
	下面的方法对所有Array都生效
	 */
	// 复制一份当前相同类型的IArray结构，但是数据是空的
	NewEmpty() IArray
	// 判断是否是空数组
	IsEmpty() bool
	// 判断是否是空数组
	IsNotEmpty() bool
	// 放入一个元素到数组中，对所有Array生效
	Append(item interface{}) error
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
	Merge(arr IArray)
	// 合并，将当前数组每个元素作为key，传入参数每个元素作为value，组成一个map
	Combine(arr IArray) (IMap, error)
	// 进行笛卡尔乘积组成数组
	CrossJoin(arr IArray) (IMap, error)

	// 每个元素都调用一次的方法
	Each(func(item interface{}, key int))
	// 每个元素都调用一次的方法, 并组成一个新的元素
	Map(func(item interface{}, key int) IMix) IArray
	// 合并一些元素，并组成一个新的元素
	Reduce(func(carry IMix, item IMix) IMix) IMix
	// 判断每个对象是否都满足
	Every(func(item interface{}, key int) bool)
	// 按照分页进行返回
	ForPage(page int, perPage int) IArray
	// 获取第n位值组成数组
	Nth(n int) IArray
	// 组成的个数
	Pad(start int, def interface{}) IArray
	// 弹出结构
	Pop() IMix
	// 推入元素
	Push(item interface{})
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
	ToJson() string

	/*
	下面的方法对ObjArray生效
	 */
	// 返回数组中对象的某个key组成的数组，仅对ObjectArray生效, key为对象属性名称，必须为public的属性
	Column(string) (IArray, error)
	// 将数组中对象某个key作为map的key，整个对象作为value，作为map返回，如果key有重复会进行覆盖，仅对ObjectArray生效
	KeyBy(key string) (IMap, error)
	// 将对象的某个key作为map的key，对象的val作为map的value，作为map返回
	Pluck(val string, key string) (IMap, error)
	// 按照某个字段进行排序
	SortBy(key string) (IArray, error)
	// 按照某个字段进行排序,倒序
	SortByDesc(key string) (IArray, error)


	/*
	下面的方法对基础Array生效，但是ObjArray一旦设置了Compare函数也生效
	 */
	// 比较a和b，如果a>b, 返回1，如果a<b, 返回-1，如果a=b, 返回0
	// 设置比较函数，理论上所有Array都能设置比较函数，但是强烈不建议基础Array设置。
	SetCompare(func(a interface{}, b interface{}) int)
	// 数组中最大的元素，仅对基础Array生效, 可以传递一个比较函数
	Max() IMix
	// 数组中最小的元素，仅对基础Array生效
	Min() IMix
	// 判断是否包含某个元素，（并不进行定位），对基础Array生效
	Contains(obj interface{}) bool
	// 根据key对象计数
	CountBy() IMap
	// 比较两个数组，获取第一个数组不在第二个数组中的元素，组成新数组
	Diff(arr IArray) (IArray, error)
	// 进行排序, 升序
	Sort() IArray
	// 进行排序，倒序
	SortDesc() IArray

	/*
	下面方法对stringArray起作用
	 */
	Join(split string) string

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
	ToString() ([]string, error)
	// 转化为golang原生的Int64数组，仅对Int64Array生效
	ToInt64() ([]int64, error)
	// 转化为golang原生的Int数组，仅对IntArray生效
	ToInt() ([]int, error)
	// 转化为obj数组
	ToMix() []IMix
	// 转化为float64数组
	ToFloat64() ([]float64, error)
	// 转化为float32数组
	ToFloat32() ([]float32, error)
}