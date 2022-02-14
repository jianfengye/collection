package collection

// ICollection 表示数组结构，有几种类型
type ICollection interface {
	// Err ICollection错误信息，链式调用的时候需要检查下这个error是否存在，每次调用之后都检查一下
	Err() error
	// SetErr 设置ICollection的错误信息
	SetErr(error) ICollection

	/*
		下面的方法对所有Collection都生效
	*/
	// NewEmpty 复制一份当前相同类型的ICollection结构，但是数据是空的
	NewEmpty(err ...error) ICollection
	// IsEmpty 判断是否是空数组
	IsEmpty() bool
	// IsNotEmpty 判断是否是空数组
	IsNotEmpty() bool
	// Append 放入一个元素到数组中，对所有Collection生效, 仅当item和Collection结构不一致的时候返回错误
	Append(item interface{}) ICollection
	// Remove 删除一个元素, 需要自类实现
	Remove(index int) ICollection
	// Insert 增加一个元素。
	Insert(index int, item interface{}) ICollection
	// Search 查找数据中是否包含，-1不包含，>=0 返回数组中元素下标，对所有Collection生效
	Search(item interface{}) int
	// Unique 过滤数组中重复的元素，仅对基础Collection生效
	Unique() ICollection
	// Filter 按照某个方法进行过滤, 保留符合的
	Filter(func(item interface{}, key int) bool) ICollection
	// Reject 按照某个方法进行过滤，去掉符合的
	Reject(func(item interface{}, key int) bool) ICollection
	// First 获取满足条件的第一个, 如果没有填写过滤条件，就获取所有的第一个
	First(...func(item interface{}, key int) bool) IMix
	// Last 获取满足条件的最后一个，如果没有填写过滤条件，就获取所有的最后一个
	Last(...func(item interface{}, key int) bool) IMix
	// Slice 获取数组片段，对所有Collection生效
	Slice(...int) ICollection
	// Index 获取某个下标，对所有Collection生效
	Index(i int) IMix
	// SetIndex 设置数组的下标为某个值
	SetIndex(i int, val interface{}) ICollection
	// Copy 复制当前数组
	Copy() ICollection
	// Count 获取数组长度，对所有Collection生效
	Count() int
	// Merge 将两个数组进行合并，参数的数据挂在当前数组中，返回当前数组，对所有Collection生效
	Merge(arr ICollection) ICollection

	// Each 每个元素都调用一次的方法
	Each(func(item interface{}, key int))
	// Map 每个元素都调用一次的方法, 并组成一个新的元素
	Map(func(item interface{}, key int) interface{}) ICollection
	// Reduce 合并一些元素，并组成一个新的元素
	Reduce(func(carry IMix, item IMix) IMix) IMix
	// Every 判断每个对象是否都满足, 如果collection是空，返回true
	Every(func(item interface{}, key int) bool) bool
	// ForPage 按照分页进行返回
	ForPage(page int, perPage int) ICollection
	// Nth 获取从索引offset开始为0，每n位值组成数组
	Nth(n int, offset int) ICollection
	// Pad 将数组填充到count个数，只能数值型生效
	Pad(count int, def interface{}) ICollection
	// Pop 从队列右侧弹出结构
	Pop() IMix
	// Push 推入元素
	Push(item interface{}) ICollection
	// Prepend 前面插入一个元素
	Prepend(item interface{}) ICollection
	// Random 随机获取一个元素
	Random() IMix
	// Reverse 倒置
	Reverse() ICollection
	// Shuffle 随机乱置
	Shuffle() ICollection
	// GroupBy 类scala groupby 设计, 根据某个函数分组
	GroupBy(func(interface{}, int) interface{}) map[interface{}]ICollection
	// Split 按照size个数进行分组
	Split(size int) []ICollection
	// DD 打印出当前数组结构
	DD()

	/*
		下面的方法对ObjCollection生效
	*/
	// Pluck 返回数组中对象的某个key组成的数组，仅对ObjectCollection生效, key为对象属性名称，必须为public的属性
	Pluck(key string) ICollection
	// SortBy 按照某个字段进行排序
	SortBy(key string) ICollection
	// SortByDesc 按照某个字段进行排序,倒序
	SortByDesc(key string) ICollection

	/*
		下面的方法对基础Collection生效，但是ObjCollection一旦设置了Compare函数也生效
	*/
	// SetCompare 比较a和b，如果a>b, 返回1，如果a<b, 返回-1，如果a=b, 返回0
	// 设置比较函数，理论上所有Collection都能设置比较函数，但是强烈不建议基础Collection设置
	SetCompare(func(a interface{}, b interface{}) int) ICollection
	// GetCompare 获取比较函数
	GetCompare() func(a interface{}, b interface{}) int
	// Max 数组中最大的元素，仅对基础Collection生效, 可以传递一个比较函数
	Max() IMix
	// Min 数组中最小的元素，仅对基础Collection生效
	Min() IMix
	// Contains 判断是否包含某个元素，（并不进行定位），对基础Collection生效
	Contains(obj interface{}) bool
	// ContainsCount 判断包含某个元素的个数，返回0代表没有找到，返回正整数代表个数。必须设置compare函数
	ContainsCount(obj interface{}) int
	// Diff 比较两个数组，获取第一个数组不在第二个数组中的元素，组成新数组
	Diff(arr ICollection) ICollection
	// Sort 进行排序, 升序
	Sort() ICollection
	// SortDesc 进行排序，倒序
	SortDesc() ICollection
	// Join 进行拼接
	Join(split string, format ...func(item interface{}) string) string
	// Union 两个集合的并集
	Union(arr ICollection) ICollection
	// Intersect 两个集合的交集
	Intersect(arr ICollection) ICollection

	/*
		下面的方法对基础Collection生效
	*/
	// Avg 获取平均值
	Avg() IMix
	// Median 获取中位值
	Median() IMix
	// Mode 获取Mode值
	Mode() IMix
	// Sum 获取sum值
	Sum() IMix

	/*
		下面的方法对根据不同的对象，进行不同的调用转换
	*/
	// ToStrings 转化为golang原生的字符数组，仅对StrCollection生效
	ToStrings() ([]string, error)
	// ToInt64s 转化为golang原生的Int64数组，仅对Int64Collection生效
	ToInt64s() ([]int64, error)
	// ToInt32s 转化为golang原生的Int32数组，仅对Int32Collection生效
	ToInt32s() ([]int32, error)
	// ToInts 转化为golang原生的Int数组，仅对IntCollection生效
	ToInts() ([]int, error)
	// ToUInt64s 转化为golang原生的UInt64数组，仅对UInt64Collection生效
	ToUInt64s() ([]uint64, error)
	// ToUInt32s 转化为golang原生的UInt32数组，仅对UInt32Collection生效
	ToUInt32s() ([]uint32, error)
	// ToUInts 转化为golang原生的UInt数组，仅对UIntCollection生效
	ToUInts() ([]uint, error)
	// ToMixs 转化为obj数组
	ToMixs() ([]IMix, error)
	// ToFloat64s 转化为float64数组
	ToFloat64s() ([]float64, error)
	// ToFloat32s 转化为float32数组
	ToFloat32s() ([]float32, error)
	// ToInterfaces 转化为interface{} 数组
	ToInterfaces() ([]interface{}, error)
	// ToObjs 转化为objs{}数组
	ToObjs(interface{}) error

	// ToJson 转换为Json
	ToJson() ([]byte, error)
	// FromJson 从json数组转换
	FromJson([]byte) error
}
