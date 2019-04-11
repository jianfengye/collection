package collection

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
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

func (arr *AbsArray) Insert(index int, obj interface{}) error {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Insert(index, obj)
}


func (arr *AbsArray) Remove(index int) error {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Remove(index)
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

func (arr *AbsArray) DD()  {
	panic("DD: not Implement")
}

func (arr *AbsArray) ToJson() []byte {
	panic("ToJson: not Implement")
}

/*
下面这些函数是所有函数体都一样
 */

func (arr *AbsArray) Append(item interface{}) error {
	return arr.Parent.Insert(arr.Count(), item)
}

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

func (arr *AbsArray) CrossJoin(arr2 IArray) (IMap, error) {
	if arr.Count() == 0 || arr2.Count() == 0 {
		return nil, errors.New("CrossJoin: count can not be zero")
	}

	ret := NewEmptyMap(reflect.TypeOf(arr.Index(0).ToInterface()), reflect.TypeOf(arr2.Index(0).ToInterface()))
	for i := 0; i < arr.Count(); i++ {
		for j := 0; j < arr2.Count(); j++ {
			key := arr.Index(i).ToInterface()
			val := arr2.Index(j).ToInterface()
			ret.Set(key, val)
		}
	}
	return ret, nil
}

func (arr *AbsArray) Each(f func(item interface{}, key int)) {
	for i := 0; i < arr.Count(); i++ {
		f(arr.Index(i).ToInterface(), i)
	}
}

func newMixArray(mix IMix) IArray {
	switch mix.Type().Kind() {
	case reflect.String:
		return NewStrArray([]string{})
	case reflect.Int:
		return NewIntArray([]int{})
	}
	return nil
}

func (arr *AbsArray) Map(f func(item interface{}, key int) IMix) IArray {
	// call first f for map type
	if arr.Count() == 0 {
		return nil
	}

	first := f(arr.Index(0).ToInterface(), 0)
	ret := newMixArray(first)
	ret.Append(first.ToInterface())
	for i := 1; i < arr.Count(); i++ {
		ret.Append(f(arr.Index(i).ToInterface(), 0).ToInterface())
	}
	return ret
}

func (arr *AbsArray) Reduce(f func(carry IMix, item IMix) IMix) IMix {
	if arr.Count() == 0 {
		return nil
	}

	if arr.Count() == 1 {
		return NewMix(arr.Index(0).ToInterface())
	}

	carry := f(NewMix(arr.Index(0).ToInterface()), NewMix(arr.Index(1).ToInterface()))

	for i := 2; i < arr.Count(); i++ {
		carry = f(carry, NewMix(arr.Index(i).ToInterface()))
	}
	return carry
}

func (arr *AbsArray) Every(f func(item interface{}, key int) bool) bool {
	if arr.Count() == 0 {
		return true
	}

	for i := 0; i < arr.Count(); i++ {
		if f(arr.Index(i).ToInterface(), i) == false {
			return false
		}
	}
	return true
}

func (arr *AbsArray) ForPage(page int, perPage int) IArray {
	start := page * perPage
	return arr.Slice(start, perPage)
}

func (arr *AbsArray) Nth(n int, offset int) IArray {
	newArr := arr.NewEmpty()
	for i := 0; i < arr.Count(); i++ {
		if (i - offset) % n == 0 {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsArray) Pad(start int, def interface{}) (IArray, error) {
	newArr := arr.NewEmpty()

	if start > 0 {
		if start <= arr.Count() {
			return arr.Slice(0, start), nil
		}

		for i:=0; i < arr.Count(); i++ {
			newArr.Append(arr.Index(i).ToInterface())
		}
		for i := arr.Count(); i < start; i++ {
			err := newArr.Append(def)
			if err != nil {
				return nil, err
			}
		}

		return newArr, nil
	}

	if start < 0 {
		if start >= -arr.Count() {
			return arr.Slice(arr.Count() + start, arr.Count()), nil
		}

		for i := start; i < -arr.Count(); i++ {
			err := newArr.Append(def)
			if err != nil {
				return nil, err
			}
		}

		for i := 0; i < arr.Count(); i++ {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}

	return newArr, nil
}

func (arr *AbsArray) Pop() IMix {
	ret := arr.Index(arr.Count() - 1)
	arr.Remove(arr.Count() - 1)

	return ret
}

func (arr *AbsArray) Push(item interface{}) error {
	return arr.Append(item)
}

func (arr *AbsArray) Prepend(item interface{}) error {
	return arr.Insert(0, item)
}

func (arr *AbsArray) Random() IMix {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	index := r.Intn(arr.Count())
	return arr.Index(index)
}

func (arr *AbsArray) Reverse() IArray {
	newArr := arr.NewEmpty()
	for i := 0; i < newArr.Count() - 1; i++ {
		newArr.Append(arr.Index(newArr.Count() - 1 - i).ToInterface())
	}
	return newArr
}

func (arr *AbsArray) Shuffle() IArray {
	indexs := make([]int, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		indexs[i] = i
	}
	rand.Shuffle(arr.Count(), func(i, j int) {
		indexs[i], indexs[j] = indexs[j], indexs[i]
	})

	newArr := arr.NewEmpty()
	for i := 0; i < len(indexs); i ++ {
		newArr.Append(arr.Index(indexs[i]).ToInterface())
	}
	return newArr
}

func (arr *AbsArray) Column(string) (IArray, error) {
	return nil, errors.New("format not support")
}

func (arr *AbsArray) KeyBy(key string) (IMap, error) {
	return nil, errors.New("format not support")
}

func (arr *AbsArray) Pluck(key string) (IArray, error) {
	return nil, errors.New("format not support")
}

func (arr *AbsArray) SortBy(key string) (IArray, error) {
	return nil, errors.New("format not support")
}

func (arr *AbsArray) SortByDesc(key string) (IArray, error) {
	return nil, errors.New("format not support")
}

func (arr *AbsArray) SetCompare(func(a interface{}, b interface{}) int) {
	panic("SetCompare: not Implement")
}

func (arr *AbsArray) Max() IMix {
	max := 0
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), arr.Index(max).ToInterface()) > 0 {
			max = i
		}
	}
	return arr.Index(max)
}

func (arr *AbsArray) Min() IMix {
	min := 0
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), arr.Index(min).ToInterface()) < 0 {
			min = i
		}
	}
	return arr.Index(min)
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
	// TODO: 这个等map思考清楚再进行开发
	panic("CountBy: not Implement")
}

func (arr *AbsArray) Diff(arr2 IArray) IArray {
	newArr := arr.NewEmpty()
	for i := 0; i < arr.Count(); i++ {
		if arr2.Contains(arr.Index(i).ToInterface()) == false {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsArray) Sort() IArray {
	newArr := arr.NewEmpty()

	isContained := func(arr []int, item int) bool {
		for i := 0; i < len(arr); i++ {
			if item == arr[i] {
				return true
			}
		}
		return false
	}


	sorted := make([]int, 0, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		min := -1
		for j := 0; j < arr.Count(); j++ {
			if isContained(sorted, j) {
				continue
			}

			if min == -1 {
				min = j
				continue
			}

			if arr.compare(arr.Index(j).ToInterface(), arr.Index(min).ToInterface()) <= 0 {
				min = j
				continue
			}
		}

		sorted = append(sorted, min)
		newArr.Append(arr.Index(min).ToInterface())
	}
	return newArr
}

func (arr *AbsArray) SortDesc() IArray {
	newArr := arr.NewEmpty()

	isContained := func(arr []int, item int) bool {
		for i := 0; i < len(arr); i++ {
			if item == arr[i] {
				return true
			}
		}
		return false
	}


	sorted := make([]int, 0, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		max := -1
		for j := 0; j < arr.Count(); j++ {
			if isContained(sorted, j) {
				continue
			}

			if max == -1 {
				max = j
				continue
			}

			if arr.compare(arr.Index(j).ToInterface(), arr.Index(max).ToInterface()) >= 0 {
				max = j
				continue
			}
		}

		sorted = append(sorted, max)
		newArr.Append(arr.Index(max).ToInterface())
	}
	return newArr
}

func (arr *AbsArray) Join(split string, format ...func(item interface{}) string) string {
	var ret strings.Builder
	for i := 0; i < arr.Count(); i++ {
		if len(format) == 0 {
			ret.WriteString(fmt.Sprintf("%v", arr.Index(i).ToInterface()))
		} else {
			f := format[0]
			ret.WriteString(f(arr.Index(i).ToInterface()))
		}

		if i != arr.Count() - 1 {
			ret.WriteString(split)
		}
	}
	return ret.String()
}

func (arr *AbsArray) Avg() IMix{
	panic("Avg: not Implement")
}

func (arr *AbsArray) Median() (IMix, error) {
	newArr := arr.Sort()
	if newArr.Count() % 2 == 0 {
		imax, err := newArr.Index(newArr.Count() / 2 - 1).Add(newArr.Index(newArr.Count() / 2 + 1))
		if err != nil {
			return nil, err
		}
		return imax.Div(2)
	}

	return newArr.Index(newArr.Count() / 2 + 1), nil
}

func (arr *AbsArray) Mode() IMix {
	m := arr.CountBy()
	maxCount, _ := m.Values().Max().ToInt()
	k := m.Search(maxCount)
	return k
}

func (arr *AbsArray) Sum() IMix {
	mix := NewMix(arr.Index(0).ToInterface())
	for i := 0; i < arr.Count(); i++ {
		mix.Add(arr.Index(i))
	}
	return mix
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
	ret := make([]string, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToString()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsArray) ToInt64() ([]int64, error) {
	ret := make([]int64, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToInt64()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsArray) ToInt() ([]int, error) {
	ret := make([]int, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToInt()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsArray) ToMix() []IMix {
	ret := make([]IMix, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		ret[i] = arr.Index(i)
	}
	return ret
}

func (arr *AbsArray) ToFloat64() ([]float64, error) {
	ret := make([]float64, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToFloat64()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsArray) ToFloat32() ([]float32, error) {
	ret := make([]float32, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToFloat32()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}
