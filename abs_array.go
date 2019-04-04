package collection

import (
	"errors"
	"reflect"
)

// 这个是一个虚函数，能实现的都实现，不能实现的panic
type AbsArray struct {
	compare func(interface{}, interface{}) int // 比较函数

	IArray
	Parent IArray
}

/*
下面的几个函数必须要实现
 */
func (arr *AbsArray) NewEmpty() IArray {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.NewEmpty()
}

func (arr *AbsArray) Append(item interface{}) error {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Append(item)
}

func (arr *AbsArray) Index(i int) IMix {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Index(i)
}

func (arr *AbsArray) Count() int {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Count()
}

/*
下面这些函数是所有函数体都一样
 */
func (arr *AbsArray) IsEmpty() bool {
	return arr.Count() == 0
}

func (arr *AbsArray) IsNotEmpty() bool {
	return arr.Count() != 0
}

func (arr *AbsArray) Search(item interface{}) int {
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), item) == 0 {
			return i
		}
	}
	return -1
}

func (arr *AbsArray) Unique() IArray {
	newArr := arr.NewEmpty()
	for i := 0; i < arr.Count(); i++ {
		if newArr.Contains(arr.Index(i).ToInterface()) == false {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsArray) Reject(f func(item interface{}, key int) bool) IArray {
	newArr := arr.NewEmpty()
	for i := 0; i < arr.Count(); i++ {
		if f(arr.Index(i).ToInterface(), i) == false {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsArray) Last(fs ...func(item interface{}, key int) bool) IMix {
	if len(fs) > 1 {
		panic("Last 参数个数错误")
	}

	if len(fs) == 0 {
		return arr.Index(arr.Count() - 1)
	}

	newArr := arr.Filter(fs[0])
	return newArr.Last()
}

func (arr *AbsArray) Slice(ps ...int) IArray {
	if len(ps) > 2 || len(ps) == 0{
		panic("Slice params count error")
	}

	start := ps[0]
	count := arr.Count()
	if len(ps) == 2 && ps[1] != -1 {
		count = ps[1]
	}

	newArr := arr.NewEmpty()
	for i := 0; i < arr.Count(); i++ {
		if i >= start {
			newArr.Append(arr.Index(i).ToInterface())
			if newArr.Count() >= count {
				break
			}
		}
	}
	return newArr
}

func (arr *AbsArray) Merge(arr2 IArray) {
	for i := 0; i < arr2.Count(); i++ {
		arr.Append(arr2.Index(i).ToInterface())
	}
}

func (arr *AbsArray) Combine(arr2 IArray) (IMap, error) {
	if arr.Count() == 0 {
		return nil, errors.New("combine: count can not be zero")
	}

	if arr.Count() != arr2.Count() {
		return nil, errors.New("combine: count not match")
	}

	ret := NewEmptyMap(reflect.TypeOf(arr.Index(0).ToInterface()), reflect.TypeOf(arr2.Index(0).ToInterface()))
	for i := 0; i < arr.Count(); i++ {
		key := arr.Index(i).ToInterface()
		val := arr2.Index(i).ToInterface()
		ret.Set(key, val)
	}
	return ret, nil
}

func (arr *AbsArray) CrossJoin(arr2 IArray) IMap {
	panic("CrossJoin: not Implement")
}

func (arr *AbsArray) Each(f func(item interface{}, key int)) {
	panic("Each: not Implement")
}

func (arr *AbsArray) Map(func(item interface{}, key int)) IArray {
	panic("Map: not Implement")
}

func (arr *AbsArray) Reduce(func(carry IMix, item IMix) IMix) IMix {
	panic("Reduce: not Implement")
}

func (arr *AbsArray) Every(func(item interface{}, key int) bool) {
	panic("Every: not Implement")
}

func (arr *AbsArray) ForPage(page int, perPage int) IArray {
	panic("ForPage: not Implement")
}

func (arr *AbsArray) Nth(n int) IArray {
	panic("Nth: not Implement")
}

func (arr *AbsArray) Pad(start int, def interface{}) IArray {
	panic("Pad: not Implement")
}

func (arr *AbsArray) Pop() IMix {
	panic("Pop: not Implement")
}

func (arr *AbsArray) Push(item interface{}) {
	panic("Push: not Implement")
}

func (arr *AbsArray) Prepend(item interface{}) IArray {
	panic("Prepend: not Implement")
}

func (arr *AbsArray) Random() IMix {
	panic("Random: not Implement")
}

func (arr *AbsArray) Reverse() IArray {
	panic("Reverse: not Implement")
}

func (arr *AbsArray) Shuffle() IArray {
	panic("Shuffle: not Implement")
}

func (arr *AbsArray) DD()  {
	panic("DD: not Implement")
}

func (arr *AbsArray) ToJson() string {
	panic("ToJson: not Implement")
}

func (arr *AbsArray) Column(string) (IArray, error) {
	panic("Column: not Implement")
}

func (arr *AbsArray) KeyBy(key string) (IMap, error) {
	panic("KeyBy: not Implement")
}

func (arr *AbsArray) Pluck(val string, key string) (IMap, error) {
	panic("Pluck: not Implement")
}

func (arr *AbsArray) SortBy(key string) (IArray, error) {
	panic("SortBy: not Implement")
}

func (arr *AbsArray) SortByDesc(key string) (IArray, error) {
	panic("Reverse: not Implement")
}

func (arr *AbsArray) SetCompare(func(a interface{}, b interface{}) int) {
	panic("Reverse: not Implement")
}

func (arr *AbsArray) Max() IMix {
	panic("Max: not Implement")
}

func (arr *AbsArray) Min() IMix {
	panic("Min: not Implement")
}

func (arr *AbsArray) Contains(obj interface{}) bool {
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), obj) == 0 {
			return true
		}
	}
	return false
}

func (arr *AbsArray) CountBy() IMap {
	panic("CountBy: not Implement")
}

func (arr *AbsArray) Diff(arr2 IArray) (IArray, error) {
	panic("Diff: not Implement")
}

func (arr *AbsArray) Sort() IArray {
	panic("Sort: not Implement")
}

func (arr *AbsArray) SortDesc() IArray {
	panic("ToString: not Implement")
}

func (arr *AbsArray) Join(split string) string {
	panic("Join: not Implement")
}

func (arr *AbsArray) Avg() IMix{
	panic("Avg: not Implement")
}

func (arr *AbsArray) Median() IMix {
	panic("Median: not Implement")
}

func (arr *AbsArray) Mode() IMix {
	panic("Mode: not Implement")
}

func (arr *AbsArray) Sum() IMix {
	panic("Sum: not Implement")
}

func (arr *AbsArray) Filter(f func(obj interface{}, index int) bool) IArray {
	ret := arr.Parent.NewEmpty()
	l := arr.Parent.Count()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f(obj, i) == true {
			ret.Append(obj)
		}
	}
	return ret
}

func (arr *AbsArray) First(f ...func(obj interface{}, index int) bool) IMix {
	if len(f) == 0 {
		return arr.Parent.Index(0)
	}
	fun := f[0]

	l := arr.Parent.Count()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if fun(obj, i) == true {
			return arr.Parent.Index(i)
		}
	}
	return nil
}

func (arr *AbsArray) ToString() ([]string, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToInt64() ([]int64, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToInt() ([]int, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToMix() []IMix {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToFloat64() ([]float64, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToFloat32() ([]float32, error) {
	panic("Sum: not Implement")
}
