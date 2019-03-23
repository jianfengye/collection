package collection


type IArray interface {
	// 返回当前空的结构的IArray
	NewEmptyIArray() IArray

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

	// 按照某个方法进行过滤
	Filter(func(obj interface{}, index int) bool) IArray
	// 获取满足条件的第一个
	First(func(obj interface{}, index int) bool) *Mix

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
	// 转化为obj数组
	//ToMix() []*Mix
}

// 这个是一个虚函数，能实现的都实现，不能实现的panic
type VArray struct {
	IArray
	Parent IArray
}

func (arr *VArray) Append(obj interface{}) {
	panic("Append: not Implement")
}

// Search find string in arr, -1 present not found, >=0 present index
func (arr *VArray) Search(obj interface{}) int {
	panic("Search: not Implement")
}

func (arr *VArray) Column(key string) IArray {
	panic("Column: not Implement")
}

func (arr *VArray) Unique() IArray {
	panic("Unique: not Implement")
}

func (arr *VArray) Max() *Mix {
	panic("Max: not Implement")
}

func (arr *VArray) Min() *Mix {
	panic("Min: not Implement")
}

func (arr *VArray) ToString() []string {
	panic("ToString: not Implement")
}

func (arr *VArray) ToInt64() []int64 {
	panic("ToInt64: not Implement")
}

func (arr *VArray) ToInt() []int{
	panic("ToInt: not Implement")
}

func (arr *VArray) KeyBy(key string) *Map {
	panic("KeyBy: not Implement")
}

func (arr *VArray) Slice(start, end int) IArray {
	panic("Slice: not Implement")
}

func (arr *VArray) Index(i int) *Mix {
	panic("Index: not Implement")
}

func (arr *VArray) Len() int {
	panic("Len: not Implement")
}

func (arr *VArray) Has(obj interface{}) bool {
	if arr.Parent.Search(obj) >= 0 {
		return true
	}
	return false
}

func (arr *VArray) Merge(bArr IArray) IArray {
	l := bArr.Len()
	for i := 0; i < l; i++{
		arr.Append(bArr.Index(i).ToInterface())
	}
	return arr
}

func (arr *VArray) NewEmptyIArray() IArray {
	panic("NewEmptyIArray: not Implement")
}

func (arr *VArray) Filter(f func(obj interface{}, index int) bool) IArray {
	ret := arr.Parent.NewEmptyIArray()
	l := arr.Parent.Len()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f(obj, i) == true {
			ret.Append(obj)
		}
	}
	return ret
}

func (arr *VArray) First(f func(obj interface{}, index int) bool) *Mix {
	l := arr.Parent.Len()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f(obj, i) == true {
			return NewMix(obj)
		}
	}
	return nil
}